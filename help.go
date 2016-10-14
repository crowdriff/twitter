package twitter

import (
	"context"
	"encoding/json"
	"net/url"
)

// GetConfiguration calls the Twitter /help/configuration.json endpoint.
func (c *Client) GetConfiguration(ctx context.Context) (*ConfigurationResponse, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/configuration.json", url.Values{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var config Configuration
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &ConfigurationResponse{
		Configuration: config,
		RateLimit:     getRateLimit(resp.Header),
	}, nil
}

// GetLanguages calls Twitter help/langauges.json endpoint.
func (c *Client) GetLanguages(ctx context.Context) (*LanguagesResponse, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/langauges.json", url.Values{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var ls []Language
	err = json.NewDecoder(resp.Body).Decode(&ls)
	if err != nil {
		return nil, err
	}
	return &LanguagesResponse{
		Languages: ls,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

// GetPrivacy calls Twitter help/privacy.json endpoint.
func (c *Client) GetPrivacy(ctx context.Context) (*PrivacyResponse, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/privacy.json", url.Values{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var p map[string]string
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return nil, err
	}
	return &PrivacyResponse{
		Privacy:   p,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}

// GetTOS calls Twitter help/tos.json endpoint.
func (c *Client) GetTOS(ctx context.Context) (*TOSResponse, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/tos.json", url.Values{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = checkResponse(resp); err != nil {
		return nil, err
	}
	var tos map[string]string
	err = json.NewDecoder(resp.Body).Decode(&tos)
	if err != nil {
		return nil, err
	}
	return &TOSResponse{
		TOS:       tos,
		RateLimit: getRateLimit(resp.Header),
	}, nil
}
