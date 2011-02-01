package facebook

import (
	"testing"
)

func TestGetAlbum(t *testing.T) {
	for _, id := range albums {
		album, err := GetAlbum(id)
		if err != nil {
			t.Errorf("%s\n", err)
		}
		if album == nil {
			t.Errorf("Album is nil of object: %s\n", album)
		}
		if len(album.ID) == 0 {
			t.Errorf("Album.ID is empty of object: %s\n", album)
		}
	}
}
