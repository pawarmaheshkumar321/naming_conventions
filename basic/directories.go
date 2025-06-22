package main

import (
	"fmt"
	"os"
)

func main() {

	//checkError(os.Mkdir("output", 0777))
	checkError(os.MkdirAll("output", 0777))
	checkError(os.WriteFile("output/file1", []byte("hello, world!!!"), 0777))

	result, err := os.ReadDir("output")
	checkError(err)

	fmt.Println("Result :", result)
	fmt.Println("Result :", result[0])
	//////////////////////////////
	dir, err := os.ReadDir("output")
	checkError(err)

	fmt.Println("Current Dir :", dir)
	/////////////////////////////////////
	checkError(os.Chdir("../"))
	dir1, err := os.Getwd()
	checkError(err)

	fmt.Println("Working Dir :", dir1)

	//////////////////

	/*	filepath.WalkDir(){

		}*/

	err = os.Remove("output")
	if err != nil {
		fmt.Println("Error :", err)
		panic(err)
	}

	err = os.RemoveAll("output")
	if err != nil {
		fmt.Println("Error :", err)
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error :", err)
		panic(err)
	}
}
