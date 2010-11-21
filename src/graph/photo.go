package graph

import (
	"time"
	"os"
)

// An individual photo within an album
type Photo struct {
	// The photo ID. Available to everyone on Facebook by default.
	ID string

	// The profile (user or page) that posted this photo. Available to everyone on Facebook by default.
	From Object

	// The tagged users and their positions in this photo.
	// Available to everyone on Facebook by default.
	// An array of JSON objects, the x and y coordinates are percentages from the left and top edges of the photo, respectively.
	// Tags TODO

	// The caption given to this photo. Available to everyone on Facebook by default.
	Name string

	// The icon that Facebook displays when photos are published to the Feed. Available to everyone on Facebook by default. Contains a valid URL.
	Icon string

	// The full-sized source of the photo. Available to everyone on Facebook by default. Contains a valid URL.
	Source string

	// The height of the photo in pixels. Available to everyone on Facebook by default.
	Height float64

	// The width of the photo in pixels. Available to everyone on Facebook by default.
	Widht float64

	// A link to the photo on Facebook. Available to everyone on Facebook by default. Contains a valid URL.
	Link string

	// The time the photo was initially published. Available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime.
	CreatedTime *time.Time

	// The last time the photo or its caption was updated. Available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime.
	UpdatedTime *time.Time

	// Connections
	comments string
	likes    string
	picture  string
}

// Gets all of the comments on this Photo. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (p *Photo) GetComments() (cs []Comment, err os.Error) {
	if p.comments == "" {
		err = os.NewError("Error: Photo.GetComments: The comments URL is empty.")
	}
	return getComments(p.comments)
}

// Gets the likes on this Photo. Available to everyone on Facebook.
// Returns an array of objects containing the id and name fields.
func (p *Photo) GetLikes() (likes []Object, err os.Error) {
	if p.likes == "" {
		err = os.NewError("Error: Photo.GetLikes: The likes URL is empty.")
	}
	data, err := getData(p.likes)
	if err != nil {
		return
	}
	likes = parseObjects(data)
	return
}

// Gets the album-sized view of the photo. Available to everyone on Facebook by default.
// Publicly available. Returns an HTTP 302 URL string with the location set to the picture URL.
func (p *Photo) GetPicture() (pic *Picture, err os.Error) {
	if p.picture == "" {
		err = os.NewError("Error: Photo.GetPicture: The picture URL is empty.")
	}
	return NewPicture(p.picture), err
}

func getPhotos(url string) (ps []Photo, err os.Error) {
	data, err := getData(url)
	if err != nil {
		return
	}
	ps = make([]Photo, len(data))
	for i, v := range data {
		ps[i], err = parsePhoto(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

/*
 * Parses Photo data. Returns nil for err if no error appeared.
 */
func parsePhoto(value map[string]interface{}) (p Photo, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			p.ID = val.(string)
		case "from":
			p.From = parseObject(val.(map[string]interface{}))
		case "name":
			p.Name = val.(string)
		case "icon":
			p.Icon = val.(string)
		case "soure":
			p.Source = val.(string)
		case "height":
			p.Height = val.(float64)
		case "width":
			p.Widht = val.(float64)
		case "link":
			p.Link = val.(string)
		case "created_time":
			p.CreatedTime, err = parseTime(val.(string))
		case "updated_time":
			p.UpdatedTime, err = parseTime(val.(string))
			// Connections
			// TODO
		}
	}
	return
}
