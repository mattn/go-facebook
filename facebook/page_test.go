package facebook

import (
	"testing"
)

func TestGetPage(t *testing.T) {
	for _, id := range pages {
		page, err := GetPage(id)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		if page == nil {
			t.Fatalf("Page is nil of object: %x\n", page)
		}
		if len(page.ID) == 0 {
			t.Errorf("Page.ID is empty of object: %x\n", page)
		}
	}
}
