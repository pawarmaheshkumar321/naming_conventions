package main

import (
	"fmt"
	"time"
)

func main() {

	timeNow := time.Now()

	fmt.Println(timeNow)

	fmt.Println(timeNow.Add(time.Hour * 24))

	fmt.Println(timeNow.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println(timeNow.In(loc).Format("Monday 2006-01-02 15:04:05"))
}
