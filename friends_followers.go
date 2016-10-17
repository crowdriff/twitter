package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// FListParams represents the generic query parameters for /followers/list
// and /friends/list.
type FListParams struct {
	UserID              string
	ScreenName          string
	Cursor              string
	Count               int
	SkipStatus          bool
	ExcludeUserEntities bool
}

// FListResponse represents a response from Twitter containing a
// followers or friends list.
type FListResponse struct {
	Users     []User
	RateLimit RateLimit
}

func fListToQuery(params FListParams) url.Values {
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
	values.Set("skip_status", strconv.FormatBool(params.SkipStatus))
	if params.ExcludeUserEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// getFList calls the Twitter /{following,friends}/list.json endpoint.
func (c *Client) getFList(ctx context.Context, params FListParams, direction string) (
	*FListResponse, error) {
	values := fListToQuery(params)
	resp, err := c.do(ctx, "GET", fmt.Sprintf("https://api.twitter.com/1.1/%s/list.json",
		direction), values)
	if err != nil {
		return nil, err
	}
	err = checkResponse(resp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var users struct {
		Users []User `json:"users"`
	}
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}
	return &FListResponse{
		Users:     users.Users,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

// FollowersList calls the Twitter /followers/list.json endpoint.
func (c *Client) FollowersList(ctx context.Context, params FListParams) (*FListResponse, error) {
	return c.getFList(ctx, params, "followers")
}

// FriendsList calls the Twitter /friends/list.json endpoint.
func (c *Client) FriendsList(ctx context.Context, params FListParams) (*FListResponse, error) {
	return c.getFList(ctx, params, "friends")
}

// FIDsParams reprensents the generic parameters required for /followers/ids
// and /following/ids endpoints.
type FIDsParams struct {
	UserID       string
	ScreenName   string
	Cursor       string
	StringifyIDs bool
	Count        int
}

func fIDsToQuery(params FIDsParams) url.Values {
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

// FollowerIDs calls the Twitter /followers/ids.json endpoint.
func (c *Client) FollowerIDs(ctx context.Context, params FIDsParams) (*IDsResponse, error) {
	values := fIDsToQuery(params)
	urlStr := "https://api.twitter.com/1.1/followers/ids.json"
	return c.handleIDsResponse(ctx, "GET", urlStr, values)
}

// FriendsIDs calls the Twitter /friends/ids.json endpoint.
func (c *Client) FriendsIDs(ctx context.Context, params FIDsParams) (*IDsResponse, error) {
	values := fIDsToQuery(params)
	urlStr := "https://api.twitter.com/1.1/friends/ids.json"
	return c.handleIDsResponse(ctx, "GET", urlStr, values)
}
