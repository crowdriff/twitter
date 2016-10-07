package twitter

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
)

// Credentials represents a token/secret pair of auth credentials.
type Credentials struct {
	Token  string
	Secret string
}

// OAuthClient is the interface for making an OAuth HTTP request using the
// provided context, HTTP method, optional consumer credentials, URL string,
// and URL query parameters. It returns the corresponding HTTP response or
// error.
type OAuthClient interface {
	Do(ctx context.Context, method string, creds *Credentials, urlStr string, values url.Values) (*http.Response, error)
}

// HTTPClient is the interface for making HTTP requests. It accepts an HTTP
// request and returns the corresponding HTTP response or error.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client represents the client used to make requests to the Twitter API 1.1.
type Client struct {
	httpClient   HTTPClient
	oauthClient  *oauth.Client
	defaultCreds *oauth.Credentials
}

// NewClient returns a new Client instance using the provided application
// credentials, optional default consumer credentials, and optional HTTPClient.
func NewClient(appCreds Credentials, defaultCreds Credentials, httpClient HTTPClient) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	oauthClient := oauth.Client{
		Credentials: oauth.Credentials{
			Token:  appCreds.Token,
			Secret: appCreds.Secret,
		},
	}
	return &Client{
		httpClient:  httpClient,
		oauthClient: &oauthClient,
		defaultCreds: &oauth.Credentials{
			Token:  defaultCreds.Token,
			Secret: defaultCreds.Secret,
		},
	}
}

// Do implements the OAuthClient interface. It is used to make an OAuth HTTP
// request with the provided context, HTTP method, OAuth credentials, URL
// string, and URL query parameters. It returns the corresponding HTTP response
// or error.
func (c *Client) Do(ctx context.Context, method string, creds *Credentials, urlStr string, values url.Values) (*http.Response, error) {
	// Set credentials.
	oauthCreds := c.defaultCreds
	if creds != nil {
		oauthCreds = &oauth.Credentials{
			Token:  creds.Token,
			Secret: creds.Secret,
		}
	}
	// Create HTTP request.
	req, err := http.NewRequest(method, urlStr, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	err = c.oauthClient.SetAuthorizationHeader(req.Header, oauthCreds, method, req.URL, values)
	if err != nil {
		return nil, err
	}
	// Make request.
	return c.httpClient.Do(req)
}

// StartFilterStream starts and returns a new Stream using the provided context,
// stream filter parameters, and optional stream error callback.
func (c *Client) StartFilterStream(ctx context.Context, params StreamFilterParams, errFn StreamErrFn) *Stream {
	return newFilterStream(ctx, c, params, errFn)
}
