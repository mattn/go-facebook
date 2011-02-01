package facebook

import (
	"testing"
)

var urls = []string{
	"https://graph.facebook.com/btaylor",
	"https://graph.facebook.com/cocacola",
	"https://graph.facebook.com/331218348435",
	"https://graph.facebook.com/195466193802264",
	"https://graph.facebook.com/2439131959",
	"https://graph.facebook.com/98423808305",
	"https://graph.facebook.com/99394368305",
}

var ids = []string{
	"btaylor",
	"cocacola",
	"331218348435",
	"195466193802264",
	"2439131959",
	"98423808305",
	"99394368305",
}

var users = []string{
	"btaylor",
}

var pages = []string{
	"platform",
	"cocacola",
}

func TestCall(t *testing.T) {
	for _, id := range ids {
		t.Logf("Fetching Facebook object %s.\n", id)
		resp, err := Call(id, map[string]string{})
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if resp == nil {
			t.Errorf("Error: The Response is empty.")
		} else if resp.Data == nil {
			t.Errorf("Error: Empty Response.Data.")
		}
	}
}
