package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// NoRetweetIDs calls the Twitter /friendships/no_retweets/ids.json endpoint.
func (c *Client) NoRetweetIDs(ctx context.Context, stringifyIDs bool) (*UserIDsResponse, error) {
	values := url.Values{}
	if stringifyIDs {
		values.Set("stringify_ids", "true")
	} else {
		values.Set("stringify_ids", "false")
	}
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/friendships/no_retweets/ids.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var userIDs []string
	err = json.NewDecoder(resp.Body).Decode(&userIDs)
	if err != nil {
		return nil, err
	}
	return &UserIDsResponse{
		UserIDs:   userIDs,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

// friendshipsDirection calls the Twitter /friendships/{direction}.json endpoint.
func (c *Client) friendshipsDirection(ctx context.Context, direction string,
	cursor string, stringifyIDs bool) (*UserIDsPageResponse, error) {
	values := url.Values{}
	if cursor != "" {
		values.Set("cursor", cursor)
	}
	if stringifyIDs {
		values.Set("stringify_ids", "true")
	} else {
		values.Set("stringify_ids", "false")
	}
	resp, err := c.do(ctx, "GET", fmt.Sprintf(
		"https://api.twitter.com/1.1/friendships/%s.json", direction), values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var res UserIDPage
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &UserIDsPageResponse{
		UserIDPage: res,
		RateLimit:  getRateLimit(resp.Header),
	}, nil
}

// FriendshipsIncoming calls the Twitter /friendships/incoming.json endpoint.
func (c *Client) FriendshipsIncoming(ctx context.Context, direction string,
	cursor string, stringifyIDs bool) (*UserIDsPageResponse, error) {
	return c.friendshipsDirection(ctx, "incoming", cursor, stringifyIDs)
}

// FriendshipsOutgoing calls the Twitter /friendships/outgoing.json endpoint.
func (c *Client) FriendshipsOutgoing(ctx context.Context, direction string,
	cursor string, stringifyIDs bool) (*UserIDsPageResponse, error) {
	return c.friendshipsDirection(ctx, "outgoing", cursor, stringifyIDs)
}
