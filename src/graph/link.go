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

	// Connections
	comments string
}

// Gets all of the comments on this link. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (l *Link) GetComments() (cs []Comment, err os.Error) {
	if l.comments == "" {
		err = os.NewError("Error: Link.GetComments: The comments URL is empty.")
	}
	return getComments(l.comments)
}

// Requires access token
func getLinks(url string) (ls []Link, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
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
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, va := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "comments":
					l.comments = va.(string)
				}
			}
		}
	}
	return
}
