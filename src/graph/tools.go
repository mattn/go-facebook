package graph

import (
	"http"
	"io/ioutil"
	"json"
	"os"
	"strconv"
	"fmt"
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

func getData(url string) (value []interface{}, err os.Error) {
	if len(url) == 0 {
		return value, os.NewError("getData: url is empty.")
	}
	resp, _, err := http.Get(url) // Response, final URL, error
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	value, ok := data["data"].([]interface{})
	if !ok {
		err = os.NewError("getData: Couldn't parse data,")
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

func getObject(id string) (data map[string]interface{}, err os.Error) {
	return getObjByURL(GRAPHURL + id + "?metadata=1")
}

func getObjByURL(url string) (data map[string]interface{}, err os.Error) {
	if url == "" {
		return data, os.NewError("getObjByURL: url is empty.")
	}
	resp, _, err := http.Get(url) // Response, final URL, error
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	data, err = getJsonMap(b)
	return
}

func getJsonMap(body []byte) (data map[string]interface{}, err os.Error) {
	var values interface{}
	if err = json.Unmarshal(body, &values); err != nil {
		return
	}
	switch values.(type) {
	case bool:
		return data, os.NewError("JsonMap is false. Probably wrong UID request.")
	case map[string]interface{}:
		data = values.(map[string]interface{})
	default:
		return data, os.NewError("Unsupported JSON-Data. Body: " + string(body))
	}
	if e, ok := data["error"]; ok == true {
		error := e.(map[string]interface{})
		t := error["type"].(string)
		message := error["message"].(string)
		err = os.NewError(t + ": " + message)
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

func debugInterface(value interface{}, key, funcName string) {
	var str string
	switch value.(type) {
	case float64:
		str = strconv.Ftoa64(value.(float64), 'e', -1)
	}
	fmt.Printf("%s: Unknown pair: %s : %s\n", funcName, key, str)
}
