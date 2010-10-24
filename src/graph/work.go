package facebook

type Work struct {
	Employer  Employer
	Position  Position
	StartDate string
	EndDate   string
}

type Employer struct {
	ID   string
	Name string
}

type Position struct {
	ID   string
	Name string
}
