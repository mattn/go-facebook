package graph

import (
	"os"
)

// A subscription to an application to get real-time updates for an Graph object type.
// For more details, see the Real-time Overview.
type Subscription struct {
	// The object type to subscribe to.Available to everyone in Facebook by default. Contains code or permissions.
	Object string
	// The list of fields for the object type. Available to everyone in Facebook by default. Contains a comma-seperated list of field names.
	// Fields TODO
	// An endpoint on your domain which can handle the real-time notifications. Available to everyone in Facebook by default. Contains a valid URL.
	CallbackURL string
	// Whether or not the subscription is active or not. Available to everyone in Facebook by default.
	Active bool
}

func getSubscriptions(url string) (s []Subscription, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
	s = make([]Subscription, len(data))
	for i, v := range data {
		s[i], err = parseSubscription(v.(map[string]interface{}))
		if err != nil {
			return
		}
	}
	return
}

// Parses Subscription data. Returns nil for err if no error appeared.
func parseSubscription(value map[string]interface{}) (s Subscription, err os.Error) {
	for key, val := range value {
		switch key {
		case "object":
			s.Object = val.(string)
		case "fields":
			//
		case "callback_url":
			s.CallbackURL = val.(string)
		case "active":
			s.Active = val.(bool)
		}
	}
	return
}
