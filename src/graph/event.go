package graph

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

	// Connections

	feed      string
	noReply   string
	maybe     string
	invited   string
	attending string
	declined  string
	picture   string
}

func getEvents(url string) (es []Event, err os.Error) {
	data, err := getData(url)
	if err != nil {
		return
	}
	es = make([]Event, len(data))
	for i, v := range data {
		es[i], err = parseEvent(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Gets the event's wall. Returns an array of Post objects.
func (e *Event) GetFeed() (feed []Post, err os.Error) {
	if e.feed == "" {
		err = os.NewError("Error: Event.GetFeed: The feed URL is empty.")
	}
	return fetchPosts(e.feed)
}

// Gets all of the users who have not yet been responded to their invitation to this event.
// Returns an array of Invitation objects.
func (e *Event) GetNoReply() (invs []Invitation, err os.Error) {
	if e.noReply == "" {
		err = os.NewError("Error: Event.GetNoReply: The noreply URL is empty.")
	}
	return getInvitations(e.noReply)
}

// Gets all of the users who have been responded "Maybe" to their invitation to this event.
// Returns an array of Invitation objects.
func (e *Event) GetMaybe() (invs []Invitation, err os.Error) {
	if e.noReply == "" {
		err = os.NewError("Error: Event.GetMaybe: The maybe URL is empty.")
	}
	return getInvitations(e.maybe)
}

// Gets all of the users who have been invited to this event.
// Returns an array of Invitation objects.
func (e *Event) GetInvited() (invs []Invitation, err os.Error) {
	if e.noReply == "" {
		err = os.NewError("Error: Event.GetInvited: The invited URL is empty.")
	}
	return getInvitations(e.invited)
}

// Gets all of the users who are attending this event.
// Returns an array of Invitation objects.
func (e *Event) GetAttending() (invs []Invitation, err os.Error) {
	if e.noReply == "" {
		err = os.NewError("Error: Event.GetAttending: The attending URL is empty.")
	}
	return getInvitations(e.attending)
}

// Gets all of the users who declined their invitation to this event.
// Returns an array of Invitation objects.
func (e *Event) GetDeclined() (invs []Invitation, err os.Error) {
	if e.noReply == "" {
		err = os.NewError("Error: Event.GetDeclined: The declined URL is empty.")
	}
	return getInvitations(e.declined)
}

// The event's profile picture
func (e *Event) GetPicture() (pic *Picture, err os.Error) {
	if e.picture == "" {
		err = os.NewError("Error: Event.GetPicture: The picture URL is empty.")
	}
	return NewPicture(e.picture), err
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
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "feed":
					e.feed = v.(string)
				case "noreply":
					e.noReply = v.(string)
				case "maybe":
					e.maybe = v.(string)
				case "invited":
					e.invited = v.(string)
				case "attending":
					e.attending = v.(string)
				case "declined":
					e.declined = v.(string)
				case "picture":
					e.picture = v.(string)
				}
			}
		}
	}
	return
}
