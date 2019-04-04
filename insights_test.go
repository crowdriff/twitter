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

var _ = Describe("Insights", func() {
	Context("GetInsights", func() {
		It("should return an error when http request fails", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 400,
						Body:       ioutil.NopCloser(strings.NewReader(`{"errors":[{"code": 400, "message": "oops"}]}`)),
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

			ctx := context.Background()
			params := PostInsightsParams{
				PostIDs: []string{"failure"},
			}

			_, err := c.GetTotalPostInsights(ctx, params)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("oops"))
		})
		It("should return successfully", func() {
			hm := HTTPMock{
				DoFn: func(req *http.Request) (*http.Response, error) {
					r := &http.Response{
						StatusCode: 200,
						Body: ioutil.NopCloser(strings.NewReader(`{
							"data": {
								"232323": {"replies": "12", "favorites": "11"},
								"7767": {"retweets": "0"}
								}
							}`)),
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

			ctx := context.Background()
			params := PostInsightsParams{
				PostIDs: []string{"232323", "7767"},
			}

			resp, err := c.GetTotalPostInsights(ctx, params)
			tweetID := TweetIDs{
				"232323": MediaInsights{
					Favourites: "11",
					Replies:    "12",
				},
				"7767": MediaInsights{
					Retweets: "0",
				},
			}
			d := InsightsData{
				"data": tweetID,
			}

			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp.Insights).Should(Equal(d))
		})
	})
})
