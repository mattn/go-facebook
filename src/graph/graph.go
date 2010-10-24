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
