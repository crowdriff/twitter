package twitter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
)

//PostInsightsParams represents the body parameters for the /insights/engagement/total request
type PostInsightsParams struct {
	PostIDs []string
}

type getInsightsQueryResponse struct {
	ContentType string
	Body        io.Reader
}

type insightsBody struct {
	TweetIDs        []string `json:"tweet_ids"`
	EngagementTypes []string `json:"engagement_types"`
}

var engagements = []string{"impressions", "engagements", "favorites", "retweets", "video_views", "replies"}

//GetTotalPostInsights calls the twitter data api and retrieves insight totals (lifetime metrics) for up to 250 post ids
// Posts older than 90 days cannot be queried using this endpoint
func (c *Client) GetTotalPostInsights(ctx context.Context, params PostInsightsParams) (*PostInsightsResponse, error) {
	urlStr := "https://data-api.twitter.com/insights/engagement/totals"
	query, err := insightsToQuery(params)
	if err != nil {
		return nil, err
	}
	return c.handleGetInsights(ctx, "POST", urlStr, query)
}

func insightsToQuery(params PostInsightsParams) (getInsightsQueryResponse, error) {
	queryResp := getInsightsQueryResponse{}
	if len(params.PostIDs) > 250 {
		return queryResp, errors.New("cannot query more than 250 post ids at one time")
	}

	body := insightsBody{
		EngagementTypes: engagements,
		TweetIDs:        params.PostIDs,
	}
	b, err := json.Marshal(&body)
	if err != nil {
		return queryResp, err
	}

	buf := &bytes.Buffer{}
	_, err = buf.Write(b)
	if err != nil {
		return queryResp, err
	}

	queryResp.Body = buf
	queryResp.ContentType = "application/json"

	return queryResp, err
}
