package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// NoRetweetIDs calls the Twitter /friendships/no_retweets/ids.json endpoint.
func (c *Client) NoRetweetIDs(ctx context.Context, stringifyIDs bool) (*UserIDsResponse, error) {
	values := url.Values{}
	values.Set("stringify_ids", strconv.FormatBool(stringifyIDs))
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/friendships/no_retweets/ids.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
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
	values.Set("stringify_ids", strconv.FormatBool(stringifyIDs))
	resp, err := c.do(ctx, "GET", fmt.Sprintf(
		"https://api.twitter.com/1.1/friendships/%s.json", direction), values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
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
func (c *Client) FriendshipsIncoming(ctx context.Context, cursor string,
	stringifyIDs bool) (*UserIDsPageResponse, error) {
	return c.friendshipsDirection(ctx, "incoming", cursor, stringifyIDs)
}

// FriendshipsOutgoing calls the Twitter /friendships/outgoing.json endpoint.
func (c *Client) FriendshipsOutgoing(ctx context.Context, cursor string,
	stringifyIDs bool) (*UserIDsPageResponse, error) {
	return c.friendshipsDirection(ctx, "outgoing", cursor, stringifyIDs)
}

// FriendshipsCreate calls the Twitter /friendships/create.json endpoint.
func (c *Client) FriendshipsCreate(ctx context.Context, screenName,
	userID string, follow bool) (*UserResponse, error) {
	values := url.Values{}
	if screenName != "" {
		values.Set("screen_name", screenName)
	}
	if userID != "" {
		values.Set("user_id", userID)
	}
	values.Set("follow", strconv.FormatBool(follow))
	resp, err := c.do(ctx, "POST",
		"https://api.twitter.com/1.1/friendships/create.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res User
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		User:      res,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

// FriendshipsDestroy calls the Twitter /friendships/destroy.json endpoint.
func (c *Client) FriendshipsDestroy(ctx context.Context, screenName,
	userID string) (*UserResponse, error) {
	values := url.Values{}
	if screenName != "" {
		values.Set("screen_name", screenName)
	}
	if userID != "" {
		values.Set("user_id", userID)
	}
	resp, err := c.do(ctx, "POST",
		"https://api.twitter.com/1.1/friendships/destroy.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res User
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		User:      res,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

type friendshipsDetail struct {
	RelationshipDetail `json:"relationship"`
}

// FriendshipsUpdate calls the Twitter /friendships/update.json endpoint.
func (c *Client) FriendshipsUpdate(ctx context.Context, screenName,
	userID string, device, retweets bool) (*RelationshipDetailResponse, error) {
	values := url.Values{}
	if screenName != "" {
		values.Set("screen_name", screenName)
	}
	if userID != "" {
		values.Set("user_id", userID)
	}
	values.Set("device", strconv.FormatBool(device))
	values.Set("retweets", strconv.FormatBool(retweets))

	resp, err := c.do(ctx, "POST",
		"https://api.twitter.com/1.1/friendships/update.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res friendshipsDetail
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &RelationshipDetailResponse{
		Relationship: res.RelationshipDetail,
		RateLimit:    getRateLimit(resp.Header),
	}, nil
}

// FriendshipsShow calls the Twitter /friendships/show.json endpoint.
func (c *Client) FriendshipsShow(ctx context.Context, sourceID, targetID int64,
	sourceScreenName, targetScreenName string) (*RelationshipDetailResponse, error) {
	values := url.Values{}
	if sourceID > 0 {
		values.Set("source_id", strconv.FormatInt(sourceID, 10))
	}
	if targetID > 0 {
		values.Set("target_id", strconv.FormatInt(targetID, 10))
	}
	if sourceScreenName != "" {
		values.Set("source_screen_name", sourceScreenName)
	}
	if targetScreenName != "" {
		values.Set("target_screen_name", targetScreenName)
	}
	resp, err := c.do(ctx, "GET",
		"https://api.twitter.com/1.1/friendships/show.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res friendshipsDetail
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &RelationshipDetailResponse{
		Relationship: res.RelationshipDetail,
		RateLimit:    getRateLimit(resp.Header),
	}, nil
}

type friendshipsRelationship struct {
	Relationship `json:"relationship"`
}

// FriendshipsLookup calls the Twitter /friendships/lookup.json endpoint.
func (c *Client) FriendshipsLookup(ctx context.Context,
	screenNames, userIDs []string) (*RelationshipResponse, error) {
	values := url.Values{}
	if len(screenNames) > 0 {
		values.Set("screen_name", strings.Join(screenNames, ","))
	}
	if len(userIDs) > 0 {
		values.Set("user_id", strings.Join(userIDs, ","))
	}
	resp, err := c.do(ctx, "GET",
		"https://api.twitter.com/1.1/friendships/lookup.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res friendshipsRelationship
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &RelationshipResponse{
		Relationship: res.Relationship,
		RateLimit:    getRateLimit(resp.Header),
	}, nil
}
