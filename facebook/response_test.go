package facebook

import (
	"testing"
)

func TestGet(t *testing.T) {
	for _, url := range urls {
		t.Logf("Fetching Facebook object from %s url.\n", url)
		resp, err := Get(url)
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

/* TODO: Test PostForm
func TestPostForm(t *testing.T) {
	resp, err := PostForm("https://graph.facebook.com/me/albums", map[string]string{"name": "Test album"})
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
	if resp == nil {
		t.Errorf("Error: The Response is empty.")
	} else if resp.Data == nil {
		t.Errorf("Error: Empty Response.Data.")
	}
}
*/
