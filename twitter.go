package twitter

// TweetsResponse represents a response from Twitter containing multiple Tweets.
type TweetsResponse struct {
	Tweets    []Tweet
	RateLimit RateLimit
}
