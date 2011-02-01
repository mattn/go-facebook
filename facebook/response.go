package facebook

import (
	"os"
	"http"
	"io/ioutil"
)

type Response struct {
	Data     string
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
	r.Data = string(b)
	return
}
