package graph

import (
	"os"
	"log"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

type Graph struct {
	Logger *log.Logger
	AccesToken string
}

func NewGraph(logger *log.Logger, accessToken string) (g *Graph) {
	g = new(Graph)
	g.Logger = logger
	g.AccesToken = accessToken
	return
}

// Fetches the Group with the provided ID.
func (g *Graph) FetchGroup(id string) (gr *Group, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	group, err := parseGroup(data)
	return &group, err
}
// Fetches the Event with the provided ID.
func (g *Graph) FetchEvent(id string) (e *Event, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	ev, err := parseEvent(data)
	return &ev, err
}
// Fetches Events from an URL.
func (g *Graph) FetchEvents(url string) (es *[]Event, err os.Error) {
	e, err := getEvents(url)
	return &e, err
}

// Fetches the Application with the provided ID.
func (g *Graph) FetchApplication(id string) (app *Application, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	applic, err := parseApplication(data)
	return &applic, err
}
// Fetches the Post with the provided ID.
func (g *Graph) FetchPost(id string) (p *Post, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	post, err := parsePost(data)
	return &post, err
}
/* Fetches posts from an facebook GraphAPI URL.
 * Returns err is nil if no error appeared.
 */
func (g *Graph) FetchPosts(url string) (*[]Post,os.Error) {
	var err os.Error
	post, err := fetchPosts(url)
	return &post, err
}
// Fetches the Insights with the provided ID.
func (g *Graph) FetchInsights(id string) (i *Insights, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	in, err := parseInsights(data["data"].([]interface{}))
	return &in, err
}
// Fetches the Note with the provided ID.
func (g *Graph) FetchNote(id string) (n *Note, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	note, err := parseNote(data)
	return &note,err
}
// Fetches the StatusMessage with the provided ID.
func (g *Graph) FetchStatusMessage(id string) (sm *StatusMessage, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	stm, err := parseStatusMessage(data)
	return &stm, err
}
// Fetches the Photo with the provided ID.
func (g *Graph) FetchPhoto(id string) (p *Photo, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	photo, err := parsePhoto(data)
	return &photo, err
}
// Fetches the Page with the provided ID.
func (g *Graph) FetchPage(id string) (p *Page, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	page, err := parsePage(data)
	return &page, err
}
// Fetches the User with the provided ID or name.
func (g *Graph) FetchUser(name string) (u *User, err os.Error) {
	data, err := getObject(name)
	if err != nil {
		return
	}
	var user User
	user, err = parseUser(data)
	return &user, err
}
// Fetches the Video with the provided ID or name.
func (g *Graph) FetchVideo(id string) (v *Video, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	vid, err := parseVideo(data)
	return &vid, err
}
// Fetches the Album with the provided ID or name.
func (g *Graph) FetchAlbum(id string) (a *Album, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	album, err := parseAlbum(data)
	return &album, err
}
// Fetches the Link with the provided ID or name.
func (g *Graph) FetchLink(id string) (l *Link, err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	link, err := parseLink(data)
	return &link, err
}
// Fetches the Checkin with the provided ID or name.
func (g *Graph) FetchCheckin(id string) (c *Checkin,err os.Error) {
	data, err := getObject(id)
	if err != nil {
		return
	}
	checkin, err := parseCheckin(data)
	return &checkin, err
}
