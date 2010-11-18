package graph

type Venue struct {
	Street    string
	City      string
	State     string
	Country   string
	Latitude  float64
	Longitude float64
}

func parseVenue(value map[string]interface{}) (v Venue) {
	v.Street = value["street"].(string)
	v.City = value["city"].(string)
	v.State = value["state"].(string)
	v.Country = value["country"].(string)
	v.Latitude = value["latitude"].(float64)
	v.Longitude = value["longitude"].(float64)
	return
}
