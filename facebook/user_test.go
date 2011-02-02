package facebook

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	for _, id := range users {
		user, err := GetUser(id)
		if err != nil {
			t.Errorf("%s\n", err)
		}
		if user == nil {
			t.Errorf("User is nil of object: %x\n", user)
		}
		if len(user.ID) == 0 {
			t.Errorf("User.ID is empty of object: %x\n", user)
		}
		MetadataTest(user.Metadata, t)
	}
}

func MetadataTest(m *Metadata, t *testing.T) {
	if m == nil {
		t.Errorf("No metadata included.\n")
	}
	// Activities()
	_, err := m.GetActivities()
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
