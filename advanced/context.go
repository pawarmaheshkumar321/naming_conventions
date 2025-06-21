package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context TODO):", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result with context timeout:", result)

	time.Sleep(time.Second * 2)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result after context timeout:", result)

}

func checkEvenOdd(ctx context.Context, num2 int) string {

	select {
	case <-ctx.Done():
		return "Operation Cancelled"
	default:
		if num2%2 == 0 {
			return fmt.Sprintf("Number %v is Even", num2)
		} else {
			return fmt.Sprintf("Number %v is Odd", num2)
		}
	}

}
