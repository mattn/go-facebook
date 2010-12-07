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
func parseTime(value string) (t *time.Time, err os.Error) {
	t, err = time.Parse("RFC3339", value)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05+0000", value)
		if err != nil {
			t, err = time.Parse("January 2, 2006", value)
			if err != nil {
				t, err = time.Parse("Jan 2006", value)
			}
		}
	}
	return
}

func post(url string, data map[string]string) (err os.Error) {
	_, err = http.PostForm(url, data)
	return
}
