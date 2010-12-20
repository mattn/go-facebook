package graph

import (
	"testing"
)

// Tests

type PageTest struct {
	ID   string
	Name string
}

type UserTest struct {
	ID   string
	Name string
}

type PostTest struct {
	ID      string
	Message string
}

type GroupTest struct {
	ID string
}

type EventTest struct {
	ID   string
	Name string
}

var PageTests = []PageTest{
	{"19292868552", "Facebook Platform"},
	{"20531316728", "Facebook"},
	{"40796308305", "Coca-Cola"},
}

var UserTests = []UserTest{
	{"220439", "Bret Taylor"},
}

var PostTests = []PostTest{
	{"19292868552_118464504835613", "We're getting ready for f8! Check out the latest on the f8 Page, including a video from the first event, when Platform launched :: http://bit.ly/ahHl7j"},
}

var GroupTests = []GroupTest{
	{"2204501798"},
}

var EventTests = []EventTest{
	{"331218348435", "Facebook Developer Garage Austin - SXSW Edition"},
}

var g = NewGraph("")

func TestGroups(t *testing.T) {
	for _, v := range GroupTests {
		t.Logf("Fetching facebook group %s\n", v.ID)
		_, err := g.FetchGroup(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
	}
}

func TestEvents(t *testing.T) {
	for _, v := range EventTests {
		t.Logf("Fetching facebook event %s\n", v.ID)
		e, err := g.FetchEvent(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if e.Name != v.Name {
			t.Errorf("Error: %s expected %s \n", e.Name, v.Name)
		}
	}
}

func TestPages(t *testing.T) {
	for _, v := range PageTests {
		t.Logf("Fetching facebook page %s\n", v.ID)
		p, err := g.FetchPage(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if p.Name != v.Name {
			t.Errorf("Error: %s expected %s \n", p.Name, v.Name)
		}

		// Continue with connections

		// Check Wall
		t.Logf("Fetching facebook Page's Wall.\n")
		_, err = p.GetWall()
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}

		// TODO: Test different types of posts

		// Picture
		t.Logf("Fetching facebook Page's Picture.\n")
		pic, err := p.GetPicture()
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if len(pic.URL) == 0 {
			t.Errorf("Error: Page.GetPicture.URL is empty.")
		}

		// Tagged
		t.Logf("Fetching facebook Page's Tagged.\n")
		tags, err := p.GetTagged()
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		for _, v := range tags {
			switch v {
			case nil:
				t.Errorf("Error: Page.GetTagged[i] is empty.")
			}
		}

		// Links, requires access token
		/*
			t.Logf("Fetching facebook Page's Links.\n")
			links, err := p.GetLinks()
			if err != nil {
				t.Errorf("Error: %s\n", err.String())
			}
			for _, v := range links {
				if len(v.Name) == 0 {
					t.Errorf("Error: Page.GetLinks[i].Name is empty.")
				}
			}
		*/
		t.Logf("Fetching facebook Page's Photos.\n")
		photos, err := p.GetPhotos()
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		for _, v := range photos {
			if len(v.ID) == 0 {
				t.Errorf("Error: Page.GetPhotos[i].ID is empty")
			}
		}

		t.Logf("Fetching facebook Page's Albums.\n")
		albums, err := p.GetAlbums()
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		for _, v := range albums {
			if len(v.ID) == 0 {
				t.Errorf("Error: Page.GetAlbums[i].ID is empty")
			}
		}

		/* FIXME: Requires access token
		t.Logf("Fetching facebook Page's Statuses.\n")
		statuses, err := p.GetStatuses()
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		for _, v := range statuses {
			if len(v.ID) == 0 {
				t.Errorf("Error: Page.GetStatuses[i].ID is empty")
			}
		}
		*/
	}
}

func TestUsers(t *testing.T) {
	for _, v := range UserTests {
		t.Logf("Fetching facebook user %s\n", v.ID)
		u, err := g.FetchUser(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if u.Name != v.Name {
			t.Errorf("Error: %s expected %s \n", u.Name, v.Name)
		}
	}
}

func TestPosts(t *testing.T) {
	for _, v := range PostTests {
		t.Logf("Fetching facebook user %s\n", v.ID)
		_, err := g.FetchPost(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
	}
}


// Benchmarks
/*
func BenchmarkPage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPage("19292868552")
	}
}
*/
