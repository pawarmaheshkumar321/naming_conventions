package main

import (
	"fmt"
	"os"
)

func main() {

	tempFile, err := os.CreateTemp("", "temporary_file.txt")
	if err != nil {
		fmt.Println("Error :", err)
		panic(err)
	}

	fmt.Println("Temporary tempFile created :", tempFile.Name())

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

}
