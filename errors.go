package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	er      error
}

func main() {

	err := dosomething()
	if err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Println("Operation successfull")
	}

}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func dosomething() error {
	err := dosomethingelse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}

	return nil
}

func dosomethingelse() error {
	return errors.New("internal error")
}
