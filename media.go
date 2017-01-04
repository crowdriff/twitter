package twitter

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

// MediaUploadParameters represents the query parameters for /media/upload.json request
//
// INIT requires: Command=INIT, MediaType, TotalBytes
// APPEND requires: Command=APPEND, MediaID, Media OR MediaData, SegmentIndex
// FINALIZE requires: Command=FINALIZE, MediaID
//
type MediaUploadParameters struct {
	Command          string
	MediaType        string
	TotalBytes       int
	Media            []byte
	MediaID          string
	MediaData        string
	MediaCategory    string
	SegmentIndex     int
	AdditionalOwners []string
}

// MediaUpload calls the Twitter endpoint /media/upload.json
func (c *Client) MediaUpload(ctx context.Context, params MediaUploadParameters) (*MediaUploadResponse, error) {
	query, err := mediaUploadToQuery(params)
	if err != nil {
		return nil, err
	}

	urlStr := "https://upload.twitter.com/1.1/media/upload.json"
	return c.handleMediaUpload(ctx, "POST", urlStr, query)
}

type mediaUploadQueryResponse struct {
	ContentType string
	Body        io.Reader
}

func mediaUploadToQuery(params MediaUploadParameters) (mediaUploadQueryResponse, error) {
	queryResponse := mediaUploadQueryResponse{}
	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)
	if params.Command != "" {
		bodyWriter.WriteField("command", params.Command)
	}

	if params.MediaType != "" {
		bodyWriter.WriteField("media_type", params.MediaType)
	}

	if params.TotalBytes != 0 {
		bodyWriter.WriteField("total_bytes", strconv.Itoa(params.TotalBytes))
	}

	if params.MediaID != "" {
		bodyWriter.WriteField("media_id", params.MediaID)
	}

	if params.MediaCategory != "" {
		bodyWriter.WriteField("media_category", params.MediaCategory)
	}

	if params.Command == "APPEND" {
		bodyWriter.WriteField("segment_index", strconv.Itoa(params.SegmentIndex))
	}

	if len(params.AdditionalOwners) > 0 {
		bodyWriter.WriteField("additional_owners", strings.Join(params.AdditionalOwners, ","))
	}

	if len(params.Media) > 0 {
		part, err := bodyWriter.CreateFormFile("media", "")
		if err != nil {
			return queryResponse, err
		}
		_, err = part.Write(params.Media)
		if err != nil {
			return queryResponse, err
		}
	} else if params.MediaData != "" {
		bodyWriter.WriteField("media_data", params.MediaData)
	}

	queryResponse.Body = body
	queryResponse.ContentType = bodyWriter.FormDataContentType()

	err := bodyWriter.Close()
	if err != nil {
		return queryResponse, err
	}

	return queryResponse, nil
}
