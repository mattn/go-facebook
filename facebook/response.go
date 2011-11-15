package facebook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FBError struct {
	Err *Error
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

func Get(requestUrl string) (r *Response, err error) {
	resp, err := http.Get(requestUrl)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, errors.New("Get(" + requestUrl + "): " + err.Error())
	}
	if loc, err := resp.Location(); err == nil {
		r = &Response{Url: requestUrl, FinalUrl: loc.String()}
	} else {
		r = &Response{Url: requestUrl, FinalUrl: requestUrl}
	}
	r.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// Check for error
	var value FBError
	err = json.Unmarshal(r.Data, &value)
	if err == nil {
		if value.Err != nil {
			err = errors.New(value.Err.Type + ": " + value.Err.Message)
			return
		}
	}
	return r, nil // Dont return the check of an Error error
}

func PostForm(requestUrl string, data map[string]string) (r *Response, err error) {
	var postData url.Values
	for k,v := range data {
		postData[k] = []string{v}
	}
	resp, err := http.PostForm(requestUrl, postData)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, errors.New("Post(" + requestUrl + "): " + err.Error())
	}
	if loc, err := resp.Location(); err == nil {
		r = &Response{Url: requestUrl, FinalUrl: loc.String()}
	} else {
		r = &Response{Url: requestUrl, FinalUrl: requestUrl}
	}
	r.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// Check for error
	var value FBError
	err = json.Unmarshal(r.Data, &value)
	if err == nil {
		if value.Err != nil {
			err = errors.New(value.Err.Type + ": " + value.Err.Message)
			return
		}
	}
	return r, nil // Dont return the check of an Error error
}
