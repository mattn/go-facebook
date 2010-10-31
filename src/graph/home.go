package facebook

import (
	"os"
)

type Home struct {
	News   []News
	Paging Paging
}

func FetchHome(name string) (home Home, err os.Error) {
	home, err = FetchHomeByURL(name + "/home")
	return
}

func FetchHomeByURL(url string) (home Home, err os.Error) {
	body, err := fetchPage(url)
	if err != nil {
		return
	}
	data, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range data {
		switch key {
		case "data":
			home.News, err = parseNews(value.([]interface{}))
		case "paging":
			home.Paging = parsePaging(value.(map[string]interface{}))
		}
	}
	return
}
