package facebook

type Town struct {
	ID   string
	Name string
}

func parseTown(value map[string]interface{}) (ht Town) {
	for k, v := range value {
		switch k {
		case "id":
			ht.ID = v.(string)
		case "name":
			ht.Name = v.(string)
		}
	}
	return
}
