package main

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "12q"

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(num)

	floatStr := "12.03"
	floatNum, err := strconv.ParseFloat(floatStr, 32)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Println("floatNum", floatNum)

	fmt.Printf("floatNum = %.2f", floatNum)

}
