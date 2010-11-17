package facebook

import (
	"os"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

type Graph struct {
	// Facebook GraphAPI objects
	// messages
	// threads
	groups       map[string]Group
	events       map[string]Event
	applications map[string]Application
	posts        map[string]Post
	photos       map[string]Photo
	insights     map[string]Insights
	notes        map[string]Note
	// subscriptions // NOTE: Unsure if this will be used
	// status messages
	pages map[string]Page
	users map[string]User
	// videos
	// albums
	// links
	// checkins
}

func NewGraph() (g *Graph) {
	g = new(Graph)

	g.groups = make(map[string]Group)
	g.events = make(map[string]Event)
	g.applications = make(map[string]Application)
	g.posts = make(map[string]Post)
	g.photos = make(map[string]Photo)
	g.insights = make(map[string]Insights)
	g.notes = make(map[string]Note)

	g.pages = make(map[string]Page)
	g.users = make(map[string]User)

	return
}

// ### Groups ###

// Fetches the Group with the provided ID.
func (g *Graph) FetchGroup(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1") // Get metadata
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.groups[id], err = parseGroup(data)
	return
}

// Gets the Group with the provided ID.
func (g *Graph) GetGroup(id string) *Group {
	gr, ok := g.groups[id]
	if ok {
		return &gr
	}
	err := g.FetchGroup(id)
	if err != nil {
		return nil
	}
	gr = g.groups[id]
	return &gr
}

// ### Events ###

// Fetches the Event with the provided ID.
func (g *Graph) FetchEvent(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.events[id], err = parseEvent(data)
	return
}

// Fetches Events from an URL.
func (g *Graph) FetchEvents(url string) (err os.Error) {
	body, err := fetchPage(url)
	if err != nil {
		return
	}
	d, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range d {
		switch key {
		case "data":
			data := value.([]interface{})
			for _, val := range data {
				var event Event
				event, err = parseEvent(val.(map[string]interface{}))
				g.events[event.ID] = event
			}
		case "paging":
		}
	}
	return
}

// Gets the Event with the provided ID.
func (g *Graph) GetEvent(id string) *Event {
	p, ok := g.events[id]
	if ok {
		return &p
	}
	err := g.FetchEvent(id)
	if err != nil {
		return nil
	}
	p = g.events[id]
	return &p
}

// ### Applications ###

// Fetches the Application with the provided ID.
func (g *Graph) FetchApplication(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.applications[id], err = parseApplication(data)
	return
}

// Gets the Application with the provided ID.
func (g *Graph) GetApplication(id string) *Application {
	a, ok := g.applications[id]
	if ok {
		return &a
	}
	err := g.FetchApplication(id)
	if err != nil {
		return nil
	}
	a = g.applications[id]
	return &a
}

// ### Posts ###

// Fetches the Post with the provided ID.
func (g *Graph) FetchPost(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.posts[id], err = parsePost(data)
	return
}

/* Fetches posts from an facebook GraphAPI URL.
 * At the moment url isn't checked.
 * Returns err is nil if no error appeared.
 */
func (g *Graph) FetchPosts(url string) (posts []Post, err os.Error) {
	body, err := fetchPage(url)
	if err != nil {
		return
	}
	d, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range d {
		switch key {
		case "data":
			data := value.([]interface{})
			for _, val := range data {
				var post Post
				post, err = parsePost(val.(map[string]interface{}))
				g.posts[post.ID] = post
			}
		case "paging":
		}
	}
	return
}

// Gets the Post with the provided ID.
func (g *Graph) GetPost(id string) *Post {
	p, ok := g.posts[id]
	if ok {
		return &p
	}
	err := g.FetchPost(id)
	if err != nil {
		return nil
	}
	p = g.posts[id]
	return &p
}

// ### Insights ###

// Fetches the Insights with the provided ID.
func (g *Graph) FetchInsights(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.insights[id], err = parseInsights(id, data["data"].([]interface{}))
	return
}

// Gets the Insights with the provided ID.
func (g *Graph) GetInsights(id string) *Insights {
	i, ok := g.insights[id]
	if ok {
		return &i
	}
	err := g.FetchInsights(id)
	if err != nil {
		return nil
	}
	i = g.insights[id]
	return &i
}

// ### Notes ###

// Fetches the Note with the provided ID.
func (g *Graph) FetchNote(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.notes[id], err = parseNote(data)
	return
}

// Gets the Note with the provided ID.
func (g *Graph) GetNote(id string) *Note {
	n, ok := g.notes[id]
	if ok {
		return &n
	}
	err := g.FetchNote(id)
	if err != nil {
		return nil
	}
	n = g.notes[id]
	return &n
}

// ### Photos ###

// Fetches the Photo with the provided ID.
func (g *Graph) FetchPhoto(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.photos[id], err = parsePhoto(data)
	return
}

// Gets the Photo with the provided ID.
func (g *Graph) GetPhoto(id string) *Photo {
	p, ok := g.photos[id]
	if ok {
		return &p
	}
	err := g.FetchPhoto(id)
	if err != nil {
		return nil
	}
	p = g.photos[id]
	return &p
}

// ### Pages ###

// Fetches the Page with the provided ID.
func (g *Graph) FetchPage(id string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(id + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	g.pages[id], err = parsePage(data)
	return
}

// Gets the Page with the provided ID.
func (g *Graph) GetPage(id string) *Page {
	p, ok := g.pages[id]
	if ok {
		return &p
	}
	err := g.FetchPage(id)
	if err != nil {
		return nil
	}
	p = g.pages[id]
	return &p
}

// ### Users ###

// Fetches the User with the provided ID or name.
func (g *Graph) FetchUser(name string) (err os.Error) {
	// TODO: Check for valid ID
	b, err := fetchBody(name + "?metadata=1")
	if err != nil {
		return
	}
	data, err := getJsonMap(b)
	if err != nil {
		return
	}
	var user User
	user, err = parseUser(data)
	if err != nil {
		return
	}
	g.users[user.ID] = user
	return
}

// Gets the User with the provided ID.
func (g *Graph) GetUser(id string) *User {
	u, ok := g.users[id]
	if ok {
		return &u
	}
	err := g.FetchUser(id)
	if err != nil {
		return nil
	}
	u = g.users[id]
	return &u
}
