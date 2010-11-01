package facebook

import (
	"time"
)

type Feed struct {
	ID          string
	From        Object
	To          []Object
	Message     string
	Actions     []Link
	Privacy     []string // TODO: Replace with a Privacy struuct
	Picture     Picture
	Link        string
	Name        string
	Icon        string
	Type        string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Comments    []Comment
}
