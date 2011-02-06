package facebook

import (
	"os"
	"json"
)

// A photo album.
type Album struct {
	// The photo album ID. Publicly available.
	ID string
	// The profile that created this album. Publicly available. Contains the id and name fields.
	From *Object
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
	Created_Time string
	// The last time the photo album was updated. Publicly available. Contains a IETF RFC 3339 datetime.
	Updated_Time string
	// Metadata contains Connections
	*Metadata
}

func GetAlbum(id string) (album *Album, err os.Error) {
	resp, err := Call(id, RequestMetadata)
	if err != nil {
		return
	}
	var value Album
	err = json.Unmarshal(resp.Data, &value)
	album = &value
	return
}

func PostAlbum(profileID, name, description string) (err os.Error) {
	_, err = Publish(profileID, "albums", map[string]string{"name": name, "description": description})
	return
}
