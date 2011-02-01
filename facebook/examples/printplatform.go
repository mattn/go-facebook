package main

import (
	facebook "github.com/Agon/go-facebook/facebook"
	"fmt"
)

func main() {
	resp, err := facebook.Call("platform", map[string]string{})
	if err != nil {
		fmt.Println(err.String())
		return
	}
	platform := resp.Data.(map[string]interface{})
	info := "Name: " + platform["name"].(string) + "\n"
	fmt.Print(info)
}