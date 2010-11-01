package facebook

import (
	"os"
	"time"
)

type Home struct {
	News   []News
	Paging Paging
}

func NewHome() DataParser {
	var h Home
	return h
}

func (h Home) ParseData(value []interface{}) (err os.Error) {
	news := make([]News, len(value))
	for i, v := range value {
		wp := v.(map[string]interface{})
		for key, val := range wp {
			switch key {
			case "id":
				news[i].ID = val.(string)
			case "from":
				news[i].From = parseObject(val.(map[string]interface{}))
			case "to":
				data := val.(map[string]interface{})
				news[i].To = parseObjects(data["data"].([]interface{}))
			case "message":
				news[i].Message = val.(string)
			case "actions":
				news[i].Actions = parseLinks(val.([]interface{}))
			case "type":
				news[i].Type = val.(string)
			case "created_time":
				news[i].CreatedTime, err = time.Parse("RFC3339", val.(string)) // Eg.: 2010-10-21T22:54:34+0000
			case "updated_time":
				news[i].UpdatedTime, err = time.Parse("RFC3339", val.(string))
			case "likes":
				news[i].Likes = val.(float64)
			case "comments":
				data := val.(map[string]interface{})
				news[i].Comments = parseComments(data["data"].([]interface{}))
			}
		}
	}
	h.News = news
	return
}

func (h Home) SavePaging(p Paging) {
	h.Paging = p
}
