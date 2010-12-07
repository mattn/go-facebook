package main

import (
  "facebook/graph"
  "fmt"
)

func main() {
	g := graph.NewGraph("")
	platform, err := g.FetchPage("platform")
	if err != nil {
		fmt.Println(err.String())
	}
	info := "Name: " + platform.Name + "\n"
	info += "Website: " + platform.Website + "\n"
	info += "Username: " + platform.Username + "\n"
	info += "Founded: " + platform.Founded.String() + "\n"
	info += "CompanyOverview: " + platform.CompanyOverview + "\n"
	info += "Mission: " + platform.Mission + "\n"
	info += "Products: " + platform.Products + "\n"
	info += fmt.Sprintf("FanCount: %f\n", platform.FanCount)
	info += "ID: " + platform.ID + "\n"
	info += "Category: " + platform.Category + "\n"
	fmt.Print(info)
}
