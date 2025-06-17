package main

import (
	"fmt"
	"time"
)

func main() {

	// Mon Jan 2 15:04:05 MST 2006
	//timeNow := time.Now()

	layout := "2006-01-02"

	strTime := "01-06-2025T15:04:05"

	fmt.Println(time.Parse(layout, strTime))
}
