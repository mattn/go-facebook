package facebook

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
		objs[i] = parseObject(v.(map[string]interface{}))
	}
	return
}

func getData(URL string) (value []interface{}, err os.Error) {
	b, err := fetchPage(URL)
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

type Link struct {
	Name string
	URL  string
}

func parseLink(value map[string]interface{}) (link Link) {
	link.Name = value["name"].(string)
	link.URL = value["link"].(string)
	return
}

func parseLinks(value []interface{}) (links []Link) {
	links = make([]Link, len(value))
	for i, v := range value {
		links[i] = parseLink(v.(map[string]interface{}))
	}
	return
}

func getJsonMap(body []byte) (data map[string]interface{}, err os.Error) {
	var values interface{}

	if err = json.Unmarshal(body, &values); err != nil {
		return
	}
	data = values.(map[string]interface{})
	if e, ok := data["error"]; ok == true {
		error := e.(map[string]interface{})
		t := error["type"].(string)
		message := error["message"].(string)
		err = os.NewError(t + ": " + message)
	}
	return
}

func fetchBody(method string) (body []byte, err os.Error) {
	body, err = fetchPage(GRAPHURL + method)
	return
}

func fetchPage(url string) (body []byte, err os.Error) {
	resp, _, err := http.Get(url) // Response, final URL, error
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
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

func debugInterface(value interface{}, key, funcName string) {
	var str string
	switch value.(type) {
	case float64:
		str = strconv.Ftoa64(value.(float64), 'e', -1)
	}
	fmt.Printf("%s: Unknown pair: %s : %s\n", funcName, key, str)
}
