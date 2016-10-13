package twitter

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
