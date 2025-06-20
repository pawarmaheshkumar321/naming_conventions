package main

import "fmt"

func main() {

	greeting := make(chan string)
	greetingString := "Hello Mahesh"

	go func() {
		greeting <- greetingString
	}()

	receiver := <-greeting

	fmt.Printf("Received Value %s", receiver)

}
