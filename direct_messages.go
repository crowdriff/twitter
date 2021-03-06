package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// GetDirectMessagesParams ...
type GetDirectMessagesParams struct {
	SinceID         string
	MaxID           string
	Count           int
	ExcludeEntities bool
	SkipStatus      bool
}

// GetDirectMessages calls the Twitter /direct_messages.json endpoint.
func (c *Client) GetDirectMessages(ctx context.Context, params GetDirectMessagesParams) (*DirectMessagesResponse, error) {
	values := getDirectMessagesToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/direct_messages.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var directMessages []DirectMessage
	err = json.NewDecoder(resp.Body).Decode(&directMessages)
	if err != nil {
		return nil, err
	}
	return &DirectMessagesResponse{
		DirectMessages: directMessages,
		RateLimit:      getRateLimit(resp.Header),
	}, nil
}

// getDirectMessagesToQuery ...
func getDirectMessagesToQuery(params GetDirectMessagesParams) url.Values {
	values := url.Values{}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	if params.SkipStatus {
		values.Set("skip_status", "true")
	}
	return values
}

// DestroyDirectMessageParams  ...
type DestroyDirectMessageParams struct {
	ID              string
	ExcludeEntities bool
}

// DestroyDirectMessage calls the Twitter /direct_messages/destroy.json endpoint.
func (c *Client) DestroyDirectMessage(ctx context.Context, params DestroyDirectMessageParams) (*DirectMessageResponse, error) {
	values := url.Values{}
	values.Set("id", params.ID)
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	resp, err := c.do(ctx, "POST", "https://api.twitter.com/1.1/direct_messages/destroy.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var directMessage DirectMessage
	err = json.NewDecoder(resp.Body).Decode(&directMessage)
	if err != nil {
		return nil, err
	}
	return &DirectMessageResponse{
		DirectMessage: directMessage,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}

// NewDirectMessageParams ...
type NewDirectMessageParams struct {
	UserID     string
	ScreenName string
	Text       string
}

// NewDirectMessage calls the Twitter /direct_messages/sent.json endpoint.
func (c *Client) NewDirectMessage(ctx context.Context, params NewDirectMessageParams) (*DirectMessageResponse, error) {
	values := newDirectMessageToQuery(params)
	resp, err := c.do(ctx, "POST", "https://api.twitter.com/1.1/direct_messages/new.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var directMessage DirectMessage
	err = json.NewDecoder(resp.Body).Decode(&directMessage)
	if err != nil {
		return nil, err
	}
	return &DirectMessageResponse{
		DirectMessage: directMessage,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}

// newDirectMessageToQuery ...
func newDirectMessageToQuery(params NewDirectMessageParams) url.Values {
	values := url.Values{}
	if params.UserID != "" {
		values.Set("user_id", params.UserID)
	}
	if params.ScreenName != "" {
		values.Set("screen_name", params.ScreenName)
	}
	values.Set("text", params.Text)
	return values
}

// SentDirectMessagesParams ...
type SentDirectMessagesParams struct {
	SinceID         string
	MaxID           string
	Count           int
	Page            int
	ExcludeEntities bool
}

// SentDirectMessages calls the Twitter /direct_messages/sent.json endpoint.
func (c *Client) SentDirectMessages(ctx context.Context, params SentDirectMessagesParams) (*DirectMessagesResponse, error) {
	values := sentDirectMessagesToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/direct_messages/sent.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var directMessages []DirectMessage
	err = json.NewDecoder(resp.Body).Decode(&directMessages)
	if err != nil {
		return nil, err
	}
	return &DirectMessagesResponse{
		DirectMessages: directMessages,
		RateLimit:      getRateLimit(resp.Header),
	}, nil
}

// sentDirectMessagesToQuery ...
func sentDirectMessagesToQuery(params SentDirectMessagesParams) url.Values {
	values := url.Values{}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.Page > 0 {
		values.Set("page", strconv.Itoa(params.Page))
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// ShowDirectMessage calls the Twitter /direct_messages/show.json endpoint.
func (c *Client) ShowDirectMessage(ctx context.Context, id string) (*DirectMessageResponse, error) {
	values := url.Values{"id": []string{id}}
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/direct_messages/show.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var directMessage DirectMessage
	err = json.NewDecoder(resp.Body).Decode(&directMessage)
	if err != nil {
		return nil, err
	}
	return &DirectMessageResponse{
		DirectMessage: directMessage,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}
