package main

import "fmt"

func main() {

	sequence := adder(0)
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder(5)
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	subtracter := func() func(int) int {
		countdown := 99

		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))

}

func adder(start_point int) func() int {
	//start_point := 0
	fmt.Println("Previous value of i:", start_point)
	return func() int {
		start_point++
		fmt.Println("added 1 to start_point")
		return start_point
	}
}
