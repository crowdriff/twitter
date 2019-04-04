package twitter

import "time"

// Tweet represents a Twitter tweet object.
type Tweet struct {
	Contributor          []Contributor          `json:"contributors"`
	Coordinates          *Coordinates           `json:"coordinates"`
	CreatedAt            string                 `json:"created_at"`
	Entities             Entities               `json:"entities"`
	ExtendedEntities     ExtendedEntities       `json:"extended_entities"`
	ExtendedTweet        *ExtendedTweet         `json:"extended_tweet"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	FullText             string                 `json:"full_text"`
	ID                   int64                  `json:"id"`
	IDStr                string                 `json:"id_str"`
	InReplyToScreenName  string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string                 `json:"in_reply_to_user_id_str"`
	Lang                 string                 `json:"lang"`
	Place                *Place                 `json:"place"`
	PossiblySensitive    bool                   `json:"possibly_sensitive"`
	QuotedStatusID       int64                  `json:"quoted_status_id"`
	QuotedStatusIDStr    string                 `json:"quoted_status_id_str"`
	QuotedStatus         *Tweet                 `json:"quoted_status"`
	RetweetCount         int                    `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	RetweetedStatus      *Tweet                 `json:"retweeted_status"`
	Scopes               map[string]interface{} `json:"scopes"`
	Source               string                 `json:"source"`
	Text                 string                 `json:"text"`
	Truncated            bool                   `json:"truncated"`
	User                 User                   `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
	WithheldScope        string                 `json:"withheld_scope"`
}

// ExtendedTweet represents the full information for an extended tweet.
type ExtendedTweet struct {
	FullText         string           `json:"full_text"`
	DisplayTextRange []int            `json:"display_text_range"`
	Entities         Entities         `json:"entities"`
	ExtendedEntities ExtendedEntities `json:"extended_entities"`
}

// Coordinates represents the geographic location of a Tweet as reported by the
// user or client application. The coordinates array is formatted as geoJSON
// (longitude first, then latitude).
type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

// DirectMessage ...
type DirectMessage struct {
	CreatedAt           string   `json:"created_at"`
	Entities            Entities `json:"entities"`
	ID                  int64    `json:"id"`
	IDStr               string   `json:"id_str"`
	Recipient           User     `json:"recipient"`
	RecipientID         int64    `json:"recipient_id"`
	RecipientScreenName string   `json:"recipient_screen_name"`
	Sender              User     `json:"sender"`
	SenderID            int64    `json:"sender_id"`
	SenderScreenName    string   `json:"sender_screen_name"`
	Text                string   `json:"text"`
}

// ExtendedEntities provides metadata about the media entities present.
type ExtendedEntities struct {
	Media []MediaEntity `json:"media"`
}

// Entities provides metadata and additonal contextual information about twitter
// content.
type Entities struct {
	Hashtags     []HashtagEntity     `json:"hashtags"`
	Media        []MediaEntity       `json:"media"`
	URL          URLEntities         `json:"url"`
	URLs         []URLEntity         `json:"urls"`
	UserMentions []UserMentionEntity `json:"user_mentions"`
}

// HashtagEntity represents hashtags which have been parsed out of the Tweet
// text.
type HashtagEntity struct {
	Indices []int  `json:"indices"`
	Text    string `json:"text"`
}

// MediaEntity represents media elements uploaded with a Tweet.
type MediaEntity struct {
	DisplayURL        string     `json:"display_url"`
	ExpandedURL       string     `json:"expanded_url"`
	ID                int64      `json:"id"`
	IDStr             string     `json:"id_str"`
	Indices           []int      `json:"indices"`
	MediaURL          string     `json:"media_url"`
	MediaURLHTTPS     string     `json:"media_url_https"`
	Sizes             MediaSizes `json:"sizes"`
	SourceStatusID    int64      `json:"source_status_id"`
	SourceStatusIDStr string     `json:"source_status_id_str"`
	Type              string     `json:"type"`
	URL               string     `json:"url"`
	VideoInfo         VideoInfo  `json:"video_info"`
}

// VideoInfo represents information specific to videos.
type VideoInfo struct {
	AspectRatio    []int          `json:"aspect_ratio"`
	DurationMillis int            `json:"duration_millis"`
	Variants       []VideoVariant `json:"variants"`
}

// VideoVariant represents video info for a single bitrate.
type VideoVariant struct {
	Bitrate     int    `json:"bitrate"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}

// MediaSize represents the height, width, and resize string of a media element.
type MediaSize struct {
	Height int    `json:"h"`
	Resize string `json:"resize"`
	Width  int    `json:"w"`
}

// MediaSizes represents a range of media sizes.
type MediaSizes struct {
	Large  MediaSize `json:"large"`
	Medium MediaSize `json:"medium"`
	Small  MediaSize `json:"small"`
	Thumb  MediaSize `json:"thumb"`
}

// UserMentionEntity represents other Twitter users mentioned in the text of a
// tweet.
type UserMentionEntity struct {
	ID         int64  `json:"id"`
	IDStr      string `json:"id_str"`
	Indices    []int  `json:"indices"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

// BoundingBox represents a bounding box of coordinates which encloses a Place.
type BoundingBox struct {
	Coordinates [][][2]float64 `json:"coordinates"`
	Type        string         `json:"type"`
}

// Place represents a specific named location with corresponding geo
// coordinates.
type Place struct {
	Attributes  map[string]string `json:"attributes"`
	BoundingBox BoundingBox       `json:"bounding_box"`
	Country     string            `json:"country"`
	CountryCode string            `json:"country_code"`
	FullName    string            `json:"full_name"`
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	PlaceType   string            `json:"place_type"`
	URL         string            `json:"url"`
}

// Contributor represents a user who contributed to the authorship of a tweet.
type Contributor struct {
	ID         int64  `json:"id"`
	IDStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

// CreatedAtTime returns a time.Time version of the created date.
func (t *Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// URLEntities represents a list of url entities.
type URLEntities struct {
	URLs []URLEntity `json:"urls"`
}

// URLEntity represents URLs included in the text of a Tweet or within textual
// fields of a user object.
type URLEntity struct {
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
	Indices     []int  `json:"indices"`
	URL         string `json:"url"`
}

// User represents a twitter user object, as received using the twitter REST
// API.
type User struct {
	ContributorsEnabled            bool     `json:"contributors_enabled"`
	CreatedAt                      string   `json:"created_at"`
	DefaultProfile                 bool     `json:"default_profile"`
	DefaultProfileImage            bool     `json:"default_profile_image"`
	Description                    string   `json:"description"`
	Entities                       Entities `json:"entities"`
	FavouritesCount                int      `json:"favourites_count"`
	FollowRequestSent              bool     `json:"follow_request_sent"`
	Following                      bool     `json:"following"`
	FollowersCount                 int      `json:"followers_count"`
	FriendsCount                   int      `json:"friends_count"`
	GeoEnabled                     bool     `json:"geo_enabled"`
	ID                             int64    `json:"id"`
	IDStr                          string   `json:"id_str"`
	IsTranslator                   bool     `json:"is_translator"`
	Lang                           string   `json:"lang"`
	ListedCount                    int      `json:"listed_count"`
	Location                       string   `json:"location"`
	Name                           string   `json:"name"`
	Notifications                  bool     `json:"notifications"`
	ProfileBackgroundColor         string   `json:"profile_background_color"`
	ProfileBackgroundImageURL      string   `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool     `json:"profile_background_tile"`
	ProfileBannerURL               string   `json:"profile_banner_url"`
	ProfileImageURL                string   `json:"profile_image_url"`
	ProfileImageURLHTTPS           string   `json:"profile_image_url_https"`
	ProfileLinkColor               string   `json:"profile_link_color"`
	ProfileSidebarBorderColor      string   `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string   `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string   `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool     `json:"profile_use_background_image"`
	Protected                      bool     `json:"protected"`
	ScreenName                     string   `json:"screen_name"`
	ShowAllInlineMedia             bool     `json:"show_all_inline_media"`
	Status                         *Tweet   `json:"status"` // can be null
	StatusesCount                  int      `json:"statuses_count"`
	TimeZone                       string   `json:"time_zone"`
	URL                            string   `json:"url"`
	UTCOffset                      int      `json:"utc_offset"`
	Verified                       bool     `json:"verified"`
	WithheldInCountries            string   `json:"withheld_in_countries"`
	WithheldScope                  string   `json:"withheld_scope"`
}

// Configuration represents the configuration object received from Twitter help/configuration endpoint
type Configuration struct {
	CharactersReservedPerMedia int                  `json:"characters_reserved_per_media"`
	DMTextCharacterLimit       int                  `json:"dm_text_character_limit"`
	MaxMediaPerUpload          int                  `json:"max_media_per_upload"`
	PhotoSizeLimit             int                  `json:"photo_size_limit"`
	PhotoSizes                 map[string]PhotoSize `json:"photo_sizes"`
	ShortURLLength             int                  `json:"short_url_length"`
	ShortURLLengthHTTPS        int                  `json:"short_url_length_https"`
	NonUsernamePaths           []string             `json:"non_username_paths"`
}

// PhotoSize represents the photo size object iside configuration objectt
type PhotoSize struct {
	H      int    `json:"h"`
	Resize string `json:"resize"`
	W      int    `json:"w"`
}

// Language represents the language obejct received from Twitter help/langauges
type Language struct {
	Code   string `json:"code"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// OEmbed represents the response body for an oembed request.
type OEmbed struct {
	CacheAge     string `json:"cache_age"`
	URL          string `json:"url"`
	ProviderURL  string `json:"provider_url"`
	ProviderName string `json:"provider_name"`
	AuthorName   string `json:"author_name"`
	Version      string `json:"version"`
	AuthorURL    string `json:"author_url"`
	Type         string `json:"type"`
	HTML         string `json:"html"`
	Height       *int   `json:"height"`
	Width        *int   `json:"width"`
}

// IDs represents a paginated list of string IDs.
type IDs struct {
	IDs               []string `json:"ids"`
	PreviousCursor    int64    `json:"previous_cursor"`
	PreviousCursotStr string   `json:"previous_cursor_str"`
	NextCursor        int64    `json:"next_cursor"`
	NextCursorStr     string   `json:"next_cursor_str"`
}

// RelationshipTarget represents a Twitter user's relationship with a source
type RelationshipTarget struct {
	IDStr      string `json:"id_str"`
	ID         int64  `json:"id"`
	ScreenName string `json:"screen_name"`
	Following  bool   `json:"following"`
	FollowedBy bool   `json:"followed_by"`
}

// Relationship represents a pair of Twitter users' relationship with each other
type Relationship struct {
	Target RelationshipTarget `json:"target"`
	Source RelationshipTarget `json:"source"`
}

// Friendship represents a relationship between two Twitter users
type Friendship struct {
	Relationship Relationship `json:"relationship"`
}

// FriendshipLookup represents a Twitter user's relationship to currently authorized user
type FriendshipLookup struct {
	Name        string   `json:"name"`
	ScreenName  string   `json:"screen_name"`
	ID          int64    `json:"id"`
	IDStr       string   `json:"id_str"`
	Connections []string `json:"connections"`
}

// Image represents info on a Twitter image
type Image struct {
	ImageType string `json:"image_type"`
	W         int    `json:"w"`
	H         int    `json:"h"`
}

// Video represents info on a Twitter video
type Video struct {
	VideoType string `json:"video_type"`
}

// MediaUpload represents a Twitter's response object for media upload
type MediaUpload struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	Size             int    `json:"size"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	Image            Image  `json:"image"`
	Video            Video  `json:"video"`
}

// InsightsData represents the object that wraps the response from Twitter's Insights API.
type InsightsData map[string]TweetIDs

//TweetIDs represents a subset of the return object for the Twitter Insights API.
type TweetIDs map[string]MediaInsights

// MediaInsights represents the smallest piece of Twitter's response object for a call to their insights API.
type MediaInsights struct {
	Favourites  string `json:"favorites"`
	Replies     string `json:"replies"`
	Retweets    string `json:"retweets"`
	VideoViews  string `json:"video_views"`
	Impressions string `json:"impressions"`
	Engagements string `json:"engagements"`
}
