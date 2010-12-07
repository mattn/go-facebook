package graph

import (
	"os"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

type Graph struct {
	AccesToken string
}

func NewGraph(accessToken string) (g *Graph) {
	g = new(Graph)
	g.AccesToken = accessToken
	return
}

func (g *Graph) Call(method string) (resp Response, err os.Error) {
	url := GRAPHURL + method + "?metadata=1"
	if len(g.AccesToken) > 0 {
		url += "&access_token=" + g.AccesToken
	}
	resp, err = GetResponse(url)
	if resp.Fail {
		err = os.NewError("APICall to " + resp.Url + " failed with message: " + resp.Message + " and error: " + err.String())
	}
	return
}

// Fetches the Group with the provided ID.
func (g *Graph) FetchGroup(id string) (gr *Group, err os.Error) {
	resp, err := g.Call(id)
	group, err := parseGroup(resp.Map)
	return &group, err
}
// Fetches the Event with the provided ID.
func (g *Graph) FetchEvent(id string) (e *Event, err os.Error) {
	resp, err := g.Call(id)
	ev, err := parseEvent(resp.Map)
	return &ev, err
}
// Fetches Events from an URL.
func (g *Graph) FetchEvents(url string) (es *[]Event, err os.Error) {
	e, err := getEvents(url)
	return &e, err
}

// Fetches the Application with the provided ID.
func (g *Graph) FetchApplication(id string) (app *Application, err os.Error) {
	resp, err := g.Call(id)
	applic, err := parseApplication(resp.Map)
	return &applic, err
}
// Fetches the Post with the provided ID.
func (g *Graph) FetchPost(id string) (p *Post, err os.Error) {
	resp, err := g.Call(id)
	post, err := parsePost(resp.Map)
	return &post, err
}
/* Fetches posts from an facebook GraphAPI URL.
 * Returns err is nil if no error appeared.
 */
func (g *Graph) FetchPosts(url string) (*[]Post, os.Error) {
	var err os.Error
	post, err := fetchPosts(url)
	return &post, err
}
// Fetches the Insights with the provided ID.
func (g *Graph) FetchInsights(id string) (i *Insights, err os.Error) {
	resp, err := g.Call(id)
	in, err := parseInsights(resp.Data)
	return &in, err
}
// Fetches the Note with the provided ID.
func (g *Graph) FetchNote(id string) (n *Note, err os.Error) {
	resp, err := g.Call(id)
	note, err := parseNote(resp.Map)
	return &note, err
}
// Fetches the StatusMessage with the provided ID.
func (g *Graph) FetchStatusMessage(id string) (sm *StatusMessage, err os.Error) {
	resp, err := g.Call(id)
	stm, err := parseStatusMessage(resp.Map)
	return &stm, err
}
// Fetches the Photo with the provided ID.
func (g *Graph) FetchPhoto(id string) (p *Photo, err os.Error) {
	resp, err := g.Call(id)
	photo, err := parsePhoto(resp.Map)
	return &photo, err
}
// Fetches the Page with the provided ID.
func (g *Graph) FetchPage(id string) (p *Page, err os.Error) {
	resp, err := g.Call(id)
	page, err := parsePage(resp.Map)
	return &page, err
}
// Fetches the User with the provided ID or name.
func (g *Graph) FetchUser(name string) (u *User, err os.Error) {
	resp, err := g.Call(name)
	var user User
	user, err = parseUser(resp.Map)
	return &user, err
}
// Fetches the Video with the provided ID or name.
func (g *Graph) FetchVideo(id string) (v *Video, err os.Error) {
	resp, err := g.Call(id)
	vid, err := parseVideo(resp.Map)
	return &vid, err
}
// Fetches the Album with the provided ID or name.
func (g *Graph) FetchAlbum(id string) (a *Album, err os.Error) {
	resp, err := g.Call(id)
	album, err := parseAlbum(resp.Map)
	return &album, err
}
// Fetches the Link with the provided ID or name.
func (g *Graph) FetchLink(id string) (l *Link, err os.Error) {
	resp, err := g.Call(id)
	link, err := parseLink(resp.Map)
	return &link, err
}
// Fetches the Checkin with the provided ID or name.
func (g *Graph) FetchCheckin(id string) (c *Checkin, err os.Error) {
	resp, err := g.Call(id)
	checkin, err := parseCheckin(resp.Map)
	return &checkin, err
}
