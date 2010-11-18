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

	// TODO: Connections
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
		}
	}
	return
}
