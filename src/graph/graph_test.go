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

func TestUser(t *testing.T) {
	name := "btaylor"
	t.Logf("Fetching facebook user %s\n", name)
	_, err := FetchUser(name)
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
}

func TestUserIntrospect(t *testing.T) {
	id := "btaylor"
	t.Logf("Fetching and introspecting facebook user %s\n", id)
	_, err := FetchUserIntrospect(id)
	if err != nil {
		t.Errorf("Error: %s\n", err.String())
	}
}

// Benchmarks

func BenchmarkPage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPage("19292868552")
	}
}

func BenchmarkPageIntrospect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchPageIntrospect("19292868552")
	}
}

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchUser("btaylor")
	}
}

func BenchmarkUserIntrospect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FetchUserIntrospect("btaylor")
	}
}
