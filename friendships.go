package twitter

import (
	"context"
	"net/url"
	"strings"
)

// ShowFriendshipsParameters represents parameters for /friendships/show.json Twitter endpoint
type ShowFriendshipsParameters struct {
	SourceID         string
	SourceScreenName string
	TargetID         string
	TargetScreenName string
}

// ShowFriendships calls the Twitter endpoint /friendships/show.json
func (c *Client) ShowFriendships(ctx context.Context, params ShowFriendshipsParameters) (*FriendshipResponse, error) {
	values := showFriendshipToQuery(params)
	urlStr := "https://api.twitter.com/1.1/friendships/show.json"
	return c.handleFriendshipResponse(ctx, "GET", urlStr, values)
}

func showFriendshipToQuery(params ShowFriendshipsParameters) url.Values {
	values := url.Values{}
	if params.SourceID != "" {
		values.Set("source_id", params.SourceID)
	}
	if params.SourceScreenName != "" {
		values.Set("source_screen_name", params.SourceScreenName)
	}
	if params.TargetID != "" {
		values.Set("target_id", params.TargetID)
	}
	if params.TargetScreenName != "" {
		values.Set("target_screen_name", params.TargetScreenName)
	}
	return values
}

// LookupFriendshipsParams represents parameters for /friendships/lookup.json Twitter endpoint
type LookupFriendshipsParams struct {
	ScreenName []string
	UserID     []string
}

// LookupFriendships calls Twitter endpoint /friendships/lookup.json
func (c *Client) LookupFriendships(ctx context.Context, params LookupFriendshipsParams) (*FriendshipLookupResponse, error) {
	values := lookupFriendshipsToQuery(params)
	urlStr := "https://api.twitter.com/1.1/friendships/lookup.json"
	return c.handleFriendshipsResponse(ctx, "GET", urlStr, values)
}

func lookupFriendshipsToQuery(params LookupFriendshipsParams) url.Values {
	values := url.Values{}
	values.Set("screen_name", strings.Join(params.ScreenName, ","))
	values.Set("user_id", strings.Join(params.UserID, ","))
	return values
}
