package facebook

import (
	"strconv"
	"os"
	"time"
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
	Founded         *time.Time
	CompanyOverview string
	Mission         string
	Products        string
}

func (p *Page) String() string {
	return "ID: " + p.ID + "\tName: " + p.Name + "\tPicture: " + p.Picture +
		"\tLink: " + p.Link + "\tCategory: " + p.Category + "\tWebsite: " +
		p.Website + "\tUsername: " + p.Username + "\tFounded: " + p.Founded.String() +
		"\tCompany overview: " + p.CompanyOverview + "\tMission: " + p.Mission +
		"\tProducts: " + p.Products + "\tFan count:" +
		strconv.Ftoa64(p.FanCount, 'e', -1) + "\n"
}

func (p *Page) parseData(value map[string] interface{}) {
	for key, value := range data {
		switch key {
		case "website":
			p.Website = value.(string)
		case "picture":
			p.Picture = value.(string)
		case "fan_count":
			p.FanCount = value.(float64)
		case "mission":
			p.Mission = value.(string)
		case "category":
			p.Category = value.(string)
		case "name":
			p.Name = value.(string)
		case "username":
			p.Username = value.(string)
		case "link":
			p.Link = value.(string)
		case "id":
			p.ID = value.(string)
		case "products":
			p.Products = value.(string)
		case "founded":
			p.Founded, err = parseTime(value.(string))
		case "company_overview":
			p.CompanyOverview = value.(string)
		case "type":
			// TODO: Look into type
		case "metadata":
			// TODO: get and parse connections
		default:
			debugInterface(value, key, "Page")
		}
	}
}