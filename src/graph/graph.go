package facebook

import (
	"os"
)

const (
	GRAPHURL = "http://graph.facebook.com/"
)

type Graph struct {
	groups map[string]Group
	events map[string]Event
	pages  map[string]Page
}

// ### Groups ###

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

// ### Events ###

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

// ### Pages ###

/*
 * Fetches the Page with the provided ID.
 */
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

/* 
 * Gets the Page with the provided ID.
 */
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
