package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
)

// RateLimitsRes ...
type RateLimitsRes struct {
	RateLimitContext struct {
		AccessToken string `json:"access_token"`
	} `json:"rate_limit_context"`

	Resources struct {
		Help struct {
			Configuration RateLimit `json:"/help/configuration"`
			Languages     RateLimit `json:"/help/languages"`
			Privacy       RateLimit `json:"/help/privacy"`
			TOS           RateLimit `json:"/help/tos"`
		} `json:"help"`
		Users struct {
			Contributees           RateLimit `json:"/users/contributees"`
			Contributors           RateLimit `json:"/users/contributors"`
			Lookup                 RateLimit `json:"/users/lookup"`
			ProfileBanner          RateLimit `json:"/users/profile_banner"`
			Search                 RateLimit `json:"/users/search"`
			ShowID                 RateLimit `json:"/users/show/:id"`
			Suggestions            RateLimit `json:"/users/suggestions"`
			SuggestionsSlug        RateLimit `json:"/users/suggestions/:slug"`
			SuggestionsSlugMembers RateLimit `json:"/users/suggestions/:slug/members"`
		} `json:"users"`
		Search struct {
			Tweets RateLimit `json:"/search/tweets"`
		} `json:"search"`
		Statuses struct {
			HomeTimeline     RateLimit `json:"/statuses/home_timeline"`
			Lookup           RateLimit `json:"/statuses/lookup"`
			MentionsTimeline RateLimit `json:"/statuses/mentions_timeline"`
			OEmbed           RateLimit `json:"/statuses/oembed"`
			Retweeters       RateLimit `json:"/statuses/retweeters/ids"`
			Retweets         RateLimit `json:"/statuses/retweets/:id"`
			RetweetsOfMe     RateLimit `json:"/statuses/retweets_of_me"`
			Show             RateLimit `json:"/statuses/show/:id"`
			UserTimeline     RateLimit `json:"/statuses/user_timeline"`
		} `json:"statuses"`
	} `json:"resources"`
}

// RateLimitStatus calls the Twitter /application/rate_limit_status.json endpoint.
func (c *Client) RateLimitStatus(ctx context.Context, resources []string) (*RateLimitStatusResponse, error) {
	values := url.Values{}
	resourcesParam := strings.Join(resources, ",")
	if resourcesParam != "" {
		values.Set("resources", resourcesParam)
	}
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/application/rate_limit_status.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var res RateLimitsRes
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &RateLimitStatusResponse{
		RateLimitsRes: res,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}
