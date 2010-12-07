package graph

import (
	"os"
	"time"
)

// A status message on a user's wall
type StatusMessage struct {
	// The status message ID. Publicly available to everyone on Facebook by default.
	ID string
	// The user who posted the message. Publicly available to everyone on Facebook by default. Contains id and name fields
	From Object
	// The status message content. Publicly available to everyone on Facebook by default.
	Message string
	// The time the message was published. Publicly available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime.
	UpdatedTime *time.Time

	// Connections
	comments string
	likes    string
}

func getStatusMessages(url string) (sms []StatusMessage, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail{
		return
	}
	data := resp.Data
	sms = make([]StatusMessage, len(data))
	for i, v := range data {
		sms[i], err = parseStatusMessage(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Gets all of the comments on this Message. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (m *StatusMessage) GetComments() (cs []Comment, err os.Error) {
	if m.comments == "" {
		err = os.NewError("Error: Message.GetComments: The comments URL is empty.")
	}
	return getComments(m.comments)
}

// Gets users who like the message. Available to everyone on Facebook.
// Returns an array of objects containing the id and name fields.
func (m *StatusMessage) GetLikes() (likes []Object, err os.Error) {
	if m.likes == "" {
		err = os.NewError("Error: Message.GetLikes: The likes URL is empty.")
	}
	resp, err := GetResponse(m.likes)
	if err != nil || resp.Fail {
		return
	}
	likes = parseObjects(resp.Data)
	return
}

/*
 * Parses StatusMessage data. Returns nil for err if no error appeared.
 */
func parseStatusMessage(value map[string]interface{}) (sm StatusMessage, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			sm.ID = val.(string)
		case "from":
			sm.From = parseObject(val.(map[string]interface{}))
		case "message":
			sm.Message = val.(string)
		case "updated_time":
			sm.UpdatedTime, err = parseTime(val.(string))
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "comments":
					sm.comments = v.(string)
				case "likes":
					sm.likes = v.(string)
				}
			}
		}
	}
	return
}
