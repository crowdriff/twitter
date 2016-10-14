package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// MentionsTimelineParams represents the query parameters for a
// /statuses/mentions_timeline.json request.
type MentionsTimelineParams struct {
	Count              int    `json:"count"`
	SinceID            string `json:"since_id"`
	MaxID              string `json:"max_id"`
	TrimUser           bool   `json:"trim_user"`
	ContributorDetails bool   `json:"contributor_details"`
	ExcludeEntities    bool   `json:"exclude_entities"`
}

// MentionsTimeline calls the Twitter /statuses/mentions_timeline.json endpoint.
func (c *Client) MentionsTimeline(ctx context.Context, params MentionsTimelineParams) (*TweetsResponse, error) {
	values := mentionsTimelineToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/statuses/mentions_timeline.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func mentionsTimelineToQuery(params MentionsTimelineParams) url.Values {
	values := url.Values{}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.ContributorDetails {
		values.Set("contributor_details", "true")
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// UserTimelineParams represents the query parameters for a
// /statuses/user_timeline.json request.
type UserTimelineParams struct {
	UserID             string `json:"user_id"`
	ScreenName         string `json:"screen_name"`
	SinceID            string `json:"since_id"`
	Count              int    `json:"count"`
	MaxID              string `json:"max_id"`
	TrimUser           bool   `json:"trim_user"`
	ExcludeReplies     bool   `json:"exclude_replies"`
	ContributorDetails bool   `json:"contributor_details"`
	ExcludeRTS         bool   `json:"exclude_rts"`
}

// UserTimeline calls the Twitter /statuses/user_timeline.json endpoint.
func (c *Client) UserTimeline(ctx context.Context, params UserTimelineParams) (*TweetsResponse, error) {
	values := userTimelineToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/statuses/user_timeline.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func userTimelineToQuery(params UserTimelineParams) url.Values {
	values := url.Values{}
	if params.UserID != "" {
		values.Set("user_id", params.UserID)
	}
	if params.ScreenName != "" {
		values.Set("screen_name", params.ScreenName)
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.ExcludeReplies {
		values.Set("exclude_replies", "true")
	}
	if params.ContributorDetails {
		values.Set("contributor_details", "true")
	}
	if params.ExcludeRTS {
		values.Set("include_rts", "false")
	}
	return values
}

// HomeTimelineParams represents the query parameters for a
// /statuses/home_timeline.json request.
type HomeTimelineParams struct {
	Count              int    `json:"count"`
	SinceID            string `json:"since_id"`
	MaxID              string `json:"max_id"`
	TrimUser           bool   `json:"trim_user"`
	ExcludeReplies     bool   `json:"exclude_replies"`
	ContributorDetails bool   `json:"contributor_details"`
	ExcludeEntities    bool   `json:"exclude_entities"`
}

// HomeTimeline calls the Twitter /statuses/home_timeline.json endpoint.
func (c *Client) HomeTimeline(ctx context.Context, params HomeTimelineParams) (*TweetsResponse, error) {
	values := homeTimelineToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/statuses/home_timeline.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func homeTimelineToQuery(params HomeTimelineParams) url.Values {
	values := url.Values{}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.ExcludeReplies {
		values.Set("exclude_replies", "true")
	}
	if params.ContributorDetails {
		values.Set("contributor_details", "true")
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// RetweetsOfMeParams represents the query parameters for a
// /statuses/retweets_of_me.json request.
type RetweetsOfMeParams struct {
	Count               int    `json:"count"`
	SinceID             string `json:"since_id"`
	MaxID               string `json:"max_id"`
	TrimUser            bool   `json:"trim_user"`
	ExcludeEntities     bool   `json:"exclude_entities"`
	ExcludeUserEntities bool   `json:"exclude_user_entities"`
}

// RetweetsOfMe calls the Twitter /statuses/retweets_of_me.json endpoint.
func (c *Client) RetweetsOfMe(ctx context.Context, params RetweetsOfMeParams) (*TweetsResponse, error) {
	values := retweetsOfMeToQuery(params)
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/statuses/retweets_of_me.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func retweetsOfMeToQuery(params RetweetsOfMeParams) url.Values {
	values := url.Values{}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	if params.ExcludeUserEntities {
		values.Set("include_user_entities", "false")
	}
	return values
}

// RetweetsOfTweetParams represents the query parameters for a
// /statuses/retweets/:id.json request.
type RetweetsOfTweetParams struct {
	ID       string `json:"id"`
	Count    int    `json:"count"`
	TrimUser bool   `json:"trim_user"`
}

// RetweetsOfTweet calls the Twitter /statuses/retweets/:id.json endpoint.
func (c *Client) RetweetsOfTweet(ctx context.Context, params RetweetsOfTweetParams) (*TweetsResponse, error) {
	values := retweetsOfTweetToQuery(params)
	urlStr := "https://api.twitter.com/1.1/statuses/retweets/" + params.ID + ".json"
	resp, err := c.do(ctx, "GET", urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func retweetsOfTweetToQuery(params RetweetsOfTweetParams) url.Values {
	values := url.Values{}
	if params.Count > 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	return values
}

// ShowTweetParams represents the query parameters for a
// /statuses/show/:id.json request.
type ShowTweetParams struct {
	ID               string `json:"id"`
	TrimUser         bool   `json:"trim_user"`
	IncludeMyRetweet bool   `json:"include_my_retweet"`
	ExcludeEntities  bool   `json:"exclude_entities"`
}

// ShowTweet calls the Twitter /statuses/show/:id.json endpoint.
func (c *Client) ShowTweet(ctx context.Context, params ShowTweetParams) (*TweetResponse, error) {
	values := showTweetToQuery(params)
	urlStr := "https://api.twitter.com/1.1/statuses/show.json"
	resp, err := c.do(ctx, "GET", urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func showTweetToQuery(params ShowTweetParams) url.Values {
	values := url.Values{}
	values.Set("id", params.ID)
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if params.IncludeMyRetweet {
		values.Set("include_my_retweet", "true")
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

// DestroyTweetParams represents the query parameters for a
// /statuses/destroy/:id.json request.
type DestroyTweetParams struct {
	ID       string `json:"id"`
	TrimUser bool   `json:"trim_user"`
}

// DestroyTweet calls the Twitter /statuses/destroy/:id.json endpoint.
func (c *Client) DestroyTweet(ctx context.Context, params DestroyTweetParams) (*TweetResponse, error) {
	values := destroyTweetToQuery(params)
	urlStr := "https://api.twitter.com/1.1/statuses/destroy/" + params.ID + ".json"
	resp, err := c.do(ctx, "POST", urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func destroyTweetToQuery(params DestroyTweetParams) url.Values {
	values := url.Values{}
	values.Set("id", params.ID)
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	return values
}

// UpdateTweetParams represents the query parameters for a
// /statuses/update.json request.
type UpdateTweetParams struct {
	Status             string
	InReplyToStatusID  string
	PossiblySensitive  bool
	Location           *Location
	PlaceID            string
	DisplayCoordinates bool
	TrimUser           bool
	MediaIDs           []string
}

// UpdateTweet calls the Twitter /statuses/update.json endpoint.
func (c *Client) UpdateTweet(ctx context.Context, params UpdateTweetParams) (*TweetResponse, error) {
	values := updateTweetToQuery(params)
	urlStr := "https://api.twitter.com/1.1/statuses/update.json"
	resp, err := c.do(ctx, "POST", urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func updateTweetToQuery(params UpdateTweetParams) url.Values {
	values := url.Values{}
	values.Set("status", params.Status)
	if params.InReplyToStatusID != "" {
		values.Set("in_reply_to_status_id", params.InReplyToStatusID)
	}
	if params.PossiblySensitive {
		values.Set("possibly_sensitive", "true")
	}
	if params.Location != nil {
		values.Set("lat", strconv.FormatFloat(params.Location.Lat, 'f', -1, 64))
		values.Set("long", strconv.FormatFloat(params.Location.Long, 'f', -1, 64))
	}
	if params.PlaceID != "" {
		values.Set("place_id", params.PlaceID)
	}
	if params.DisplayCoordinates {
		values.Set("display_coordinates", "true")
	}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	if len(params.MediaIDs) > 0 {
		values.Set("media_ids", strings.Join(params.MediaIDs, ","))
	}
	return values
}

// RetweetParams represents the query parameters for a
// /statuses/retweet/:id.json request.
type RetweetParams struct {
	ID       string
	TrimUser bool
}

// Retweet calls the Twitter /statuses/retweet/:id.json endpoint.
func (c *Client) Retweet(ctx context.Context, params RetweetParams) (*TweetResponse, error) {
	values := retweetToQuery(params)
	urlStr := "https://api.twitter.com/1.1/statuses/retweet/" + params.ID + ".json"
	resp, err := c.do(ctx, "POST", urlStr, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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

func retweetToQuery(params RetweetParams) url.Values {
	values := url.Values{}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	return values
}

// UnretweetParams represents the query parameters for a
// /statuses/retweet/:id.json request.
type UnretweetParams struct {
	ID       string
	TrimUser bool
}

// Unretweet calls the Twitter /statuses/retweet/:id.json endpoint.
func (c *Client) Unretweet(ctx context.Context, params UnretweetParams) (*TweetResponse, error) {
	values := unretweetToQuery(params)
	urlStr := "https://api.twitter.com/1.1/statuses/unretweet/" + params.ID + ".json"
	resp, err := c.do(ctx, "POST", urlStr, values)
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

func unretweetToQuery(params UnretweetParams) url.Values {
	values := url.Values{}
	if params.TrimUser {
		values.Set("trim_user", "true")
	}
	return values
}
