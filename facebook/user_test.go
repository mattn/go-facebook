package facebook

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	for _, id := range users {
		user, err := GetUser(id)
		if err != nil {
			t.Fatalf("Error with user %s: %s\n", id, err)
		}
		if user == nil {
			t.Fatalf("User is nil of object: %x\n", user)
		}
		if len(user.ID) == 0 {
			t.Errorf("User.ID is empty of object: %x\n", user)
		}
		MetadataTest(user.Metadata, t)
	}
}

func MetadataTest(m *Metadata, t *testing.T) {
	if m == nil {
		t.Fatalf("No metadata included.\n")
	}
	t.Logf("Metadata test of type %s.", m.Type)

	// Friends
	friends, err := m.GetFriends()
	if err != nil {
		if err != NOFRIENDSURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, obj := range friends.Objects() {
			ObjectTest(obj, t)
		}
	}

	// Activities
	activities, err := m.GetActivities()
	if err != nil {
		if err != NOACTIVITIESURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, item := range activities.Items() {
			ItemTest(m, item, t)
		}
	}

	//Interests
	interests, err := m.GetInterests()
	if err != nil {
		if err != NOINTERESTSURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, item := range interests.Items() {
			ItemTest(m, item, t)
		}
	}

	// Music
	music, err := m.GetMusic()
	if err != nil {
		if err != NOMUSICURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, item := range music.Items() {
			ItemTest(m, item, t)
		}
	}

	// Books
	books, err := m.GetBooks()
	if err != nil {
		if err != NOBOOKSURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, item := range books.Items() {
			ItemTest(m, item, t)
		}
	}

	// Movies
	movies, err := m.GetMovies()
	if err != nil {
		if err != NOMOVIESURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, item := range movies.Items() {
			ItemTest(m, item, t)
		}
	}

	// Television
	television, err := m.GetTelevision()
	if err != nil {
		if err != NOTELEVISIONURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		for _, item := range television.Items() {
			ItemTest(m, item, t)
		}
	}

	// Likes
	likes, err := m.GetLikes()
	if err != nil {
		if err != NOLIKESURLERR {
			t.Errorf("%s\n", err)
		} else {
			t.Logf("%s\n", err)
		}
	} else {
		if likes == nil {
			t.Fatalf("Likes is nil and err is nil, this can't happen.")
		}
		switch likes.(type) {
		case PostLikes:
			l := likes.(PostLikes)
			for _, obj := range l.Objects() {
				ObjectTest(obj, t)
			}
		case Likes:
			l := likes.(Likes)
			for _, item := range l.Items() {
				ItemTest(m, item, t)
			}
		}
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
