package graph

import (
	"os"
	"time"
)

// A Facebook note
type Note struct {
	// The note ID. Available to everyone on Facebook by default.
	ID string
	// The profile that created the note. Available to everyone on Facebook by default.
	From Object
	// The title of the note. Available to everyone on Facebook by default.
	Subject string
	// The content of the note. Available to everyone on Facebook by default. Contains HTML text.
	Message string
	// The time the note was initially published. Available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime
	CreatedTime *time.Time
	// The time the note was last updated. Available to everyone on Facebook by default. Contains a IETF RFC 3339 datetime
	UpdatedTime *time.Time
	// The icon that Facebook displays with notes. Available to everyone on Facebook by default. Contains a valid URL
	Icon string

	// Connections
	comments string
	likes    string
}

// Gets all of the comments on this Note. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (n *Note) GetComments() (cs []Comment, err os.Error) {
	if n.comments == "" {
		err = os.NewError("Error: Note.GetComments: The comments URL is empty.")
	}
	return getComments(n.comments)
}

// Gets users who like the note. Available to everyone on Facebook.
// Returns an array of objects containing the id and name fields.
func (n *Note) GetLikes() (likes []Object, err os.Error) {
	if n.likes == "" {
		err = os.NewError("Error: Note.GetLikes: The likes URL is empty.")
	}
	data, err := getData(n.likes)
	if err != nil {
		return
	}
	likes = parseObjects(data)
	return
}

func getNotes(url string) (ns []Note, err os.Error) {
	data, err := getData(url)
	if err != nil {
		return
	}
	ns = make([]Note, len(data))
	for i, v := range data {
		ns[i], err = parseNote(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

/*
 * Parses Note data. Returns nil for err if no error appeared.
 */
func parseNote(value map[string]interface{}) (n Note, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			n.ID = val.(string)
		case "from":
			n.From = parseObject(val.(map[string]interface{}))
		case "subject":
			n.Subject = val.(string)
		case "message":
			n.Message = val.(string)
		case "created_time":
			n.CreatedTime, err = parseTime(val.(string))
		case "updated_time":
			n.UpdatedTime, err = parseTime(val.(string))
		case "icon":
			n.Icon = val.(string)
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "comments":
					n.comments = v.(string)
				case "likes":
					n.likes = v.(string)
				}
			}
		}
	}
	return
}
