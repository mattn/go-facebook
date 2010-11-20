package graph

import (
//	"os"
)

type Comment struct {
	ID          string
	From        Object
	Message     string
	CreatedTime string
}
/*
func GetComments(url string) (comments []Comment, paging Paging, err os.Error) {
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
			data := value.([]interface{})
			for i, val := range data {
				comments[i] = parseComment(val.(map[string]interface{}))
			}
		case "paging":
			paging = parsePaging(value.(map[string]interface{}))
		}
	}
	return
}
*/
func parseComment(value map[string]interface{}) (comment Comment) {
	comment.ID = value["id"].(string)
	comment.From = parseObject(value["from"].(map[string]interface{}))
	comment.Message = value["message"].(string)
	comment.CreatedTime = value["created_time"].(string)
	return
}

func parseComments(value map[string]interface{}) (comments []Comment, count float64) {
	count = value["count"].(float64)
	data := value["data"].([]interface{})
	comments = make([]Comment, int(count))
	for i, v := range data {
		comments[i] = parseComment(v.(map[string]interface{}))
	}
	return
}
