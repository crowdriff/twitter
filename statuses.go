package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// MentionsTimelineParams represents the query parameters for a
// /statuses/mentions_timeline.json request.
type MentionsTimelineParams struct {
	Count              int    `json:"count"`
	SinceID            string `json:"since_id"`
	MaxID              string `json:"max_id"`
	TrimUser           bool   `json:"trim_user"`
	ContributorDetails bool   `json:"contributor_details"`
	ExcludeEntities    bool   `json:"exclude_entities"`
}

// MentionsTimeline calls the Twitter /statuses/mentions_timeline.json endpoint.
func (c *Client) MentionsTimeline(ctx context.Context, params MentionsTimelineParams) (*TweetsResponse, error) {
	values := mentionsTimelineToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/statuses/mentions_timeline.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var tweets []Tweet
	err = json.NewDecoder(resp.Body).Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return &TweetsResponse{
		Tweets:    tweets,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func mentionsTimelineToQuery(params MentionsTimelineParams) url.Values {
	values := url.Values{}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.ContributorDetails {
		values.Set("contributor_details", "true")
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// UserTimelineParams represents the query parameters for a
// /statuses/user_timeline.json request.
type UserTimelineParams struct {
	UserID             string `json:"user_id"`
	ScreenName         string `json:"screen_name"`
	SinceID            string `json:"since_id"`
	Count              int    `json:"count"`
	MaxID              string `json:"max_id"`
	TrimUser           bool   `json:"trim_user"`
	ExcludeReplies     bool   `json:"exclude_replies"`
	ContributorDetails bool   `json:"contributor_details"`
	ExcludeRTS         bool   `json:"exclude_rts"`
}

// UserTimeline calls the Twitter /statuses/user_timeline.json endpoint.
func (c *Client) UserTimeline(ctx context.Context, params UserTimelineParams) (*TweetsResponse, error) {
	values := userTimelineToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/statuses/user_timeline.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var tweets []Tweet
	err = json.NewDecoder(resp.Body).Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return &TweetsResponse{
		Tweets:    tweets,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func userTimelineToQuery(params UserTimelineParams) url.Values {
	values := url.Values{}
	if params.UserID != "" {
		values.Set("user_id", params.UserID)
	}
	if params.ScreenName != "" {
		values.Set("screen_name", params.ScreenName)
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.ExcludeReplies {
		values.Set("exclude_replies", "true")
	}
	if params.ContributorDetails {
		values.Set("contributor_details", "true")
	}
	if params.ExcludeRTS {
		values.Set("include_rts", "false")
	}
	return values
}
