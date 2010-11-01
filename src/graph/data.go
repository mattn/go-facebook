package facebook

import (
	"os"
)

type DataParser interface {
	ParseData(value []interface{}) os.Error
	SavePaging(p Paging)
}

func FetchData(parser DataParser, url string) (err os.Error) {
	body, err := fetchPage(url)
	if err != nil {
		return
	}
	d, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range d {
		switch key {
		case "data":
			err = parser.ParseData(value.([]interface{}))
		case "paging":
			parser.SavePaging(parsePaging(value.(map[string]interface{})))
		}
	}
	return
}
