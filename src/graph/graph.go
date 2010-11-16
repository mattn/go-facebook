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
	err = g.groups[id].parseData(data)
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
	return &g.groups[id]
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
	err = g.groups[id].parseData(data)
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
	return &g.pages[id]
}
