package graph

import (
	"os"
	"time"
)

const (
	GROUP_PRIVACY_OPEN   = "OPEN"
	GROUP_PRIVACY_CLOSED = "CLOSED"
	GROUP_PRIVACY_SECRET = "SECRET"
)

// A Facebook group.
type Group struct {
	// The group ID. Publicly accessible.
	ID string
	// The URL for the group's icon. Publicly accessible.
	Icon string
	// The profile that created this group. Publicly accessible. A JSON object containing the id and name fields
	Owner Object
	// The name of the group. Publicly accessible.
	Name string
	// A brief description of the group. Publicly accessible.
	Description string
	// The URL for the group's website. Publicly accessible.
	Link string
	// The privacy setting of the group. Publicly accessible. Contains 'OPEN', 'CLOSED', or 'SECRET'
	Privacy string
	// The last time the group was updated. Publicly accessible. Contains a IETF RFC 3339 datetime.
	UpdatedTime *time.Time

	// Connections
	feed    string
	members string
	picture string
}

// Gets group's wall. Publicly available.
func (g *Group) GetFeed() (p []Post, err os.Error) {
	if g.feed == "" {
		err = os.NewError("Error: Group.GetFeed: The Feed URL is empty")
		return
	}
	return fetchPosts(g.feed)
}

// Get all of the users who are members of this group. Publicly available. Returned object contains friend id and name fields.
func (g *Group) GetMembers() (obj []Object) {
	return
}

// The profile picture of this group. Publicly available. Returns a HTTP 302 with the URL of the user's profile picture
func (g *Group) GetPicture() (pic *Picture) {
	return
}

/*
 * Parses Group data. Returns nil for err if no error appeared.
 */
func parseGroup(value map[string]interface{}) (g Group, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			g.ID = val.(string)
		case "icon":
			g.Icon = val.(string)
		case "owner":
			g.Owner = parseObject(val.(map[string]interface{}))
		case "name":
			g.Name = val.(string)
		case "description":
			g.Description = val.(string)
		case "link":
			g.Link = val.(string)
		case "privacy":
			g.Privacy = val.(string)
		case "updated_time":
			g.UpdatedTime, err = parseTime(val.(string))
			// Connections
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "feed":
					g.feed = v.(string)
				case "members":
					g.members = v.(string)
				case "picture":
					g.picture = v.(string)
				}
			}
		}
	}
	return
}
