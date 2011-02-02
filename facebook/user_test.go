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
	as, err := m.GetActivities()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, a := range as.Data {
		if len(a.Id) == 0 {
			t.Errorf("Activity.ID is empty of metadata: %x\n", m)
		}
		if len(a.Name) == 0 {
			t.Errorf("Activity.Name is empty of metadata: %x\n", m)
		}
		if len(a.Category) == 0 {
			t.Errorf("Activity.Category is empty of metadata: %x\n", m)
		}
		if len(a.Created_Time) == 0 {
			t.Errorf("Activity.Created_Time is empty of metadata: %x\n", m)
		}
	}
}
