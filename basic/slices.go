package main

import (
	"fmt"
	"slices"
)

func slicesfun() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	ar := [5]int{1, 2, 3, 4, 5}

	slice1 := ar[1:4]

	slice1 = append(slice1, 6, 7, 8)

	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	fmt.Println(slice2)
	fmt.Println(slice1)

	var nilSlice []int
	fmt.Println(nilSlice)

	if slices.Equal(slice1, slice2) {
		fmt.Println("slices equal")
	}

	k := 1
	twoDimSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoDimSlice[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoDimSlice[i][j] = k
			k++
		}
	}

	fmt.Println(twoDimSlice)
	fmt.Printf("Capacity of slice1 is %d", cap(slice1))
	fmt.Printf("\nLength of slice1 is %d", len(slice1))

}
