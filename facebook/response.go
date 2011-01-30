package facebook

import (
	"os"
	"http"
	"io/ioutil"
	"json"
)

type Response struct {
	Data interface{}

	Url      string
	FinalUrl string
}

func Get(url string) (r *Response, err os.Error) {
	r = &Response{Url: url}

	// Http
	resp, finalUrl, err := http.Get(r.Url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, os.NewError("GetResponse(" + url + "): " + err.String())
	}
	r.FinalUrl = finalUrl
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Json
	var values interface{}
	err = json.Unmarshal(b, &values)
	if err != nil {
		return nil, err
	}
	r.Data = values
	switch values.(type) {
	case bool:
		err = os.NewError("Call for Response couldn't be handled correctly.")
		// TODO: Extract error message
	case map[string]interface{}:
		// Do nothing
	default:
		err = os.NewError("Unsupported returned JSON-Data.")
	}
	return
}
