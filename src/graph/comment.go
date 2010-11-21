package graph

import (
	"os"
)

type Comment struct {
	ID          string
	From        Object
	Message     string
	CreatedTime string
}

func getComments(url string) (comments []Comment, err os.Error) {
	data, err := getData(url)
	if err != nil {
		return
	}
	for i, val := range data {
		comments[i] = parseComment(val.(map[string]interface{}))
	}
	return
}

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
