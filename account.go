package twitter

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// AccountSettings represents all the settings returned by
// the /account/settings endpoint.
type AccountSettings struct {
	AlwaysUseHTTPS      bool   `json:"always_use_https"`
	DiscoverableByEmail bool   `json:"discoverable_by_email"`
	GeoEnabled          bool   `json:"geo_enabled"`
	Language            string `json:"language"`
	Protected           bool   `json:"protected"`
	ScreenName          string `json:"screen_name"`
	ShowAllInlineMedia  bool   `json:"show_all_inline_media"`
	SleepTime           struct {
		Enabled   bool  `json:"enabled"`
		EndTime   int64 `json:"end_time"`
		StartTime int64 `json:"start_time"`
	} `json:"sleep_time"`
	TimeZone struct {
		Name       string `json:"name"`
		TzInfoName string `json:"tzinfo_name"`
		UTCOffset  int64  `json:"utc_offset"`
	} `json:"time_zone"`
	TrendLocation []struct {
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		Name        string `json:"name"`
		ParentID    int64  `json:"parent_id"`
		PlaceType   struct {
			Code int    `json:"code"`
			Name string `json:"name"`
		} `json:"place_type"`
		URL   string `json:"url"`
		WoeID int64  `json:"woeid"`
	} `json:"trend_location"`
	UseCookiePersonalization bool   `json:"use_cookie_personalization"`
	AllowContributorRequest  string `json:"allow_contributor_request"`
}

// GetAccountSettings calls the Twitter /account/settings.json endpoint.
func (c *Client) GetAccountSettings(ctx context.Context) (
	*AccountSettingsResponse, error) {
	resp, err := c.do(ctx, "GET",
		"https://api.twitter.com/1.1/account/settings.json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res AccountSettings
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &AccountSettingsResponse{
		AccountSettings: res,
		RateLimit:       getRateLimit(resp.Header),
	}, nil
}

// AccountSettingsParams represents all the account settings that can be
// modified via the API.
type AccountSettingsParams struct {
	SleepTimeEnabled        *bool
	StartSleepTime          string // ISO8601 format (00-23)
	EndSleepTime            string // ISO8601 format (00-23)
	TimeZone                string
	TrendLocationWoeID      int64
	AllowContributorRequest string
	CurrentPassword         string
	Lang                    string
}

// accountSettingsToQuery transforms an AccountSettingsParams to a url.Values
// that can be tacked onto the URL of the HTTP request.
func accountSettingsToQuery(p *AccountSettingsParams) (url.Values, error) {
	v := url.Values{}
	if p.SleepTimeEnabled != nil {
		v.Set("sleep_time_enabled", strconv.FormatBool(*p.SleepTimeEnabled))
	}
	if p.StartSleepTime != "" {
		v.Set("start_sleep_time", p.StartSleepTime)
	}
	if p.EndSleepTime != "" {
		v.Set("end_sleep_time", p.EndSleepTime)
	}
	if p.TimeZone != "" {
		v.Set("time_zone", p.TimeZone)
	}
	if p.TrendLocationWoeID > 0 {
		v.Set("trend_location_woeid", strconv.FormatInt(p.TrendLocationWoeID, 10))
	}
	if p.AllowContributorRequest != "" {
		if p.CurrentPassword == "" {
			return nil, errors.New("CurrentPassword must be set when updating" +
				" AllowContributorRequest")
		}
		v.Set("allow_contributor_request", p.AllowContributorRequest)
		v.Set("current_password", p.CurrentPassword)
	}
	if p.Lang != "" {
		v.Set("lang", p.Lang)
	}
	return v, nil
}

// UpdateAccountSettings calls the Twitter /account/settings.json endpoint.
func (c *Client) UpdateAccountSettings(ctx context.Context, params AccountSettingsParams) (
	*AccountSettingsResponse, error) {
	values, err := accountSettingsToQuery(&params)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(ctx, "POST",
		"https://api.twitter.com/1.1/account/settings.json", values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res AccountSettings
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &AccountSettingsResponse{
		AccountSettings: res,
		RateLimit:       getRateLimit(resp.Header),
	}, nil
}

// VerifyCredentials calls the Twitter /account/verify_credentials.json endpoint.
func (c *Client) VerifyCredentials(ctx context.Context) (
	*UserResponse, error) {
	resp, err := c.do(ctx, "GET",
		"https://api.twitter.com/1.1/account/verify_credentials.json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var res User
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		User:      res,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}
