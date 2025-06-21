package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	datach := make(chan string)

	go func() {
		for i := range 5 {
			datach <- fmt.Sprintf("Hello %s", strconv.Itoa(i))
			time.Sleep(time.Millisecond * 100)
		}
		close(datach)
	}()

	for value := range datach {
		fmt.Println("Received value", value, ":", time.Now())
	}
}

// sync multiple goroutine
/*func main() {

	numGoroutines := 3
	done := make(chan int, 3)

	for i := range numGoroutines {
		go func(id int) {
			fmt.Println(" value", id)
			//time.Sleep(time.Second)
			done <- id
		}(i)
	}

	for range numGoroutines {
		temp := <-done
		fmt.Println(" value_received", temp)
	}

	fmt.Println("finished")
}

func main() {

	ch := make(chan int)

	go func() {
		ch <- 9
		time.Sleep(time.Second * 1)
		fmt.Println("Value Pushed")
	}()

	value := <-ch
	fmt.Println("Value is ", value)

}

func main() {

	done := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(time.Second * 2)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Finished")

}*/
