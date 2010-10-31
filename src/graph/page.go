package facebook

import (
	"strconv"
	"os"
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

func FetchPage(id string) (page Page, err os.Error) {
	body, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range data {
		switch key {
		case "website":
			page.Website = value.(string)
		case "picture":
			page.Picture = value.(string)
		case "fan_count":
			page.FanCount = value.(float64)
		case "mission":
			page.Mission = value.(string)
		case "category":
			page.Category = value.(string)
		case "name":
			page.Name = value.(string)
		case "username":
			page.Username = value.(string)
		case "link":
			page.Link = value.(string)
		case "id":
			page.ID = value.(string)
		case "products":
			page.Products = value.(string)
		case "founded":
			page.Founded = value.(string)
		case "company_overview":
			page.CompanyOverview = value.(string)
		case "type":
			// TODO: Look into type
		case "metadata":
			// TODO: get and parse connections
		default:
			debugInterface(value, key, "Page")
		}
	}
	return
}
