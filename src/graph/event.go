package facebook

import (
	"os"
	"time"
)

// Specifies information about an event, including the location, event name, and which invitees plan to attend.
type Event struct {
	// The event ID
	ID string
	// An object containing the name and ID of the user who owns the event
	Owner Object
	// The event title
	Name string
	// The long-form HTML description of the event
	Description string
	// The start time of the event, an ISO-8601 formatted date/time
	StartTime *time.Time
	// The end time of the event, an ISO-8601 formatted date/time
	EndTime *time.Time
	// The location for this event.
	Location string
	// The location of this event, a structured address object with the properties street, city, state, zip, country, latitude, and longitude
	Venue Venue
	// The visibility of this event. Can be 'OPEN', 'CLOSED', or 'SECRET'
	Privacy string
	// The last time the event was updated
	UpdatedTime *time.Time
/*
	// Connections
	// This event's wall
	Feed []Post
	// All of the users who have been not yet responded to their invitation to this event
	NoReply []Invitation
	// All of the users who have been responded "Maybe" to their invitation to this event
	Maybe []Invitation
	// All of the users who have been invited to this event
	Invited []Invitation
	// All of the users who are attending this event
	Attending []Invitation
	// All of the users who declined their invitation to this event
	Declined []Invitation
	// The event's profile picture
	Picture *Picture
	*/
}

/*
 * Parses Event data. Returns nil for err if no error appeared.
 */
func parseEvent(value map[string]interface{}) (e Event, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			e.ID = val.(string)
		case "owner":
			e.Owner = parseObject(val.(map[string]interface{}))
		case "name":
			e.Name = val.(string)
		case "description":
			e.Description = val.(string)
		case "start_time":
			e.StartTime, err = parseTime(val.(string))
		case "end_time":
			e.EndTime, err = parseTime(val.(string))
		case "location":
			e.Location = val.(string)
		case "venue":
			e.Venue = parseVenue(val.(map[string]interface{}))
		case "privacy":
			e.Privacy = val.(string)
		case "updated_time":
			e.UpdatedTime, err = parseTime(val.(string))
		// Connections
			/*
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "feed":
					e.Feed, err = GetPosts(v.(string))
				case "noreply":
					e.NoReply, err = GetInvitations(v.(string))
				case "maybe":
					e.Maybe, err = GetInvitations(v.(string))
				case "invited":
					e.Invited, err = GetInvitations(v.(string))
				case "attending":
					e.Attending, err = GetInvitations(v.(string))
				case "declined":
					e.Declined, err = GetInvitations(v.(string))
				case "picture":
					e.Picture = NewPicture(v.(string))
				}
			}
			*/

		}
	}
	return
}
