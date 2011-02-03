package facebook

import (
	"os"
	"json"
)

type Metadata struct {
	Connections map[string]string
	Fields      []*Field
}

func (m *Metadata) GetFriends() (f *Friends, err os.Error) {
	url, ok := m.Connections["friends"]
	if !ok {
		return nil, os.NewError("No Metadata.Connections[friends].")
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

func (m *Metadata) GetLikes() (l *Likes, err os.Error) {
	url, ok := m.Connections["likes"]
	if !ok {
		return nil, os.NewError("No Metadata.Connections[likes].")
	}
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value Likes
	err = json.Unmarshal(resp.Data, &value)
	l = &value
	return
}

func (m *Metadata) GetTelevision() (t *Television, err os.Error) {
	url, ok := m.Connections["television"]
	if !ok {
		return nil, os.NewError("No Metadata.Connections[television].")
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
		return nil, os.NewError("No Metadata.Connections[movies].")
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
		return nil, os.NewError("No Metadata.Connections[books].")
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
		return nil, os.NewError("No Metadata.Connections[music].")
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
		return nil, os.NewError("No Metadata.Connections[interests].")
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
		return nil, os.NewError("No Metadata.Connections[activities].")
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
