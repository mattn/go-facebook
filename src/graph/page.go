package facebook

import (
	"strconv"
)

// A Facebook Page.
// http://developers.facebook.com/docs/reference/api/page
type Page struct {
	// ID
	ID string
	// Name
	Name string
	// Profile picture
	Picture string
	// Category
	Category string
	// Number of fans the page has
	FanCount float64

	// Undocumented properties but streamed
	Link            string
	Website         string
	Username        string
	Founded         string
	CompanyOverview string
	Mission         string
	Products        string
}

func (p *Page) String() string {
	return "ID: " + p.ID + "\tName: " + p.Name + "\tPicture: " + p.Picture +
		"\tLink: " + p.Link + "\tCategory: " + p.Category + "\tWebsite: " +
		p.Website + "\tUsername: " + p.Username + "\tFounded: " + p.Founded +
		"\tCompany overview: " + p.CompanyOverview + "\tMission: " + p.Mission +
		"\tProducts: " + p.Products + "\tFan count:" +
		strconv.Ftoa64(p.FanCount, 'e', -1) + "\n"
}
