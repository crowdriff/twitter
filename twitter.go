package twitter

import (
	"context"
	"encoding/json"
	"net/url"
)

// DirectMessageResponse represents a response from Twitter containing a single DirectMessage.
type DirectMessageResponse struct {
	DirectMessage DirectMessage
	RateLimit     RateLimit
}

// DirectMessagesResponse represents a response from Twitter containing multiple DirectMessages.
type DirectMessagesResponse struct {
	DirectMessages []DirectMessage
	RateLimit      RateLimit
}

// Location represents a set of lat/long coordinates.
type Location struct {
	Lat  float64
	Long float64
}

// TweetResponse represents a response from Twitter containing a single Tweet.
type TweetResponse struct {
	Tweet     Tweet
	RateLimit RateLimit
}

// RateLimitStatusResponse represents a response from Twitter containing multiple RateLimitStatuses.
type RateLimitStatusResponse struct {
	RateLimitsRes RateLimitsRes
	RateLimit     RateLimit
}

// TweetsResponse represents a response from Twitter containing multiple Tweets.
type TweetsResponse struct {
	Tweets    []Tweet
	RateLimit RateLimit
}

// ConfigurationResponse represents a response from Twitter containing configuration.
type ConfigurationResponse struct {
	Configuration Configuration
	RateLimit     RateLimit
}

// LanguagesResponse represents a response from Twitter containing languages.
type LanguagesResponse struct {
	Languages []Language
	RateLimit RateLimit
}

// PrivacyResponse represents a response from Twitter containing privacy.
type PrivacyResponse struct {
	Privacy   map[string]string
	RateLimit RateLimit
}

// TOSResponse represents a response from Twitter containing terms of service.
type TOSResponse struct {
	TOS       map[string]string
	RateLimit RateLimit
}

// OEmbedResponse represents a response from Twitter oembed request.
type OEmbedResponse struct {
	OEmbed    OEmbed
	RateLimit RateLimit
}

// IDsResponse represents a response from Twitter with paginated string IDs.
type IDsResponse struct {
	IDs       IDs
	RateLimit RateLimit
}

// UserResponse represents a response from Twitter with a user object
type UserResponse struct {
	RateLimit RateLimit
	User      User
}

// UsersResponse represents a response from Twitter with a list of user objects
type UsersResponse struct {
	RateLimit RateLimit
	Users     []User
}

// FriendshipResponse represents a response from Twitter with two objects describing the relationship between two users
type FriendshipResponse struct {
	RateLimit  RateLimit
	Friendship Friendship
}

// FriendshipLookupResponse represents a response from Twitter with a list of user relationship details relative to currently authorized user
type FriendshipLookupResponse struct {
	RateLimit        RateLimit
	FriendshipLookup []FriendshipLookup
}

func (c *Client) handleTweetsResponse(ctx context.Context, method, urlStr string, values url.Values) (*TweetsResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var tweets []Tweet
	err = json.NewDecoder(resp.Body).Decode(&tweets)
	if err != nil {
		return nil, err
	}
	return &TweetsResponse{
		Tweets:    tweets,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func (c *Client) handleTweetResponse(ctx context.Context, method, urlStr string, values url.Values) (*TweetResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var tweet Tweet
	err = json.NewDecoder(resp.Body).Decode(&tweet)
	if err != nil {
		return nil, err
	}
	return &TweetResponse{
		Tweet:     tweet,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func (c *Client) handleIDsResponse(ctx context.Context, method, urlStr string, values url.Values) (*IDsResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var ids IDs
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return nil, err
	}
	return &IDsResponse{
		IDs:       ids,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func (c *Client) handleUserResponse(ctx context.Context, method, urlStr string, values url.Values) (*UserResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		User:      user,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func (c *Client) handleUsersResponse(ctx context.Context, method, urlStr string, values url.Values) (*UsersResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}
	return &UsersResponse{
		Users:     users,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

func (c *Client) handleFriendshipResponse(ctx context.Context, method, urlStr string, values url.Values) (*FriendshipResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var friendship Friendship
	err = json.NewDecoder(resp.Body).Decode(&friendship)
	if err != nil {
		return nil, err
	}
	return &FriendshipResponse{
		RateLimit:  getRateLimit(resp.Header),
		Friendship: friendship,
	}, nil
}

func (c *Client) handleFriendshipsResponse(ctx context.Context, method, urlStr string, values url.Values) (*FriendshipLookupResponse, error) {
	resp, err := c.do(ctx, method, urlStr, values)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var friendshipLookup []FriendshipLookup
	err = json.NewDecoder(resp.Body).Decode(&friendshipLookup)
	if err != nil {
		return nil, err
	}
	return &FriendshipLookupResponse{
		RateLimit:        getRateLimit(resp.Header),
		FriendshipLookup: friendshipLookup,
	}, nil
}
