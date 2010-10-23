package facebook

import (
	"testing"
)

// Tests

func TestPage(t *testing.T) {
	id := "19292868552"
	t.Logf("Fetching facebook page %s\n", id)
	_, err := FetchPage(id)
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
}

func TestPageIntrospect(t *testing.T) {
	id := "19292868552"
	t.Logf("Fetching and introspecting facebook page %s\n", id)
	_, err := FetchPageIntrospect(id)
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
}

func TestPerson(t *testing.T) {
	name := "btaylor"
	t.Logf("Fetching facebook person %s\n", name)
	_, err := FetchPerson(name)
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
}

func TestPersonIntrospect(t *testing.T) {
	id := "btaylor"
	t.Logf("Fetching and introspecting facebook person %s\n", id)
	_, err := FetchPersonIntrospect(id)
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
}

// Benchmarks

func BenchmarkPage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPerson("19292868552")
	}
}

func BenchmarkPageIntrospect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPageIntrospect("19292868552")
	}
}

func BenchmarkPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPerson("btaylor")
	}
}

func BenchmarkPersonIntrospect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPersonIntrospect("btaylor")
	}
}
