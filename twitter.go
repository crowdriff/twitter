package twitter

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
	Privacy   [string]string
	RateLimit RateLimit
}
