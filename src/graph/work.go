package facebook

import (
	"os"
	"time"
)

type Workplace struct {
	Employer  Object
	Position  Object
	StartDate *time.Time
	EndDate   string
}

func parseWork(value []interface{}) (workplaces []Workplace, err os.Error) {
	workplaces = make([]Workplace, len(value))
	for i, v := range value {
		wp := v.(map[string]interface{})
		index := 0
		for key, val := range wp {
			switch key {
			case "employer":
				workplaces[index].Employer = parseObject(val.(map[string]interface{}))
			case "position":
				workplaces[index].Position = parseObject(val.(map[string]interface{}))
			case "start_date":
				workplaces[i].StartDate, err = time.Parse("RFC3339", v.(string))
			case "end_date":
				workplaces[i].EndDate = v.(string)
			}
			index++
		}
	}
	return
}
