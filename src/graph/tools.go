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
