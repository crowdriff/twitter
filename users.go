package twitter

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// SearchUsersParams represents parameters to send to /users/search.json Twitter endpoint
type SearchUsersParams struct {
	Q               string
	Page            int
	Count           int
	IncludeEntities bool
}

// SearchUsers calls Twitter endpoint /users/search.json
func (c *Client) SearchUsers(ctx context.Context, params SearchUsersParams) (*UsersResponse, error) {
	values := searchUsersToQuery(params)
	urlStr := "https://api.twitter.com/1.1/users/search.json"
	return c.handleUsersResponse(ctx, "GET", urlStr, values)
}

func searchUsersToQuery(params SearchUsersParams) url.Values {
	values := url.Values{}
	values.Set("q", params.Q)
	if params.Page != 0 {
		values.Set("page", strconv.Itoa(params.Page))
	}
	if params.Count != 0 {
		values.Set("count", strconv.Itoa(params.Count))
	}
	if params.IncludeEntities {
		values.Set("include_entities", "true")
	}
	return values
}

// ShowUserParams reperesents parameters to send to /users/show.json Twitter endpoint
type ShowUserParams struct {
	UserID          string
	ScreenName      string
	IncludeEntities bool
}

// ShowUser calls Twitter endpoint /users/show.json
func (c *Client) ShowUser(ctx context.Context, params ShowUserParams) (*UserResponse, error) {
	values := showUserToQuery(params)
	urlStr := "https://api.twitter.com/1.1/users/show.json"
	return c.handleUserResponse(ctx, "GET", urlStr, values)
}

func showUserToQuery(params ShowUserParams) url.Values {
	values := url.Values{}

	values.Set("user_id", params.UserID)
	values.Set("screen_name", params.ScreenName)
	if params.IncludeEntities {
		values.Set("include_entities", "true")
	}

	return values
}

// LookupUsersParams represents parameters to send to /users/lookup.json Twitter endpoint
type LookupUsersParams struct {
	ScreenName      []string
	UserID          []string
	IncludeEntities bool
}

// LookupUsers calls Twitter endpoint /users/lookup.json
func (c *Client) LookupUsers(ctx context.Context, params LookupUsersParams) (*UsersResponse, error) {
	values := lookupUsersToQuery(params)
	urlStr := "https://api.twitter.com/1.1/users/lookup.json"
	return c.handleUsersResponse(ctx, "POST", urlStr, values)
}

func lookupUsersToQuery(params LookupUsersParams) url.Values {
	values := url.Values{}
	if len(params.ScreenName) > 0 {
		values.Set("screen_name", strings.Join(params.ScreenName, ","))
	}
	if len(params.UserID) > 0 {
		values.Set("user_id", strings.Join(params.UserID, ","))
	}
	if params.IncludeEntities {
		values.Set("include_entities", "true")
	}

	return values
}
