package twitter

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
