package facebook

const (
	TYPE_EDU_HIGHSCHOOL     = "High School"
	TYPE_EDU_COLLEGE        = "College"
	TYPE_EDU_GRADUATESCHOOL = "Graduate School"
)

type Education struct {
	School        *Object
	Year          *Object
	Concentration *Object
	Degree        *Object
	Type          string
}
