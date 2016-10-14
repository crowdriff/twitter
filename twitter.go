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
