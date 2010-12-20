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
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
	for i, val := range data {
		comment, err := parseComment(val.(map[string]interface{}))
		if err != nil {
			return
		}
		comments[i] = comment
	}
	return
}

func parseComment(value map[string]interface{}) (comment Comment, err os.Error) {
	for k, v := range value {
		switch k {
		case "id":
			comment.ID = v.(string)
		case "from":
			comment.From = parseObject(v.(map[string]interface{}))
		case "message":
			comment.Message = v.(string)
		case "created_time":
			comment.CreatedTime = v.(string)
		default:
			err = os.NewError("Unsupported field of type " + k + " with value " + v.(string))
			return
		}
	}
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
		comment, err := parseComment(v.(map[string]interface{}))
		if err != nil {
			return
		}
		comments[i] = comment
	}
	return
}
