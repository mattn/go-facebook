package facebook


// Connection
type Activities struct {
	Data []*Activity
}

type Activity struct {
	Id           string
	Name         string
	Category     string
	Created_Time string
}
