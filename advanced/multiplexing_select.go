package main

import "fmt"

func main() {

	intch := make(chan int)

	go func() {
		intch <- 1
		close(intch)
	}()

	for {
		select {
		case msg, ok := <-intch:
			if !ok {
				fmt.Println("Channel closed", ok)
				return
			}
			fmt.Printf("Value receive on channel is : %d\n", msg)
		}
	}

}
