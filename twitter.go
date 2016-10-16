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

// UserResponse represents a single User from Twitter.
type UserResponse struct {
	User      User
	RateLimit RateLimit
}

// UserIDsResponse represents a list of User IDs from Twitter.
type UserIDsResponse struct {
	UserIDs   []string
	RateLimit RateLimit
}

// UserIDsPageResponse represents a page of UserIDs from Twitter.
type UserIDsPageResponse struct {
	UserIDPage UserIDPage
	RateLimit  RateLimit
}
