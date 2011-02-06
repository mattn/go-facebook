package facebook

import (
	"os"
	"http"
	"io/ioutil"
	"json"
)

type FBError struct {
	Error *Error
}

type Error struct {
	Type    string
	Message string
}

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
	if err != nil {
		return
	}
	// Check for error
	var value FBError
	err = json.Unmarshal(r.Data, &value)
	if err == nil {
		if value.Error != nil {
			err = os.NewError(value.Error.Type + ": " + value.Error.Message)
			return
		}
	}
	return r, nil // Dont return the check of an Error error
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
	if err != nil {
		return
	}
	// Check for error
	var value FBError
	err = json.Unmarshal(r.Data, &value)
	if err == nil {
		if value.Error != nil {
			err = os.NewError(value.Error.Type + ": " + value.Error.Message)
			return
		}
	}
	return r, nil // Dont return the check of an Error error
}
