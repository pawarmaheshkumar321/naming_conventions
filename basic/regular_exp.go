package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[aeiou]`)
	str := "Hello world!"
	result := re.ReplaceAllString(str, "*")

	fmt.Printf(" %s", result)

}
