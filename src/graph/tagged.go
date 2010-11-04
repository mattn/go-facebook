package facebook

import (
	"os"
	"time"
)

type Tagged struct {
	Posts   []News
	Paging Paging
}

func NewTagged() DataParser {
	var t Tagged
	return t
}

func (t Tagged) ParseData(value []interface{}) (err os.Error) {
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
			case "picture":
				news[i].Picture = NewPicture(val.(string))
			case "link":
				news[i].Link = v.(string) // Just an URL
			case "name":
				news[i].Name = v.(string)
			case "icon":
				news[i].Icon = v.(string) // Just an URL
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
	t.Posts = news
	return
}

func (t Tagged) SavePaging(p Paging) {
	t.Paging = p
}