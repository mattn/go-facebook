package main

import (
	facebook "github.com/Agon/go-facebook/facebook"
	"fmt"
)

func main() {
	platform, err := facebook.GetPage("platform")
	if err != nil {
		fmt.Println(err.String())
		return
	}
	fmt.Printf("Name: %s\n", platform.Name)
}
