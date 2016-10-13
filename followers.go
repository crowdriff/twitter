package twitter

import (
	"context"
	"encoding/json"
	"net/url"
)

// FollowerListParams represents the query parameters for a /search/tweets.json
// request.
type FollowerListParams struct {
	UserID              string
	ScreenName          string
	Cursor              int
	Count               int
	SkipStatus          bool
	IncludeUserEntities bool
}

// SearchTweets calls the Twitter /search/tweets.json endpoint.
// func (c *Client) SearchTweets(ctx context.Context, params SearchTweetsParams) (*TweetsResponse, error) {
// 	values := searchTweetsToQuery(params)
// 	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/search/tweets.json", values)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	var tweets []Tweet
// 	err = json.NewDecoder(resp.Body).Decode(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &TweetsResponse{
// 		Tweets:    tweets,
// 		RateLimit: getRateLimit(resp.Header),
// 	}, nil
// }

// func searchTweetsToQuery(params SearchTweetsParams) url.Values {
// 	values := url.Values{}
// 	values.Set("q", params.Query)
// 	if params.Geocode != "" {
// 		values.Set("geocode", params.Geocode)
// 	}
// 	if params.Lang != "" {
// 		values.Set("lang", params.Lang)
// 	}
// 	if params.Locale != "" {
// 		values.Set("locale", params.Locale)
// 	}
// 	if params.ResultType != "" {
// 		values.Set("result_type", params.ResultType)
// 	}
// 	if params.Count > 0 {
// 		values.Set("count", strconv.Itoa(params.Count))
// 	}
// 	if params.Until != "" {
// 		values.Set("until", params.Until)
// 	}
// 	if params.SinceID != "" {
// 		values.Set("since_id", params.SinceID)
// 	}
// 	if params.MaxID != "" {
// 		values.Set("max_id", params.MaxID)
// 	}
// 	if params.ExcludeEntities {
// 		values.Set("include_entities", "false")
// 	}
// 	if params.Callback != "" {
// 		values.Set("callback", params.Callback)
// 	}
// 	return values
// }
