package twitter

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// GetConfiguration calls the Twitter /help/configuration.json endpoint.
func (c *Client) GetConfiguration(ctx context.Context, params SearchTweetsParams) (*Configuration, error) {
	resp, err := c.do(ctx, "GET", "https://api.twitter.com/1.1/help/configuration.json", values)
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
