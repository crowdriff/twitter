package twitter

import (
	"context"
	"net/url"
	"strconv"
)

//ListFavoritesParams represents the query parameters for a
// /favorites/list.json request
type ListFavoritesParams struct {
	UserID          string
	ScreenName      string
	Count           int
	SinceID         string
	MaxID           string
	ExcludeEntities bool
}

// ListFavorites calls the Twitter /favorites/list.json endpoint
func (c *Client) ListFavorites(ctx context.Context, params ListFavoritesParams) (*TweetsResponse, error) {
	values := listFavoritesToQuery(params)
	urlStr := "https://api.twitter.com/1.1/favorites/list.json"
	return c.handleTweetsResponse(ctx, "GET", urlStr, values)
}

func listFavoritesToQuery(params ListFavoritesParams) url.Values {
	values := url.Values{}
	if params.UserID != "" {
		values.Set("user_id", params.UserID)
	}
	if params.ScreenName != "" {
		values.Set("screen_name", params.ScreenName)
	}
	if params.Count != 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.SinceID != "" {
		values.Set("since_id", params.SinceID)
	}
	if params.MaxID != "" {
		values.Set("max_id", params.MaxID)
	}
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

//CreateFavoriteParameters represents query parameters for a
// /favorites/create.json request
type CreateFavoriteParameters struct {
	ID              string
	ExcludeEntities bool
}

// CreateFavorite calls the Twitter /favorites/create.json endpoint
func (c *Client) CreateFavorite(ctx context.Context, params CreateFavoriteParameters) (*TweetResponse, error) {
	values := createFavoriteToQuery(params)
	urlStr := "https://api.twitter.com/1.1/favorites/create.json"
	return c.handleTweetResponse(ctx, "POST", urlStr, values)
}

func createFavoriteToQuery(params CreateFavoriteParameters) url.Values {
	values := url.Values{}
	values.Set("id", params.ID)
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}

//DestroyFavoriteParameters represents query parameters for a
// /favorites/create.json request
type DestroyFavoriteParameters struct {
	ID              string
	ExcludeEntities bool
}

// DestroyFavorite calls the Twitter /favorites/create.json endpoint
func (c *Client) DestroyFavorite(ctx context.Context, params DestroyFavoriteParameters) (*TweetResponse, error) {
	values := destroyFavoriteToQuery(params)
	urlStr := "https://api.twitter.com/1.1/favorites/destroy.json"
	return c.handleTweetResponse(ctx, "POST", urlStr, values)
}

func destroyFavoriteToQuery(params DestroyFavoriteParameters) url.Values {
	values := url.Values{}
	values.Set("id", params.ID)
	if params.ExcludeEntities {
		values.Set("include_entities", "false")
	}
	return values
}
