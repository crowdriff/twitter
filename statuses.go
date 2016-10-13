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
