package facebook

import (
	"os"
	"json"
)

var NOPICTUREURLERR = os.NewError("No Metadata.Connections[picture].")
var NOFRIENDSURLERR = os.NewError("No Metadata.Connections[friends].")
var NOLIKESURLERR = os.NewError("No Metadata.Connections[likes].")
var NOTELEVISIONURLERR = os.NewError("No Metadata.Connections[television].")
var NOMOVIESURLERR = os.NewError("No Metadata.Connections[movies].")
var NOBOOKSURLERR = os.NewError("No Metadata.Connections[books].")
var NOMUSICURLERR = os.NewError("No Metadata.Connections[music].")
var NOINTERESTSURLERR = os.NewError("No Metadata.Connections[interests].")
var NOACTIVITIESURLERR = os.NewError("No Metadata.Connections[activities].")

type Metadata struct {
	Type        string
	Connections map[string]string
	Fields      []*Field
}

// size = small | normal | large
func (m *Metadata) GetPicture(size string) (url string, err os.Error) {
	url, ok := m.Connections["picture"]
	if !ok {
		return "", NOPICTUREURLERR
	}
	url += "?type=" + size
	resp, err := Get(url)
	if err != nil {
		return
	}
	return string(resp.Data), nil
}

func (m *Metadata) GetFriends() (f *Friends, err os.Error) {
	url, ok := m.Connections["friends"]
	if !ok {
		return nil, NOFRIENDSURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Friends
	err = json.Unmarshal(resp.Data, &value)
	f = &value
	return
}

// Returns ever PostLikes or Likes object
func (m *Metadata) GetLikes() (l interface{}, err os.Error) {
	url, ok := m.Connections["likes"]
	if !ok {
		return nil, NOLIKESURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	switch m.Type {
	case "post":
		var value PostLikes
		err = json.Unmarshal(resp.Data, &value)
		l = &value
	default:
		var value Likes
		err = json.Unmarshal(resp.Data, &value)
		l = &value
	}
	return
}

func (m *Metadata) GetTelevision() (t *Television, err os.Error) {
	url, ok := m.Connections["television"]
	if !ok {
		return nil, NOTELEVISIONURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Television
	err = json.Unmarshal(resp.Data, &value)
	t = &value
	return
}

func (m *Metadata) GetMovies() (movies *Movies, err os.Error) {
	url, ok := m.Connections["movies"]
	if !ok {
		return nil, NOMOVIESURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Movies
	err = json.Unmarshal(resp.Data, &value)
	movies = &value
	return
}

func (m *Metadata) GetBooks() (b *Books, err os.Error) {
	url, ok := m.Connections["books"]
	if !ok {
		return nil, NOBOOKSURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Books
	err = json.Unmarshal(resp.Data, &value)
	b = &value
	return
}

func (m *Metadata) GetMusic() (music *Music, err os.Error) {
	url, ok := m.Connections["music"]
	if !ok {
		return nil, NOMUSICURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Music
	err = json.Unmarshal(resp.Data, &value)
	music = &value
	return
}

func (m *Metadata) GetInterests() (i *Interests, err os.Error) {
	url, ok := m.Connections["interests"]
	if !ok {
		return nil, NOINTERESTSURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Interests
	err = json.Unmarshal(resp.Data, &value)
	i = &value
	return
}

func (m *Metadata) GetActivities() (a *Activities, err os.Error) {
	url, ok := m.Connections["activities"]
	if !ok {
		return nil, NOACTIVITIESURLERR
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Activities
	err = json.Unmarshal(resp.Data, &value)
	a = &value
	return
}

type Field struct {
	Name        string
	Description string
}
