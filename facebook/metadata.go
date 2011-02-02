package facebook

import (
	"os"
	"json"
)

type Metadata struct {
	Connections map[string]string
	Fields      []*Field
}

func (m *Metadata) GetActivities() (a *Activities, err os.Error) {
	url, ok := m.Connections["activities"]
	if !ok {
		return //nil, os.NewError("No Metadata.Connections[activities].")
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
