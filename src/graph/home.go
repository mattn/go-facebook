package facebook

type News struct {
	ID string
	From Object
	To []Object
	Message string
	Actions []Link
	Type string
	CreatedTime string
	UpdatedTime string
	Likes string
	Comments []Comment
}  

type Comment struct {
	ID string
	From Object
	Message string
	CreatedTime string
	
}
