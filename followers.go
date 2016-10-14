package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// FollowerListParams represents the query parameters for a /search/tweets.json
// request.
type FollowerListParams struct {
	UserID              string
	ScreenName          string
	Cursor              string
	Count               int
	SkipStatus          bool
	IncludeUserEntities bool
}

// GetFollowerList calls the Twitter /followers/list.json endpoint.
func (c *Client) GetFollowerList(ctx context.Context, params FollowerListParams) (*FollowerListResponse, error) {
	values := followerListToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/followers/list.json", values)
	if err != nil {
		return nil, err
	}
	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var users map[string][]User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return &FollowerListResponse{
		Users:     users,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func followerListToQuery(params FollowerListParams) url.Values {
	values := url.Values{}
	if params.UserID != "" {
		values.Set("user_id", params.UserID)
	}
	if params.ScreenName != "" {
		values.Set("screen_name", params.ScreenName)
	}
	if params.Cursor != "" {
		values.Set("cursor", params.Cursor)
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.SkipStatus {
		values.Set("skip_status", "false")
	}
	if params.IncludeUserEntities {
		values.Set("include_entities", "false")
	}
	return values
}

type FollowerIDsParams struct {
	UserID       string
	ScreenName   string
	Cursor       string
	StringifyIDs bool
	Count        int
}

// GetFollowerIDs calls the Twitter /followers/ids.json endpoint.
func (c *Client) GetFollowerIDs(ctx context.Context, params FollowerIDsParams) (*FollowerIDsResponse, error) {
	values := followerIDsToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/followers/list.json", values)
	if err != nil {
		return nil, err
	}
	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var ids FollowerIDs
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return nil, err
	}

	return &FollowerIDsResponse{
		FollowerIDs: ids,
		RateLimit:   getRateLimit(resp.Header),
	}, nil
}

func followerIDsToQuery(params FollowerIDsParams) url.Values {
	values := url.Values{}
	if params.UserID != "" {
		values.Set("user_id", params.UserID)
	}
	if params.ScreenName != "" {
		values.Set("screen_name", params.ScreenName)
	}
	if params.Cursor != "" {
		values.Set("cursor", params.Cursor)
	}
	if params.StringifyIDs {
		values.Set("stringify_ids", "false")
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	return values
}
