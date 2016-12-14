package twitter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

// Error represents an individual error from Twitter.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Errors represents an error response from the Twitter API, possibly containing
// multiple individual errors.
type Errors struct {
	Errors   []Error `json:"errors"`
	HTTPCode int     `json:"-"`
}

// Error implemetns the error interface.
func (e *Errors) Error() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(e.HTTPCode))
	buf.WriteString(": ")
	if len(e.Errors) == 0 {
		buf.WriteString("Unknown error")
		return buf.String()
	}
	buf.WriteByte('[')
	for i, err := range e.Errors {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteByte('{')
		buf.WriteString(strconv.Itoa(err.Code))
		buf.WriteString(": ")
		buf.WriteString(err.Message)
		buf.WriteByte('}')
	}
	buf.WriteByte(']')
	return buf.String()
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	errs := Errors{
		HTTPCode: resp.StatusCode,
	}

	_ = json.NewDecoder(resp.Body).Decode(&errs)
	return &errs
}
