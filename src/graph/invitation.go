package graph

import (
	"os"
)

type Invitation struct {
	Name string
	ID   string
	// rsvp_status, not_replied, attending, unsure or declined
	RSVPStatus string
}

func getInvitations(url string) (invs []Invitation, err os.Error) {
	data, err := getData(url)
	if err != nil {
		return
	}
	for i, v := range data {
		invs[i] = parseInvitation(v.(map[string]interface{}))
	}
	return
}

func parseInvitation(value map[string]interface{}) (invi Invitation) {
	invi.Name = value["name"].(string)
	invi.ID = value["id"].(string)
	invi.RSVPStatus = value["rsvp_status"].(string)
	return
}
