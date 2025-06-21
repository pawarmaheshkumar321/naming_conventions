package main

import (
	"fmt"
	"time"
)

func main() {

	/*ch := make(chan int)

	select {
	case msg := <-ch:
		fmt.Println("Received", msg)
	default:
		fmt.Println("No messges received")
	}

	select {
	case ch <- 1:
		fmt.Println("Message sent", 1)
	default:
		fmt.Println("Channel is not ready to receive")

	}*/

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data Received", d)

			case <-quit:
				fmt.Println("Stopping")
				close(data)
				close(quit)
				return
			default:
				fmt.Println("Waiting for data")
				time.Sleep(time.Millisecond * 100)

			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Millisecond * 100)
	}

	quit <- true
}
