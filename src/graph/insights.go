package graph

import (
	"os"
	"time"
	"strings"
)

/* Statistics about applications, pages, and domains.
 * Available metrics include application and page hits, adds, removes, and likes.
 * The FQL Insights Documentation contains a complete list of available metrics.
 * Without a session, an application can retrieve only its own metrics.
 * With a user session, it is possible to retrieve data for all applications, pages, and domains owned by the session user.
 * Reading this data requires the read_insights extended permission.
 * Use Facebook Insights to claim your domain.
 * NOTE: Insights can be retrieved only as an array.
 */
type Insights struct {
	ID       string
	Insights []Insight
}

func getInsights(url string) (is Insights, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	is, err = parseInsights(resp.Data)
	return
}

func parseInsights(value []interface{}) (is Insights, err os.Error) {
	is.Insights = make([]Insight, len(value))
	for i, v := range value {
		is.Insights[i], err = parseInsight(v.(map[string]interface{}))
	}
	is.ID = strings.Split(is.Insights[0].ID, "/", -1)[0]
	return
}

type Insight struct {
	// ID of the insight. Requires read_insight permission.
	ID string
	// Name of the insight. Requires read_insight permission.
	Name string
	// Length of the period during which the insights were collected. Requires read_insight permission. JSON string containing 'day', 'week' or 'month'.
	Period string
	// Individual data points for the insight. Requires read_insight permission.
	// A JSON array of objects containing the value (a JSON number) and end_time (A JSON string containing a IETF RFC 3339 datetime) fields.
	Values map[float64]*time.Time

	// Not documented but streamed
	Description string
}

/*
 * Parses Insight data. Returns nil for err if no error appeared.
 */
func parseInsight(value map[string]interface{}) (i Insight, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			i.ID = val.(string)
		case "name":
			i.Name = val.(string)
		case "period":
			i.Period = val.(string)
		case "values":
			i.Values = make(map[float64]*time.Time)
			for _, v := range val.([]interface{}) {
				vl := v.(map[string]interface{})
				i.Values[vl["value"].(float64)], err = parseTime(vl["end_time"].(string))
			}
		case "description":
			i.Description = val.(string)
		}
	}
	return
}
