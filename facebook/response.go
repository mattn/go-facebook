package facebook

import (
	"os"
	"http"
	"io/ioutil"
)

type Response struct {
	Data     []byte
	Url      string
	FinalUrl string
}

func Get(url string) (r *Response, err os.Error) {
	resp, finalUrl, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, os.NewError("Get(" + url + "): " + err.String())
	}
	r = &Response{Url: url}
	r.FinalUrl = finalUrl
	r.Data, err = ioutil.ReadAll(resp.Body)
	return
}

func PostForm(url string, data map[string]string) (r *Response, err os.Error) {
	resp, err := http.PostForm(url, data)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, os.NewError("Post(" + url + "): " + err.String())
	}
	r = &Response{Url: url, FinalUrl: url}
	r.Data, err = ioutil.ReadAll(resp.Body)
	return
}
