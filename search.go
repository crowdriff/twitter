package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// SearchTweetsParams represents the query parameters for a /search/tweets.json
// request.
type SearchTweetsParams struct {
	Query            string
	Geocode          string
	Lang             string
	Locale           string
	ResultType       string
	Count            int
	Until            string
	SinceID          string
	MaxID            string
	ExcludeEntities  bool
	ExtendedEntities bool
	Callback         string
}

type searchRes struct {
	Statuses []Tweet `json:"statuses"`
}

// SearchTweets calls the Twitter /search/tweets.json endpoint.
func (c *Client) SearchTweets(ctx context.Context, params SearchTweetsParams) (*TweetsResponse, error) {
	values := searchTweetsToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/search/tweets.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res searchRes
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &TweetsResponse{
		Tweets:    res.Statuses,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func searchTweetsToQuery(params SearchTweetsParams) url.Values {
	values := url.Values{}
	values.Set("q", params.Query)
	if params.Geocode != "" {
		values.Set("geocode", params.Geocode)
	}
	if params.Lang != "" {
		values.Set("lang", params.Lang)
	}
	if params.Locale != "" {
		values.Set("locale", params.Locale)
	}
	if params.ResultType != "" {
		values.Set("result_type", params.ResultType)
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.Until != "" {
		values.Set("until", params.Until)
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	if params.ExtendedEntities {
		values.Set("tweet_mode", "extended")
	}
	if params.Callback != "" {
		values.Set("callback", params.Callback)
	}
	return values
}
