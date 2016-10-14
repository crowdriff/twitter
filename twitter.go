package twitter

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
