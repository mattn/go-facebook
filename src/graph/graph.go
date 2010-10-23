package facebook

import(
  "http"
  "io/ioutil"
  "json"
  "os"
)

const (
  GRAPHURL = "http://graph.facebook.com/"
)

type Args struct {
  Metadata bool
}

type Page struct {
  ID string
  Name string
  Picture string
  Link string
  Category string
  Website string
  Username string
  Founded string
  CompanyOverview string
  Mission string
  Products string
  FanCount float64
}

type Person struct {
  ID string
  Name string
  FirstName string
  LastName string
  Link string
  Gender string
  Locale string
  UpdatedTime string
}

func FetchPage(id string) (page Page, err os.Error) {
  body, err := fetchBody(id)
  if err != nil { return }
  data, err := getJsonMap(body)
  if err != nil { return }  
  for key, value := range data {
	switch key {
	  case "website":
		page.Website = value.(string)
	  case "picture":
		page.Picture = value.(string)
	  case "fan_count":
		page.FanCount = value.(float64)
	  case "mission":
		page.Mission = value.(string)
	  case "category":
		page.Category = value.(string)
	  case "name":
		page.Name = value.(string)
	  case "username":
		page.Username = value.(string)
	  case "link":
		page.Link = value.(string)
	  case "id":
		page.ID = value.(string)
	  case "products":
		page.Products = value.(string)
	  case "founded":
		page.Founded = value.(string)
	  case "company_overview":
		page.CompanyOverview = value.(string)
	}
  }  
  return
}

func FetchPerson(name string) (person Person, err os.Error) {
  body, err := fetchBody(name)
  if err != nil { return }
  data, err := getJsonMap(body)
  if err != nil { return }
  for key, value := range data {
	switch key {
	  case "locale":
		person.Locale = value.(string)
	  case "name":
		person.Name = value.(string)
	  case "link":
		person.Link = value.(string)
	  case "Gender":
		person.Gender = value.(string)
	  case "first_name":
		person.FirstName = value.(string)
	  case "last_name":
		person.LastName = value.(string)
	  case "id":
		person.ID = value.(string)
	}
  }  
  return
}


func getJsonMap(body []byte) (data map[string] interface{}, err os.Error) {
  var values interface{}
  
  if err = json.Unmarshal(body, &values); err != nil {
	return
  }
  data = values.(map[string] interface{})
  return
}

func fetchBody(method string) (body []byte, err os.Error) {
  resp, _, err := http.Get(GRAPHURL + method) // Response, final URL, error
  if err != nil { return }
  defer resp.Body.Close()
  
  body, err = ioutil.ReadAll(resp.Body)
  return
}
