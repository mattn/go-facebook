package graph

import (
	"os"
	"time"
)

// A link shared on a user's wall
type Link struct {
	// The link ID. Publicly accessible.
	ID string
	// The user that created the link. Publicly accessible. Contains the id and name field.
	From Object
	// The URL that was shared. Publicly accessible. Contains a valid URL
	Link string
	// The name of the link. Publicly accessible.
	Name string
	// The caption of the link (appears beneath the link name). Publicly accessible.
	Caption string
	// A description of the link (appears beneath the link caption). Publicly accessible.
	Description string
	// A URL to the link icon that Facebook displays in the news feed. Publicly accessible. Contains a valid URL
	Icon string
	// A URL to the thumbnail image used in the link post. Publicly accessible. Contains a valid URL.
	Picture string
	// The optional message from the user about this link. Publicly accessible.
	Message string
	// The time the message was published. Publicly accessible. Contains a IETF RFC 3339 datetime.
	CreatedTime *time.Time

	// TODO: Connections
}

func getLinks(url string) (ls []Link, err os.Error) {
	data, err := getData(url)
	if err != nil {
		return
	}
	ls = make([]Link, len(data))
	for i, v := range data {
		ls[i], err = parseLink(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Parses Link data. Returns nil for err if no error appeared.
func parseLink(value map[string]interface{}) (l Link, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			l.ID = val.(string)
		case "from":
			l.From = parseObject(val.(map[string]interface{}))
		case "link":
			l.Link = val.(string)
		case "name":
			l.Name = val.(string)
		case "caption":
			l.Caption = val.(string)
		case "description":
			l.Description = val.(string)
		case "icon":
			l.Icon = val.(string)
		case "picture":
			l.Picture = val.(string)
		case "message":
			l.Message = val.(string)
		case "created_time":
			l.CreatedTime, err = parseTime(val.(string))
		}
	}
	return
}
