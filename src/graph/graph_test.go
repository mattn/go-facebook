package facebook

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

var PageTests = []PageTest{
	{"19292868552", "Facebook Platform"},
	{"20531316728", "Facebook"},
	{"40796308305", "Coca-Cola"},
}

var UserTests = []UserTest{
	{"220439", "btaylor"},
}

var PostTests = []PostTest{
	{"19292868552_118464504835613", "We're getting ready for f8! Check out the latest on the f8 Page, including a video from the first event, when Platform launched :: http://bit.ly/ahHl7j"},
}

func TestPage(t *testing.T) {
	for _, v := range PageTests {
		t.Logf("Fetching facebook page %s\n", v.ID)
		p, err := FetchPage(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if p.Name != v.Name {
			t.Errorf("Error: %s expected %s \n", p.Name, v.Name)
		}
	}
}

func TestUser(t *testing.T) {
	for _, v := range UserTests {
		t.Logf("Fetching facebook user %s\n", v.Name)
		u, err := FetchUser(v.Name)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if u.ID != v.ID {
			t.Errorf("Error: %s expected %s \n", u.ID, v.ID)
		}
	}
}

func TestPosts(t *testing.T) {
	for _, v := range PostTests {
		t.Logf("Fetching facebook post %s\n", v.ID)
		u, err := GetPost(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if u.ID != v.ID {
			t.Errorf("Error: %s expected %s \n", u.ID, v.ID)
		}
		if u.Message != v.Message {
			t.Errorf("Error: %s expected %s \n", u.Message, u.Message)
		}
	}
}

// Benchmarks

func BenchmarkPage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPage("19292868552")
	}
}

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchUser("btaylor")
	}
}

func BenchmarkPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPost("19292868552_118464504835613")
	}
}
