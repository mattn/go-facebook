package facebook

import (
	"os"
	"time"
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

type News struct {
	ID          string
	From        Object
	To          []Object
	Message     string
	Actions     []Link
	Type        string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Likes       string
	Comments    []Comment
}

func parseNews(value []interface{}) (news []News, err os.Error) {
	news = make([]News, len(value))
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
				news[i].Likes = val.(string)
			case "comments":
				data := val.(map[string]interface{})
				news[i].Comments = parseComments(data["data"].([]interface{}))
			}
		}
	}
	return
}

type Comment struct {
	ID          string
	From        Object
	Message     string
	CreatedTime string
}

func parseComment(value map[string]interface{}) (comment Comment) {
	comment.ID = value["id"].(string)
	comment.From = parseObject(value["from"].(map[string]interface{}))
	comment.Message = value["message"].(string)
	comment.CreatedTime = value["created_time"].(string)
	return
}

func parseComments(value []interface{}) (comments []Comment) {
	for i, v := range value {
		comments[i] = parseComment(v.(map[string]interface{}))
	}
	return
}
