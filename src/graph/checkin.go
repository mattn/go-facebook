package graph

import (
	"os"
	"time"
)

// You can search, read and publish checkins.
// Every checkin is associated with a checkin ID that represents an object in the graph.
// These checkins are then associated with locations represented by Facebook Pages or Open Graph protocol pages.
// To get a user's checkins, request the user_checkins extended permission. To see the user's friends' checkins, request the friends_checkins extended permission.
// NOTE: Requires user_checkins permission for all properties.
type Checkin struct {
	// The checkin ID.
	ID string
	// The ID and name of the user who made the checkin. Contains the name and Facebook id of the user who made the checkin.
	From Object
	// The users the author tagged in the checkin. Contains in data a list of the users tagged in this checkin.
	// Tags TODO
	// Information about the Facebook Page that represents the location of the checkin. Contains the Page id, name, and location.
	// Place TODO
	// The message the user added to the checkin.
	Message string
	// Information about the application that made the checkin. Object that contains the name and id of the application.
	Application Object
	// The time the checkin was created. Contains a IETF RFC 3339 datetime.
	CreatedTime *time.Time
}

func getCheckins(url string) (cs []Checkin, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
	cs = make([]Checkin, len(data))
	for i, v := range data {
		cs[i], err = parseCheckin(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Parses Checkin data. Returns nil for err if no error appeared.
func parseCheckin(value map[string]interface{}) (c Checkin, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			c.ID = val.(string)
		case "from":
			c.From = parseObject(val.(map[string]interface{}))
		case "tags":
			// TODO
		case "place":
			// TODO
		case "message":
			c.Message = val.(string)
		case "application":
			c.Application = parseObject(val.(map[string]interface{}))
		case "created_time":
			c.CreatedTime, err = parseTime(val.(string))
		}
	}
	return
}
