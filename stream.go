package twitter

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// StreamFilterParams represents the filter parameters used in a stream.
// https://dev.twitter.com/streaming/overview/request-parameters
type StreamFilterParams struct {
	FilterLevel   string
	Follow        []string
	Language      []string
	Locations     []string
	StallWarnings bool
	Track         []string
}

// StreamMessage represents a message received from a Twitter stream. Fields
// should be checked for existence before being used.
type StreamMessage struct {
	*Tweet
	Delete         *DeleteMessage         `json:"delete"`
	ScrubGeo       *ScrubGeoMessage       `json:"scrub_geo"`
	Limit          *LimitMessage          `json:"limit"`
	StatusWithheld *StatusWithheldMessage `json:"status_withheld"`
	UserWithheld   *UserWithheldMessage   `json:"user_withheld"`
	Disconnect     *DisconnectMessage     `json:"disconnect_message"`
}

// DeleteMessage represents a stream message that a given Tweet has been
// deleted.
type DeleteMessage struct {
	Status struct {
		ID        int64  `json:"id"`
		IDStr     string `json:"id_str"`
		UserID    int64  `json:"user_id"`
		UserIDStr string `json:"user_id_str"`
	} `json:"status"`
}

// ScrubGeoMessage represents a stream message that geolocated data must be
// stripped from a range of Tweets.
type ScrubGeoMessage struct {
	UserID          int64  `json:"user_id"`
	UserIDStr       string `json:"user_id_str"`
	UpToStatusID    int64  `json:"up_to_status_id"`
	UpToStatusIDStr string `json:"up_to_status_id_str"`
}

// LimitMessage represents a stream message that a filtered stream has matched
// more tweets than its current rate limit allows to be delivered.
type LimitMessage struct {
	Track int `json:"track"`
}

// StatusWithheldMessage represents a stream message that the indicated tweet
// has been withheld.
type StatusWithheldMessage struct {
	ID                  int64    `json:"id"`
	UserID              int64    `json:"user_id"`
	WithheldInCountries []string `json:"withheld_in_countries"`
}

// UserWithheldMessage represents a steram message that tweets from the
// indicated user have been withheld.
type UserWithheldMessage struct {
	ID                  int64    `json:"id"`
	WithheldInCountries []string `json:"withheld_in_countries"`
}

// DisconnectMessage represents a stream message that the stream will disconnect
// with the provided code and reason.
// https://dev.twitter.com/streaming/overview/messages-types#disconnect_messages
type DisconnectMessage struct {
	Code       int    `json:"code"`
	StreamName string `json:"stream_name"`
	Reason     string `json:"reason"`
}

// StreamErrFn represents a function that is called when an error is encountered
// in a stream and the connection will be retried. If the StreamErrFn returns
// a non-nil error, the stream will be immediately closed with the error.
type StreamErrFn func(Backoff, error) error

// Stream represents a Twitter stream connection. Messages from the stream can
// be read off the channel returned by Messages. At any point, the stream can be
// manually closed by calling the Close method. When the stream exits, the
// channel returned from the Done method will be closed.
type Stream struct {
	ctx    context.Context
	cancel context.CancelFunc

	client   OAuthClient
	values   url.Values
	endpoint string

	chMessage chan StreamMessage
	chDone    chan struct{}
	closeErr  error
	errFn     StreamErrFn
}

func newFilterStream(ctx context.Context, client OAuthClient, params StreamFilterParams, errFn StreamErrFn) *Stream {
	s := Stream{
		client:    client,
		values:    parseFilterParams(params),
		endpoint:  "https://stream.twitter.com/1.1/statuses/filter.json",
		chMessage: make(chan StreamMessage),
		chDone:    make(chan struct{}),
		errFn:     errFn,
	}
	s.ctx, s.cancel = context.WithCancel(ctx)
	go s.start()
	return &s
}

// Close immediately closes the stream and waits for the stream to completely
// close before returning the stream's shutdown error.
func (s *Stream) Close() error {
	s.cancel()
	<-s.chDone
	return s.Err()
}

// Done returns a channel that is closed when the stream has completely
// shutdown.
func (s *Stream) Done() <-chan struct{} {
	return s.chDone
}

// Err returns the stream's shutdown error after it has been closed. This should
// only be called after the the channel returned from Done has been closed.
func (s *Stream) Err() error {
	return s.closeErr
}

// Messages returns a read-only channel that messages are sent to as they are
// read off of the stream. Messages should be regualrly waiting on this channel,
// otherwise the stream's queue (on Twitter's end) will fill up and cause the
// stream to close.
func (s *Stream) Messages() <-chan StreamMessage {
	return s.chMessage
}

func (s *Stream) notifyError(boff Backoff, err error) error {
	if s.errFn == nil {
		return nil
	}
	return s.errFn(boff, err)
}

func (s *Stream) start() {
	var err error
	defer func() {
		s.cancel()
		s.closeErr = err
		close(s.chDone)
	}()

	boff := &backoff{}
	for {
		err = s.makeRequest(boff)
		select {
		case <-s.ctx.Done():
			err = s.ctx.Err()
			return
		default:
		}
		if err != nil {
			return
		}
		if d := boff.wait(); d > 0 {
			select {
			case <-s.ctx.Done():
				err = s.ctx.Err()
				return
			case <-time.After(d):
			}
		}
	}
}

func (s *Stream) makeRequest(boff *backoff) error {
	// Create a child context for this specific request.
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	// Make HTTP request to open stream.
	resp, err := s.client.Do(ctx, "POST", nil, s.endpoint, s.values)
	if err != nil {
		boff.incNetDelay()
		return s.notifyError(boff, err)
	}
	defer resp.Body.Close()

	// Handle HTTP response.
	switch resp.StatusCode {
	case 200:
		err = s.readMessages(cancel, resp.Body)
		boff.reset()
		return s.notifyError(boff, err)
	case 401, 403, 404, 406, 413, 416:
		err = fmt.Errorf("%d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		return err
	case 420:
		err = errors.New("420: Rate Limited")
		boff.incHTTPDelay(true)
		return s.notifyError(boff, err)
	default:
		err = fmt.Errorf("%d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		boff.incHTTPDelay(false)
		return s.notifyError(boff, err)
	}
}

func (s *Stream) readMessages(cancel context.CancelFunc, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(scanLines)
	for {
		if err := s.readMessage(cancel, scanner); err != nil {
			return err
		}
	}
}

func (s *Stream) readMessage(cancel context.CancelFunc, scanner *bufio.Scanner) error {
	// Set 90 second timeout on receiving a message.
	// https://dev.twitter.com/streaming/overview/connecting
	t := time.AfterFunc(90*time.Second, func() { cancel() })
	// Scan next token.
	ok := scanner.Scan()
	t.Stop()
	if !ok {
		return scanner.Err()
	}
	b := scanner.Bytes()
	if len(b) == 0 || (len(b) == 1 && b[0] == '\n') {
		// Keep-alive.
		log.Println("Keep-alive")
		return nil
	}
	// Parse StreamMessage JSON.
	var sm StreamMessage
	err := json.Unmarshal(b, &sm)
	if err != nil {
		return err
	}
	select {
	case <-s.ctx.Done():
		return s.ctx.Err()
	case s.chMessage <- sm:
		return nil
	}
}

var newMsgBytes = []byte("\r\n")

func scanLines(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, newMsgBytes); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func parseFilterParams(params StreamFilterParams) url.Values {
	values := url.Values{}
	// Write possible settings.
	if params.FilterLevel != "" {
		values.Set("filter_level", params.FilterLevel)
	}
	if params.StallWarnings {
		values.Set("stall_warnings", "true")
	}
	// Write possible filters.
	var buf bytes.Buffer
	if len(params.Follow) > 0 {
		values.Set("follow", commaSeparated(&buf, params.Follow))
	}
	if len(params.Language) > 0 {
		values.Set("language", commaSeparated(&buf, params.Language))
	}
	if len(params.Locations) > 0 {
		values.Set("locations", commaSeparated(&buf, params.Locations))
	}
	if len(params.Track) > 0 {
		values.Set("track", commaSeparated(&buf, params.Track))
	}
	return values
}

func commaSeparated(buf *bytes.Buffer, ss []string) string {
	buf.Reset()
	for i, s := range ss {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s)
	}
	return buf.String()
}
