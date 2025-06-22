package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error while opening file : ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	err = scanner.Err()

	if err != nil {
		fmt.Println("Error scanning file : ")
		return
	}

}
