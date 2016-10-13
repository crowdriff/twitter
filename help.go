package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// GetConfiguration calls the Twitter /help/configuration.json endpoint.
func (c *Client) GetConfiguration(ctx context.Context) (*Configuration, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/configuration.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var c Configuration
	err = json.NewDecoder(resp.Body).Decode(c)
	if err != nil {
		return nil, err
	}
	return &ConfigurationResponse{
		Configuration: c,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}

// GetLanguages calls Twitter help/langauges.json endpoint.
func (c *Client) GetLanguages(ctx context.Context) (*[]Language, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/langauges.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ls []Language
	err = json.NewDecoder(resp.Body).Decode(ls)
	if err != nil {
		return nil, err
	}
	return &LanguageResponse{
		Languages: ls,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

// GetPrivacy calls Twitter help/privacy.json endpoint.
func (c *Client) GetPrivacy(ctx context.Context) (*[string]string, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/privacy.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var p [string]string
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return nil, err
	}
	return &PrivacyResponse{
		Privacy:   p,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}
