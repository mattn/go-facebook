package facebook

import (
	"os"
)

type Invitation struct {
	Name string
	ID   string
	// rsvp_status, not_replied, attending, unsure or declined
	RSVPStatus string
}

func GetInvitations(URL string) (invs []Invitation, err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchPage(URL)
	m, err := getJsonMap(b)
	data, ok := m["data"].([]interface{})
	if !ok {
		err = os.NewError("GetInvitations: data could not be found.")
		return
	}
	for i, v := range data {
		invs[i].parseData(v.(map[string]interface{}))
	}
	return
}

func (i *Invitation) parseData(value map[string]interface{}) {
	i.Name = value["name"].(string)
	i.ID = value["id"].(string)
	i.RSVPStatus = value["rsvp_status"].(string)
	return
}
