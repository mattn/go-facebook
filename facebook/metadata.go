package facebook

import (
	"encoding/json"
	"errors"
)

var NOPICTUREURLERR = errors.New("No Metadata.Connections[picture].")
var NOFRIENDSURLERR = errors.New("No Metadata.Connections[friends].")
var NOLIKESURLERR = errors.New("No Metadata.Connections[likes].")
var NOTELEVISIONURLERR = errors.New("No Metadata.Connections[television].")
var NOMOVIESURLERR = errors.New("No Metadata.Connections[movies].")
var NOBOOKSURLERR = errors.New("No Metadata.Connections[books].")
var NOMUSICURLERR = errors.New("No Metadata.Connections[music].")
var NOINTERESTSURLERR = errors.New("No Metadata.Connections[interests].")
var NOACTIVITIESURLERR = errors.New("No Metadata.Connections[activities].")

type Metadata struct {
	Type        string
	Connections map[string]string
	Fields      []*Field
}

// size = small | normal | large
func (m *Metadata) GetPicture(size string) (url string, err error) {
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

func (m *Metadata) GetFriends() (f *Friends, err error) {
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
func (m *Metadata) GetLikes() (l interface{}, err error) {
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

func (m *Metadata) GetTelevision() (t *Television, err error) {
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

func (m *Metadata) GetMovies() (movies *Movies, err error) {
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

func (m *Metadata) GetBooks() (b *Books, err error) {
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

func (m *Metadata) GetMusic() (music *Music, err error) {
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

func (m *Metadata) GetInterests() (i *Interests, err error) {
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

func (m *Metadata) GetActivities() (a *Activities, err error) {
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
