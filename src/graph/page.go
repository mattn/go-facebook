package facebook

import (
	"strconv"
)

type Page struct {
	ID              string
	Name            string
	Picture         string
	Link            string
	Category        string
	Website         string
	Username        string
	Founded         string
	CompanyOverview string
	Mission         string
	Products        string
	FanCount        float64
}

func (p *Page) String() string {
	return "ID: " + p.ID + "\tName: " + p.Name + "\tPicture: " + p.Picture +
		"\tLink: " + p.Link + "\tCategory: " + p.Category + "\tWebsite: " +
		p.Website + "\tUsername: " + p.Username + "\tFounded: " + p.Founded +
		"\tCompany overview: " + p.CompanyOverview + "\tMission: " + p.Mission +
		"\tProducts: " + p.Products + "\tFan count:" +
		strconv.Ftoa64(p.FanCount, 'e', -1) + "\n"
}
