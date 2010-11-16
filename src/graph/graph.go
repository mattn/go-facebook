package facebook

import (
  "os"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

type Graph struct {
	groups map[string]Group
	pages  map[string]Page
	events map[string]Event
}

/*
 * Fetches the Group with the provided ID.
 */
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

/*
 * Gets the Group with the provided ID.
 */
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


/*
 * Fetches the Event with the provided ID.
 */
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

/*
 * Fetches Events from an URL.
 */
func FetchEvents(url string) (es []Event, err os.Error) {
	return
}

/* 
 * Gets the Event with the provided ID.
 */
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
