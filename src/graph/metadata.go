package facebook

type Metadata struct {
	Connections Connections
	// Name, Description
	Fields map[string]string
}

func parseMetaData(value interface{}) (metadata Metadata) {
	data := value.(map[string]interface{})
	for key, v := range data {
		switch key {
		case "connections":
			metadata.Connections = parseConnections(v)
		case "fields":
			metadata.Fields = parseFields(v)
		default:
			debugInterface(v, key, "Metadata")
		}
	}
	return
}

func parseFields(value interface{}) (fields map[string]string) {
	fields = make(map[string]string)
	var field map[string]interface{}
	data := value.([]interface{})
	for _, c := range data {
		field = c.(map[string]interface{})
		fields[field["name"].(string)] = field["description"].(string)
	}
	return
}
