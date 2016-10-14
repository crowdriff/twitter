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

// FollowerListResponse represents a response from Twitter containing follower list.
type FollowerListResponse struct {
	Users     map[string][]User
	RateLimit RateLimit
}

// FollowerIDsResponse represents a response from Twitter containing follower ids.
type FollowerIDsResponse struct {
	FollowerIDs FollowerIDs
	RateLimit   RateLimit
}
