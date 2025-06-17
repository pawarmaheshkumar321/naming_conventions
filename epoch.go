package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Unix())

	unixt := now.Unix()
	fmt.Println(unixt)

	fmt.Println(time.Unix(unixt, 0))

}
