package graph

import (
	"os"
	"time"
)

// A photo album.
type Album struct {
	// The photo album ID. Publicly available.
	ID string
	// The profile that created this album. Publicly available. Contains the id and name fields.
	From Object
	// The title of the album. Publicly available.
	Name string
	// The description of the album. Available to everyone in Facebook.
	Description string
	// The location of the album. Available to everyone on Facebook.
	Location string
	// A link to this album on Facebook. Publicly available. Contains a valid URL.
	Link string
	// The privacy settings for the album. Available to everyone on Facebook.
	Privacy string
	// The number of photos in this album. Publicly available.
	Count float64
	// The time the photo album was initially created. Publicly available. Contains a IETF RFC 3339 datetime.
	CreatedTime *time.Time
	// The last time the photo album was updated. Publicly available. Contains a IETF RFC 3339 datetime.
	UpdatedTime *time.Time

	// Connections
	photos   string
	comments string
	picture  string
}

// Gets the photos contained in this album. Publicly available.
// Returns an array of Photo objects.
func (a *Album) GetPhotos() (ps []Photo, err os.Error) {
	if a.photos == "" {
		err = os.NewError("Error: Album.GetPhotos: The photos URL is empty.")
	}
	return getPhotos(a.photos)
}

// Gets the comments made on this album. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (a *Album) GetComments() (cs []Comment, err os.Error) {
	if a.comments == "" {
		err = os.NewError("Error: Album.GetComments: The comments URL is empty.")
	}
	return getComments(a.comments)
}

// Gets the album's cover photo. Publicly available.
// Returns an HTTP 302 URL string with the location set to the picture URL.
func (a *Album) GetPicture() (pic *Picture, err os.Error) {
	if a.picture == "" {
		err = os.NewError("Error: Album.GetPicture: The picture URL is empty.")
	}
	return NewPicture(a.picture), err
}

func getAlbums(url string) (as []Album, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
	as = make([]Album, len(data))
	for i, v := range data {
		as[i], err = parseAlbum(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Parses Album data. Returns nil for err if no error appeared.
func parseAlbum(value map[string]interface{}) (a Album, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			a.ID = val.(string)
		case "from":
			a.From = parseObject(val.(map[string]interface{}))
		case "name":
			a.Name = val.(string)
		case "description":
			a.Description = val.(string)
		case "location":
			a.Location = val.(string)
		case "link":
			a.Link = val.(string)
		case "privacy":
			a.Privacy = val.(string)
		case "count":
			a.Count = val.(float64)
		case "created_time":
			a.CreatedTime, err = parseTime(val.(string))
		case "updated_time":
			a.UpdatedTime, err = parseTime(val.(string))
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, va := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "photos":
					a.photos = va.(string)
				case "comments":
					a.comments = va.(string)
				case "picture":
					a.picture = va.(string)
				}
			}
		}
	}
	return
}
