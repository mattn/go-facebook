package graph

import (
	"os"
	"time"
)

type Video struct {
	// The video ID. Available to everyone on Facebook by default.
	ID string
	// The profile (user or page) that created the video. Available to everyone on Facebook by default. Contains id and name fields
	From Object
	// The users who are tagged in this video. Available to everyone on Facebook by default. An array of objects containing id and name fields
	Tags []Object
	// The video title or caption. Available to everyone on Facebook by default.
	Name string
	// The html element that may be embedded in an Web page to play the video. Available to everyone on Facebook by default. Contains a valid URL.
	EmbedHtml string
	// The icon that Facebook displays when video are published to the Feed. Available to everyone on Facebook by default. Contains a valid URL.
	Icon string
	// A URL to the raw, playable video file. Available to everyone on Facebook by default. Contains a valid URL.
	Source string
	// The time the video was initially published. Available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime.
	CreatedTime *time.Time
	// The last time the video or its caption were updated. Available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime.
	UpdatedTime *time.Time

	// Connections
	comments string
	picture  string
}

// Gets all of the comments on this Video. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (v *Video) GetComments() (cs []Comment, err os.Error) {
	if v.comments == "" {
		err = os.NewError("Error: Video.GetComments: The comments URL is empty.")
	}
	return getComments(v.comments)
}

// Gets the image which represents the content of the video.
// Publicly available. Returns an HTTP 302 URL string with the location set to the picture URL.
func (v *Video) GetPicture() (pic *Picture, err os.Error) {
	if v.picture == "" {
		err = os.NewError("Error: Video.GetPicture: The picture URL is empty.")
	}
	return NewPicture(v.picture), err
}

func getVideos(url string) (vs []Video, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
	vs = make([]Video, len(data))
	for i, v := range data {
		vs[i], err = parseVideo(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Parses Video data. Returns nil for err if no error appeared.
func parseVideo(value map[string]interface{}) (v Video, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			v.ID = val.(string)
		case "from":
			v.From = parseObject(val.(map[string]interface{}))
		case "tags":
			v.Tags = parseObjects(val.([]interface{}))
		case "name":
			v.Name = val.(string)
		case "embed_html":
			v.EmbedHtml = val.(string)
		case "icon":
			v.Icon = val.(string)
		case "source":
			v.Source = val.(string)
		case "created_time":
			v.CreatedTime, err = parseTime(val.(string))
		case "updated_time":
			v.UpdatedTime, err = parseTime(val.(string))
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, va := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "comments":
					v.comments = va.(string)
				case "picture":
					v.picture = va.(string)
				}
			}
		}
	}
	return
}
