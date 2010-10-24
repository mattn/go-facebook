package facebook

type Connections struct {
	Albums     string
	Statuses   string
	Links      string
	Posts      string
	Notes      string
	Videos     string
	Feed       string
	Photos     string
	Tagged     string
	Events     string
	Home       string
	Friends    string
	Activities string
	Interests  string
	Music      string
	Books      string
	Movies     string
	Television string
	Likes      string
	Groups     string
	InBox      string
	OutBox     string
	Updates    string
	Accounts   string
	CheckIns   string
	Picture    string
	Family     string
}

func parseConnections(value interface{}) (connections Connections) {
	data := value.(map[string]interface{})
	for key, v := range data {
		switch key {
		case "feed":
			connections.Feed = v.(string)
		case "posts":
			connections.Posts = v.(string)
		case "tagged":
			connections.Tagged = v.(string)
		case "statuses":
			connections.Statuses = v.(string)
		case "links":
			connections.Links = v.(string)
		case "notes":
			connections.Notes = v.(string)
		case "photos":
			connections.Photos = v.(string)
		case "albums":
			connections.Albums = v.(string)
		case "events":
			connections.Events = v.(string)
		case "videos":
			connections.Videos = v.(string)
		case "home":
			connections.Home = v.(string)
		case "friends":
			connections.Home = v.(string)
		case "activities":
			connections.Activities = v.(string)
		case "interests":
			connections.Interests = v.(string)
		case "music":
			connections.Music = v.(string)
		case "books":
			connections.Books = v.(string)
		case "movies":
			connections.Movies = v.(string)
		case "television":
			connections.Television = v.(string)
		case "likes":
			connections.Likes = v.(string)
		case "groups":
			connections.Groups = v.(string)
		case "inbox":
			connections.InBox = v.(string)
		case "outbox":
			connections.OutBox = v.(string)
		case "updates":
			connections.Updates = v.(string)
		case "accounts":
			connections.Accounts = v.(string)
		case "checkins":
			connections.CheckIns = v.(string)
		case "picture":
			connections.Picture = v.(string)
		case "family":
			connections.Family = v.(string)
		default:
			debugInterface(v, key, "Connections")
		}
	}
	return
}
