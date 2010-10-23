package facebook

import (
	"http"
	"io/ioutil"
	"json"
	"os"
	"strconv"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

type Metadata struct {
	Connections Connections
	// Name, Description
	Fields map[string]string
}

type Connections struct {
	Albums   string
	Statuses string
	Links    string
	Posts    string
	Notes    string
	Videos   string
	Feed     string
	Photos   string
	Tagged   string
	Events   string
}

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

type Person struct {
	ID          string
	Name        string
	FirstName   string
	LastName    string
	Link        string
	Gender      string
	Locale      string
	UpdatedTime string
}

func (p *Person) String() string {
	return "ID: " + p.ID + "\tName: " + p.Name + "\tFirst name: " + p.FirstName +
		"\tLast name: " + p.LastName + "\tLink: " + p.Link + "\tGender: " +
		p.Gender + "\tLocale: " + p.Locale + "\tUpdated time: " + p.UpdatedTime +
		"\n"
}

func FetchPageIntrospect(id string) (page Page, err os.Error) {
	return FetchPage(id + "?metadata=1")
}

func FetchPage(id string) (page Page, err os.Error) {
	body, err := fetchBody(id)
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
		case "metadata":
			parseMetaData(value)
		}
	}
	return
}

func FetchPersonIntrospect(name string) (person Person, err os.Error) {
	return FetchPerson(name + "?metadata=1")
}

func FetchPerson(name string) (person Person, err os.Error) {
	body, err := fetchBody(name)
	if err != nil {
		return
	}
	data, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range data {
		switch key {
		case "locale":
			person.Locale = value.(string)
		case "name":
			person.Name = value.(string)
		case "link":
			person.Link = value.(string)
		case "Gender":
			person.Gender = value.(string)
		case "first_name":
			person.FirstName = value.(string)
		case "last_name":
			person.LastName = value.(string)
		case "id":
			person.ID = value.(string)
		case "metadata":
			parseMetaData(value)
		}
	}
	return
}

func parseMetaData(value interface{}) (metadata Metadata) {
	data := value.(map[string]interface{})
	for key, v := range data {
		switch key {
		case "connections":
			metadata.Connections = parseConnections(v)
		case "fields":
			metadata.Fields = parseFields(v)
		}
	}
	return
}

func parseConnections(value interface{}) (connections Connections) {
	data := value.(map[string]interface{})
	for key, v := range data {
		switch key {
		case "feed":
			connections.Feed = v.(string)
		case "posts":
			connections.Posts = v.(string)
		case "tagged":
			connections.Tagged = v.(string)
		case "statuses":
			connections.Statuses = v.(string)
		case "links":
			connections.Links = v.(string)
		case "notes":
			connections.Notes = v.(string)
		case "photos":
			connections.Photos = v.(string)
		case "albums":
			connections.Albums = v.(string)
		case "events":
			connections.Events = v.(string)
		case "videos":
			connections.Videos = v.(string)
		}
	}
	return
}

func parseFields(value interface{}) (fields map[string]string) {
	fields = make(map[string]string)
	var field map[string]interface{}
	data := value.([]interface{})
	for _, c := range data {
		field = c.(map[string]interface{})
		fields[field["name"].(string)] = field["description"].(string)
	}
	return
}

func getJsonMap(body []byte) (data map[string]interface{}, err os.Error) {
	var values interface{}

	if err = json.Unmarshal(body, &values); err != nil {
		return
	}
	data = values.(map[string]interface{})
	return
}

func fetchBody(method string) (body []byte, err os.Error) {
	resp, _, err := http.Get(GRAPHURL + method) // Response, final URL, error
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
