package facebook

type Workplace struct {
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

func parseWork(value []interface{}) (workplaces []Workplace) {
	workplaces = make([]Workplace, len(value))
	for i, v := range value {
		wp := v.(map[string]interface{})
		index := 0
		for key, val := range wp {
			switch key {
			case "employer":
				workplaces[index].Employer = parseEmployer(val.(map[string]interface{}))
			case "position":
				workplaces[index].Position = parsePosition(val.(map[string]interface{}))
			case "start_date":
				workplaces[i].StartDate = v.(string)
			case "end_date":
				workplaces[i].EndDate = v.(string)
			}
			index++
		}
	}
	return
}

func parseEmployer(value map[string]interface{}) (employer Employer) {
	employer.ID = value["id"].(string)
	employer.Name = value["name"].(string)
	return
}

func parsePosition(value map[string]interface{}) (pos Position) {
	pos.ID = value["id"].(string)
	pos.Name = value["name"].(string)
	return
}
