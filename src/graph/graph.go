package facebook

import (
	"http"
	"io/ioutil"
	"json"
	"os"
	"strconv"
	"fmt"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

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
		case "type":
			// TODO: Look into type
		case "metadata":
			parseMetaData(value)
		default:
			debugInterface(value, key, "Page")
		}
	}
	return
}

func FetchUserIntrospect(name string) (user User, err os.Error) {
	return FetchUser(name + "?metadata=1")
}

func FetchUser(name string) (user User, err os.Error) {
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
		case "id":
			user.ID = value.(string)
		case "first_name":
			user.FirstName = value.(string)
		case "last_name":
			user.LastName = value.(string)
		case "name":
			user.Name = value.(string)
		case "link":
			user.Link = value.(string)
		case "about":
			user.About = value.(string)
		case "birthday":
			user.Birthday = value.(string)
		case "work":
			user.Work = value.(string)
		case "education":
			user.Education = value.(string)
		case "email":
			user.Email = value.(string)
		case "website":
			user.Website = value.(string)
		case "hometown":
			user.Hometown = value.(string)
		case "location":
			user.Location = value.(string)
		case "bio":
			user.Bio = value.(string)
		case "quotes":
			user.Quotes = value.(string)
		case "gender":
			user.Gender = value.(string)
		case "interested_in":
			user.InterestedIn = value.(string)
		case "meeting_for":
			user.MeetingFor = value.(string)
		case "relationship_status":
			user.RelationshipStatus = value.(string)
		case "religion":
			user.Religion = value.(string)
		case "political":
			user.Political = value.(string)
		case "verified":
			user.Verified = value.(string)
		case "significant_other":
			user.SignificantOther = value.(string)
		case "timezone":
			user.Timezone = value.(string)
			
		// Connections
		case "picture":
			user.Picture = NewPicture(value.(string))

		// Not documented in the API but streamed	
		case "locale":
			user.Locale = value.(string)
		case "mission":
			user.Mission = value.(string)
		case "category":
			user.Category = value.(string)
		case "username":
			user.Username = value.(string)
		case "products":
			user.Products = value.(string)
		case "founded":
			user.Founded = value.(string)
		case "company_overview":
			user.CompanyOverview = value.(string)
		case "fan_count":
			user.FanCount = value.(float64)
		case "type":
			// TODO: Look into type

			// Parse metadata if requested
		case "metadata":
			parseMetaData(value)
		default:
			debugInterface(value, key, "Person")
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
		default:
			debugInterface(v, key, "Metadata")
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
		case "home":
			connections.Home = v.(string)
		case "friends":
			connections.Home = v.(string)
		case "activities":
			connections.Activities = v.(string)
		case "interests":
			connections.Interests = v.(string)
		case "music":
			connections.Music = v.(string)
		case "books":
			connections.Books = v.(string)
		case "movies":
			connections.Movies = v.(string)
		case "television":
			connections.Television = v.(string)
		case "likes":
			connections.Likes = v.(string)
		case "groups":
			connections.Groups = v.(string)
		case "inbox":
			connections.InBox = v.(string)
		case "outbox":
			connections.OutBox = v.(string)
		case "updates":
			connections.Updates = v.(string)
		case "accounts":
			connections.Accounts = v.(string)
		case "checkins":
			connections.CheckIns = v.(string)
		case "picture":
			connections.Picture = v.(string)
		case "family":
			connections.Family = v.(string)
		default:
			debugInterface(v, key, "Connections")
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

func debugInterface(value interface{}, key, funcName string) {
	var str string
	switch value.(type) {
	case float64:
		str = strconv.Ftoa64(value.(float64), 'e', -1)
	}
	fmt.Printf("%s: Unknown pair: %s : %s\n", funcName, key, str)
}
