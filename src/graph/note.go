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

	// TODO: Connections
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
		}
	}
	return
}
