package graph

const (
	TYPE_EDU_HIGHSCHOOL     = "High School"
	TYPE_EDU_COLLEGE        = "College"
	TYPE_EDU_GRADUATESCHOOL = "Graduate School"
)

type Education struct {
	School        Object
	Year          Object
	Concentration Object
	Degree        Object
	Type          string
}

func parseEducations(value []interface{}) (educations []Education) {
	educations = make([]Education, len(value))
	for i, v := range value {
		wp := v.(map[string]interface{})
		for key, val := range wp {
			switch key {
			case "school":
				educations[i].School = parseObject(val.(map[string]interface{}))
			case "year":
				educations[i].Year = parseObject(val.(map[string]interface{}))
			case "concentration":
				educations[i].Concentration = parseObject(val.(map[string]interface{}))
			case "degree":
				educations[i].Degree = parseObject(val.(map[string]interface{}))
			case "type":
				educations[i].Type = v.(string)
			}
		}
	}
	return
}
