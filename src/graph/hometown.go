package facebook

type Hometown struct {
	ID   string
	Name string
}

func parseHometown(value map[string]interface{}) (ht Hometown) {
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
