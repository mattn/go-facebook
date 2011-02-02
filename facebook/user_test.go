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
	for _, item := range as.Data {
		ItemTest(m, item, t)
	}

	//Interests
	interests, err := m.GetInterests()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, item := range interests.Data {
		ItemTest(m, item, t)
	}
}

func ItemTest(m *Metadata, i *Item, t *testing.T) {
	if len(i.Id) == 0 {
		t.Errorf("Item.ID is empty of metadata: %x\n", m)
	}
	if len(i.Name) == 0 {
		t.Errorf("Item.Name is empty of metadata: %x\n", m)
	}
	if len(i.Category) == 0 {
		t.Errorf("Item.Category is empty of metadata: %x\n", m)
	}
	if len(i.Created_Time) == 0 {
		t.Errorf("Item.Created_Time is empty of metadata: %x\n", m)
	}
}
