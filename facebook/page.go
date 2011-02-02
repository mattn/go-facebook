package facebook

import (
	"os"
	"json"
)

// A Facebook Page.
// http://developers.facebook.com/docs/reference/api/page
type Page struct {
	// ID
	ID string
	// Name
	Name string
	// Category
	Category string
	// Number of users who like this page
	Likes float64
	// Metadata contains Connections
	Metadata *Metadata
}

func GetPage(id string) (page *Page, err os.Error) {
	resp, err := Call(id, RequestMetadata)
	if err != nil {
		return
	}
	var value Page
	err = json.Unmarshal(resp.Data, &value)
	page = &value
	return
}
