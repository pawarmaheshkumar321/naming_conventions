package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Command : ", os.Args[0])
	//fmt.Println("Command : ", os.Args[1])

	for index, arg := range os.Args {
		fmt.Println("Command %d : %v", index, arg)
	}
}
