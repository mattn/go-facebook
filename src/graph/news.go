package facebook

import (
	"time"
)

type News struct {
	ID          string
	From        Object
	To          []Object
	Message     string
	Actions     []Link
	Picture     Picture
	Type        string
	Link        string
	Name        string
	Icon        string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Comments    []Comment
	Likes       float64
}
