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
	// Friends
	friends, err := m.GetFriends()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, obj := range friends.Data {
		ObjectTest(obj, t)
	}

	// Activities
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

	// Music
	music, err := m.GetMusic()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, item := range music.Data {
		ItemTest(m, item, t)
	}

	// Books
	books, err := m.GetBooks()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, item := range books.Data {
		ItemTest(m, item, t)
	}

	// Movies
	movies, err := m.GetMovies()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, item := range movies.Data {
		ItemTest(m, item, t)
	}

	// Television
	television, err := m.GetTelevision()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, item := range television.Data {
		ItemTest(m, item, t)
	}

	// Likes
	likes, err := m.GetLikes()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	for _, item := range likes.Data {
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

func ObjectTest(o *Object, t *testing.T) {
	if o == nil {
		t.Errorf("ObjectTest: Object %x is nil.\n", o)
	}
	if len(o.ID) == 0 {
		t.Errorf("Object.ID is empty.n")
	}
	if len(o.Name) == 0 {
		t.Errorf("Object.Name is empty.n")
	}
}
