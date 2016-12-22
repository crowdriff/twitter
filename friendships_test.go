package twitter

import (
	"context"
	"net/http"
	"strings"

	"io/ioutil"

	"github.com/garyburd/go-oauth/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Friendships", func() {
	Context("ShowFriendships", func() {
		It("should return error when http request fails", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 400,
						Body:       ioutil.NopCloser(strings.NewReader(`{"errors":[{"code": 400, "message": "oops"}]}`)),
					}
					return r, nil
				},
			}

			client := Client{
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
			ctx := context.Background()
			_, err := client.ShowFriendships(ctx, ShowFriendshipsParameters{})
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("oops"))
		})

		It("should return successfully", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader("{}")),
					}

					Ω(req.FormValue("source_id")).Should(Equal("12345"))
					Ω(req.FormValue("target_id")).Should(Equal("23456"))

					return r, nil
				},
			}

			client := Client{
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

			ctx := context.Background()

			_, err := client.ShowFriendships(ctx, ShowFriendshipsParameters{
				SourceID: "12345",
				TargetID: "23456",
			})
			Ω(err).ShouldNot(HaveOccurred())

		})
	})

	Context("LookupFriendships", func() {
		It("should return error when http request fails", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 400,
						Body:       ioutil.NopCloser(strings.NewReader(`{"errors":[{"code": 400, "message": "oops"}]}`)),
					}
					return r, nil
				},
			}

			client := Client{
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
			ctx := context.Background()
			_, err := client.LookupFriendships(ctx, LookupFriendshipsParams{})
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("oops"))
		})

		It("should return successfully", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader(`[]`)),
					}

					Ω(req.FormValue("screen_name")).Should(Equal("ascreenname,another"))

					return r, nil
				},
			}

			client := Client{
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

			ctx := context.Background()

			_, err := client.LookupFriendships(ctx, LookupFriendshipsParams{
				ScreenName: []string{"ascreenname", "another"},
			})
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
