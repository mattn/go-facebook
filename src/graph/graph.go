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

func debugInterface(value interface{}, key, funcName string) {
	var str string
	switch value.(type) {
	case float64:
		str = strconv.Ftoa64(value.(float64), 'e', -1)
	}
	fmt.Printf("%s: Unknown pair: %s : %s\n", funcName, key, str)
}
