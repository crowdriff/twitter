package twitter

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
)

// Client represents the client used to make requests to the Twitter API 1.1.
type Client struct {
	httpClient  HTTPClient
	oauthClient *oauth.Client
	accessCreds *oauth.Credentials

	gzipDisabled bool
}

// NewClient returns a new Client instance using the provided application
// credentials, optional default consumer credentials, and optional HTTPClient.
func NewClient(consumerCreds ConsumerCredentials, accessCreds AccessCredentials, httpClient HTTPClient) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	oauthClient := oauth.Client{
		Credentials: oauth.Credentials{
			Token:  consumerCreds.Key,
			Secret: consumerCreds.Secret,
		},
	}
	return &Client{
		httpClient:  httpClient,
		oauthClient: &oauthClient,
		accessCreds: &oauth.Credentials{
			Token:  accessCreds.Token,
			Secret: accessCreds.Secret,
		},
	}
}

// WithGzipDisabled returns a new shallow copy of the Client with Gzip disabled.
func (c *Client) WithGzipDisabled() *Client {
	newC := *c
	newC.gzipDisabled = true
	return &newC
}

// do readies the request body/query url for simple queries and calls execute
func (c *Client) do(ctx context.Context, method, urlStr string, values url.Values) (*http.Response, error) {
	// Set up request URL and body.
	var body io.Reader
	switch method {
	case "GET", "HEAD":
		urlStr = urlStr + "?" + values.Encode()
		values = nil
	default:
		body = strings.NewReader(values.Encode())
	}

	return c.execute(ctx, method, urlStr, "application/x-www-form-urlencoded", body, values)
}

// execute implements the OAuthClient interface. It is used to make an OAuth HTTP
// request with the provided context, HTTP method, body ioReader, OAuth credentials, URL
// string, and URL query parameters. It returns the corresponding HTTP response
// or error.
func (c *Client) execute(ctx context.Context, method, urlStr, contentType string, body io.Reader, values url.Values) (*http.Response, error) {
	accessCreds := c.accessCredentials(ctx)

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", contentType)
	if !c.gzipDisabled {
		req.Header.Set("Accept-Encoding", "gzip")
	}
	err = c.oauthClient.SetAuthorizationHeader(req.Header, accessCreds, req.Method, req.URL, values)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if c.gzipDisabled || err != nil || !isGzipped(resp.Header) {
		return resp, err
	}
	return gzipResponse(resp)
}

// HTTPClient is the interface for making HTTP requests. It accepts an HTTP
// request and returns the corresponding HTTP response or error.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ConsumerCredentials represents a token/secret pair of auth credentials.
type ConsumerCredentials struct {
	Key    string
	Secret string
}

// oauthClient is the interface for making an OAuth HTTP request using the
// provided context, HTTP method, optional consumer credentials, URL string,
// and URL query parameters. It returns the corresponding HTTP response or
// error.
type oauthClient interface {
	do(ctx context.Context, method string, urlStr string, values url.Values) (*http.Response, error)
}

// RateLimit represents the Twitter rate limit for the credentials used in the
// associated request.
type RateLimit struct {
	Limit     int
	Remaining int
	Reset     int
}

func getRateLimit(h http.Header) RateLimit {
	limit, _ := strconv.Atoi(h.Get("X-Rate-Limit-Limit"))
	remaining, _ := strconv.Atoi(h.Get("X-Rate-Limit-Remaining"))
	reset, _ := strconv.Atoi(h.Get("X-Rate-Limit-Reset"))
	return RateLimit{
		Limit:     limit,
		Remaining: remaining,
		Reset:     reset,
	}
}

// AccessCredentials represents an OAuth token and secret pair.
type AccessCredentials struct {
	Token  string
	Secret string
}

type accCredsKeyType int

const accCredsKey accCredsKeyType = 0

func (c *Client) accessCredentials(ctx context.Context) *oauth.Credentials {
	if val := ctx.Value(accCredsKey); val != nil {
		return val.(*oauth.Credentials)
	}
	return c.accessCreds
}

// WithAccessCredentials returns a new context with the provided parent context
// and AccessCredentials.
func WithAccessCredentials(parent context.Context, creds AccessCredentials) context.Context {
	return context.WithValue(parent, accCredsKey, &oauth.Credentials{
		Token:  creds.Token,
		Secret: creds.Secret,
	})
}
