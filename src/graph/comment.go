package facebook

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
