package twitter

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Media", func() {
	Context("MediaUpload", func() {
		hm := HTTPMock{
			DoFn: func(req *http.Request) (*http.Response, error) {
				r := &http.Response{
					Body: ioutil.NopCloser(strings.NewReader("")),
				}

				switch req.FormValue("command") {
				case "INIT":
					Ω(req.FormValue("media_type")).ShouldNot(Equal(""))
					Ω(req.FormValue("total_bytes")).Should(Equal("1"))
					r.StatusCode = 201
					r.Body = ioutil.NopCloser(strings.NewReader(`{"media_id":12345, "media_id_string":"12345"}`))
				case "APPEND":
					Ω(req.FormValue("media_id")).Should(Equal("12345"))
					r.StatusCode = 204
				case "FINALIZE":
					Ω(req.FormValue("media_id")).Should(Equal("12345"))
					r.StatusCode = 200
					r.Body = ioutil.NopCloser(strings.NewReader(`{"media_id":12345, "media_id_string":"12345"}`))
				case "FAIL":
					r.StatusCode = 400
					r.Body = ioutil.NopCloser(strings.NewReader(`{"errors": [{"code": 400, "message": "oops"}]}`))
					return r, nil
				default:
					return r, errors.New("command not understood")
				}

				return r, nil
			},
		}

		c := Client{
			httpClient: &hm,
			oauthClient: &oauth.Client{
				Credentials: oauth.Credentials{
					Token:  "",
					Secret: "",
				},
			},
			accessCreds: &oauth.Credentials{
				Token:  "",
				Secret: "",
			},
		}

		It("should return an error when http request fails", func() {
			ctx := context.Background()

			params := MediaUploadParameters{
				Command: "FAIL",
			}

			_, err := c.MediaUpload(ctx, params)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("oops"))
		})

		It("should return a media id when command = INIT", func() {
			ctx := context.Background()

			params := MediaUploadParameters{
				Command:    "INIT",
				TotalBytes: 1,
				MediaType:  "image/jpeg",
			}

			res, err := c.MediaUpload(ctx, params)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(res.Media.MediaID).Should(Equal(int64(12345)))
		})

		It("should return when command = APPEND", func() {
			ctx := context.Background()

			params := MediaUploadParameters{
				Command:      "APPEND",
				MediaID:      "12345",
				Media:        []byte("imagewhatever"),
				SegmentIndex: 0,
			}

			_, err := c.MediaUpload(ctx, params)
			Ω(err).ShouldNot(HaveOccurred())
		})

	})
})
