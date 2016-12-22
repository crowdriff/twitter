package twitter

import "net/http"

// HTTPMock represents mock struct for http client to implement HTTPClient interface
type HTTPMock struct {
	DoFn func(req *http.Request) (*http.Response, error)
}

// Do method for HTTPMock in implementing HTTPClient interface
func (h *HTTPMock) Do(req *http.Request) (*http.Response, error) {
	return h.DoFn(req)
}
