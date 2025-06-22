package main

import (
	"fmt"
	"net/url"
)

func main() {

	rawUrl := "http://example.com/path?name=maheshkumar&age=30"

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Printf("Name is %s and Age is %s", parsedUrl.Query().Get("name"), parsedUrl.Query().Get("age"))

}
