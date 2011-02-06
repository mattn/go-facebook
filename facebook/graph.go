package facebook

import (
	"os"
)

const (
	GRAPH_URL = "graph.facebook.com"
	SECURE    = "https://"
	UNSECURE  = "http://"
)

var RequestMetadata = map[string]string{"metadata": "1"}

type Graph struct {
	Https bool
}

func (g *Graph) GetGraphUrl() string {
	if g.Https {
		return SECURE + GRAPH_URL
	}
	return UNSECURE + GRAPH_URL
}

func (g *Graph) Call(id string, params map[string]string) (*Response, os.Error) {
	cmd := g.GetGraphUrl() + "/" + id
	if len(params) > 0 {
		cmd += "?"
		for key, val := range params {
			cmd += key + "=" + val + "&"
		}
		cmd = cmd[:len(cmd)-1] // Remove last &
	}
	return Get(cmd)
}

func (g *Graph) Publish(profile, id string, params map[string]string) (*Response, os.Error) {
	url := g.GetGraphUrl() + "/" + profile + "/" + id
	// TODO: Parse reponse here to handle error reponses.
	return PostForm(url, params)
}

var MainGraph = new(Graph)

func Call(id string, params map[string]string) (*Response, os.Error) {
	return MainGraph.Call(id, params)
}

func Publish(profile, id string, params map[string]string) (*Response, os.Error) {
	return MainGraph.Publish(profile, id, params)
}
