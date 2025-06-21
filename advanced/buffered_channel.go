package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println("Value1", <-ch)
	fmt.Println("Value2", <-ch)

	ch <- 3

	fmt.Println("Value3", <-ch)

	fmt.Println("Buffered Chanel")

}
