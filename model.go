package twitter

import "time"

// Tweet represents a Twitter tweet object.
type Tweet struct {
	Contributor          []Contributor          `json:"contributors"`
	Coordinates          *Coordinates           `json:"coordinates"`
	CreatedAt            string                 `json:"created_at"`
	Entities             Entities               `json:"entities"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
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

// UserIDPage represents a single page of UserIDs returned by the "following or "followers" endpoint.
type UserIDPage struct {
	IDs               []string `json:"ids"`
	NextCursor        int64    `json:"next_cursor"`
	NextCursorStr     string   `json:"next_cursor_str"`
	PreviousCursor    int64    `json:"previous_cursor"`
	PreviousCursorStr string   `json:"previous_cursor_str"`
}
