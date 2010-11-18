package graph

type Paging struct {
	Previous string
	Next     string
}

func parsePaging(value map[string]interface{}) (paging Paging) {
	paging.Previous = value["previous"].(string)
	paging.Next = value["next"].(string)
	return
}
