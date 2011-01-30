package facebook

import (
	"testing"
)

var urls = []string{
	"https://graph.facebook.com/btaylor",
	"https://graph.facebook.com/cocacola",
	"https://graph.facebook.com/331218348435",
	"https://graph.facebook.com/195466193802264",
	"https://graph.facebook.com/2439131959",
	"https://graph.facebook.com/98423808305",
	"https://graph.facebook.com/99394368305",
}

func TestGet(t *testing.T) {
	for _, url := range urls {
		t.Logf("Fetching Facebook object from %s url.\n", url)
		resp, err := Get(url)
		if err != nil {
			t.Errorf("Error: %s\n", err.String())
		}
		if resp == nil {
			t.Errorf("Error: The Response is empty.")
		} else if resp.Data == nil {
			t.Errorf("Error: Empty Response.Data.")
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
