package twitter

import (
	"context"
	"net/http"

	"io/ioutil"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Favorites", func() {
	Context("ListFavorites", func() {
		It("should return error when http request fails", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 400,
						Body:       ioutil.NopCloser(strings.NewReader(`{"errors": [{"code": 400, "message": "oops"}]}`)),
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

			_, err := client.ListFavorites(ctx, ListFavoritesParams{})
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

					Ω(req.FormValue("screen_name")).Should(Equal("sumscreenname"))
					Ω(req.FormValue("user_id")).Should(Equal("sumid"))

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

			_, err := client.ListFavorites(ctx, ListFavoritesParams{
				UserID:          "sumid",
				ScreenName:      "sumscreenname",
				Count:           3,
				SinceID:         "12345",
				MaxID:           "12346",
				IncludeEntities: true,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Context("CreateFavorite", func() {
		It("should return error when http request fails", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 400,
						Body:       ioutil.NopCloser(strings.NewReader(`{"errors": [{"code": 400, "message": "oops"}]}`)),
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

			_, err := client.CreateFavorite(ctx, CreateFavoriteParameters{})
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("oops"))
		})

		It("should return successfully", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
					}

					Ω(req.FormValue("id")).Should(Equal("12345"))
					Ω(req.FormValue("include_entities")).Should(Equal("true"))

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

			_, err := client.CreateFavorite(ctx, CreateFavoriteParameters{
				ID:              "12345",
				IncludeEntities: true,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Context("DestroyFavorite", func() {
		It("should return error when http request fails", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 400,
						Body:       ioutil.NopCloser(strings.NewReader(`{"errors": [{"code": 400, "message": "oops"}]}`)),
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

			_, err := client.DestroyFavorite(ctx, DestroyFavoriteParameters{})
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("oops"))
		})

		It("should return successfully", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
					}

					Ω(req.FormValue("id")).Should(Equal("12345"))
					Ω(req.FormValue("include_entities")).Should(Equal("true"))

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

			_, err := client.DestroyFavorite(ctx, DestroyFavoriteParameters{
				ID:              "12345",
				IncludeEntities: true,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
