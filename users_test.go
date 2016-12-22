package twitter

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users", func() {
	Context("SearchUsers", func() {
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

			_, err := client.SearchUsers(ctx, SearchUsersParams{})
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

					Ω(req.FormValue("q")).Should(Equal("#crazy"))
					Ω(req.FormValue("page")).Should(Equal("1"))
					Ω(req.FormValue("count")).Should(Equal("10"))
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

			_, err := client.SearchUsers(ctx, SearchUsersParams{
				Q:               "#crazy",
				Page:            1,
				Count:           10,
				IncludeEntities: true,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Context("ShowUser", func() {
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

			_, err := client.ShowUser(ctx, ShowUserParams{})
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

					Ω(req.FormValue("user_id")).Should(Equal("12345"))
					Ω(req.FormValue("screen_name")).Should(Equal("sumname"))
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

			_, err := client.ShowUser(ctx, ShowUserParams{
				UserID:          "12345",
				ScreenName:      "sumname",
				IncludeEntities: true,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Context("LookupUsers", func() {
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

			_, err := client.LookupUsers(ctx, LookupUsersParams{})
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

					Ω(req.FormValue("user_id")).Should(Equal("12345"))
					Ω(req.FormValue("screen_name")).Should(Equal("sumname"))
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

			_, err := client.LookupUsers(ctx, LookupUsersParams{
				UserID:          []string{"12345"},
				ScreenName:      []string{"sumname"},
				IncludeEntities: true,
			})

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

})
