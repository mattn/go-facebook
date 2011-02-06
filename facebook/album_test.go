package facebook

import (
	"testing"
)

func TestGetAlbum(t *testing.T) {
	for _, id := range albums {
		album, err := GetAlbum(id)
		if err != nil {
			t.Fatalf("Album: %x\tError: %s\n", album, err)
		}
		if album == nil {
			t.Fatalf("Album is nil of object: %x\n", album)
		}
		if len(album.ID) == 0 {
			t.Errorf("Album.ID is empty of object: %x\n", album)
		}
	}
}
