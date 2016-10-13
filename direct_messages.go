package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// DirectMessage ...
type DirectMessage struct {
	CreatedAt           string   `json:"created_at"`
	Entities            Entities `json:"entities"`
	ID                  int64    `json:"id"`
	IDStr               string   `json:"id_str"`
	Recipient           User     `json:"recipient"`
	RecipientID         int64    `json:"recipient_id"`
	RecipientScreenName string   `json:"recipient_screen_name"`
	Sender              User     `json:"sender"`
	SenderID            int64    `json:"sender_id"`
	SenderScreenName    string   `json:"sender_screen_name"`
	Text                string   `json:"text"`
}

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
	var directMessages []DirectMessage
	err = json.NewDecoder(resp.Body).Decode(resp.Body)
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
	if params.ID != "" {
		values.Set("id", params.ID)
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	resp, err := c.do(ctx, "POST", "https://api.twitter.com/1.1/direct_messages/destroy.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var directMessage DirectMessage
	err = json.NewDecoder(resp.Body).Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return &DirectMessageResponse{
		DirectMessage: directMessage,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}

// SentDirectMessagesParams ...
type SentDirectMessagesParams struct {
	SinceID         string
	MaxID           string
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
	var directMessages []DirectMessage
	err = json.NewDecoder(resp.Body).Decode(resp.Body)
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
	if params.Page > 1 {
		values.Set("page", strconv.Itoa(params.Page))
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// ShowDirectMessage calls the Twitter /direct_messages/show.json endpoint.
func (c *Client) ShowDirectMessage(ctx context.Context, directMessageID string) (*DirectMessageResponse, error) {
	values := url.Values{}
	if directMessageID != "" {
		values.Set("id", directMessageID)
	}
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/direct_messages/show.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var directMessage DirectMessage
	err = json.NewDecoder(resp.Body).Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return &DirectMessageResponse{
		DirectMessage: directMessage,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}
