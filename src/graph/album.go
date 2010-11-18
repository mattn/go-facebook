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
	Count string
	// The time the photo album was initially created. Publicly available. Contains a IETF RFC 3339 datetime.
	CreatedTime *time.Time
	// The last time the photo album was updated. Publicly available. Contains a IETF RFC 3339 datetime.
	UpdatedTime *time.Time

	// TODO: Connections
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
			a.Count = val.(string)
		case "created_time":
			a.CreatedTime, err = parseTime(val.(string))
		case "updated_time":
			a.UpdatedTime, err = parseTime(val.(string))
		}
	}
	return
}
