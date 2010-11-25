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
	if len(value) <= 0 {
		return
	}
	count = value["count"].(float64)
	if count <= 0 {
		return
	}
	data, ok := value["data"].([]interface{})
	if !ok {
		return
	}
	comments = make([]Comment, len(data))
	for i, v := range data {
		comments[i] = parseComment(v.(map[string]interface{}))
	}
	return
}
