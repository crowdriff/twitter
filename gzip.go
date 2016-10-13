package twitter

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipReadCloser struct {
	r  io.ReadCloser
	gr *gzip.Reader
}

func (r *gzipReadCloser) Read(p []byte) (int, error) {
	return r.gr.Read(p)
}

func (r *gzipReadCloser) Close() error {
	r.gr.Close()
	return r.r.Close()
}

func isGzipped(h http.Header) bool {
	return strings.Contains(h.Get("Content-Encoding"), "gzip")
}

func gzipResponse(resp *http.Response) (*http.Response, error) {
	gr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &gzipReadCloser{
		r:  resp.Body,
		gr: gr,
	}
	return resp, nil
}
