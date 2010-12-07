package graph

import (
	"os"
	"http"
	"io/ioutil"
	"json"
)

type Response struct {
	Data []interface{}
	Map  map[string]interface{}

	Url     string
	Fail    bool
	Message string
}

func GetResponse(url string) (r Response, err os.Error) {
	r.Url = url
	// Http
	resp, _, err := http.Get(r.Url)
	if err != nil {
		return r, os.NewError("GetResponse(" + url + "): " + err.String())
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// Json
	var values interface{}
	err = json.Unmarshal(b, &values)
	if err != nil {
		return
	}
	switch values.(type) {
	case bool:
		r.Fail = values.(bool)
		r.Message = "Call for Response couldn't be handled correctly."
	case map[string]interface{}:
		r.Fail = false
		data := values.(map[string]interface{})
		if val, ok := data["data"]; ok {
			r.Data = val.([]interface{})
		} else {
			r.Map = data
		}
	default:
		r.Fail = true
		r.Message = "Unsupported returned JSON-Data."
	}
	return
}
