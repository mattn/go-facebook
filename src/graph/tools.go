package graph

import (
	"http"
	"os"
	"time"
)

type Object struct {
	ID   string
	Name string
}

func parseObject(value map[string]interface{}) (obj Object) {
	obj.ID = value["id"].(string)
	obj.Name = value["name"].(string)
	return
}

func parseObjects(value []interface{}) (objs []Object) {
	objs = make([]Object, len(value))
	for i, v := range value {
		if v == nil {
			// TODO: Do something about this  
		} else {
			objs[i] = parseObject(v.(map[string]interface{}))
		}
	}
	return
}

func parseUnknownObject(value map[string]interface{}) (obj interface{}, err os.Error) {
	t, ok := value["type"].(string)
	if ok {
		switch t {
		case "user":
			obj, err = parseUser(value)
			if err == nil {
				return
			}
		case "application":
			obj, err = parseApplication(value)
			if err == nil {
				return
			}
		case "page":
			obj, err = parsePage(value)
			if err == nil {
				return
			}
		case "group":
			obj, err = parseGroup(value)
			if err == nil {
				return
			}
		case "photo":
			obj, err = parsePhoto(value)
			if err == nil {
				return
			}
		case "album":
			obj, err = parseAlbum(value)
			if err == nil {
				return
			}
		case "checkin":
			obj, err = parseCheckin(value)
			if err == nil {
				return
			}
		case "comment":
			obj, err = parseComment(value)
			if err == nil {
				return
			}
		case "event":
			obj, err = parseEvent(value)
			if err == nil {
				return
			}
		case "friendlist":
			// TODO
		case "insights":
			// TODO
		case "link":
			obj, err = parseLink(value)
			if err == nil {
				return
			}
		case "message":
			// TODO
		case "note":
			obj, err = parseNote(value)
			if err == nil {
				return
			}
		case "status": //StatusMessage
			obj, err = parseStatusMessage(value)
			if err == nil {
				return
			}
		case "subscription":
			obj, err = parseSubscription(value)
			if err == nil {
				return
			}
		case "thread":
			// TODO
		case "video":
			obj, err = parseVideo(value)
			if err == nil {
				return
			}
		case "post":
			obj, err = parsePost(value)
			if err == nil {
				return
			}
		}
	}
	return obj, os.NewError("No field type detected, can't figure out the object type currently. TODO: Try and error over all possible FB obj types.")
}

type URL struct {
	Name string
	URL  string
}

func parseURL(value map[string]interface{}) (url URL) {
	url.Name = value["name"].(string)
	url.URL = value["link"].(string)
	return
}

func parseURLs(value []interface{}) (urls []URL) {
	urls = make([]URL, len(value))
	for i, v := range value {
		urls[i] = parseURL(v.(map[string]interface{}))
	}
	return
}

var times = []string{"RFC3339", "2006-01-02T15:04:05+0000", "2006-01-02T15:04:05", "January 2, 2006", "Jan 2006"}

func parseTime(value string) (t *time.Time, err os.Error) {
	for _, v := range times {
		t, err = time.Parse(v, value)
		if err == nil {
			return
		}
	}
	return
}

func post(url string, data map[string]string) (err os.Error) {
	_, err = http.PostForm(url, data)
	return
}
