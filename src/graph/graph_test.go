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

var GroupTests = []GroupTest{
	{"2204501798"},
}

func TestGraph(t *testing.T) {
	g := NewGraph()
	for _, v := range PageTests {
		t.Logf("Fetching facebook page %s\n", v.ID)
		err := g.FetchPage(v.ID)
		p := g.GetPage(v.ID)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if p.Name != v.Name {
			t.Errorf("Error: %s expected %s \n", p.Name, v.Name)
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
