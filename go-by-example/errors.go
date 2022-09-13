package main

import (
	"errors"
	"fmt"
)

// by convention errors are the last return value and have type error (a built in interface)
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New constructs a basic error Value with the given error message
		return -1, errors.New("can't work with 42")
	}

	// nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// here, we create a custom types as errors 
// we must implement the Error method on custom errors
type argError struct {
	arg int
	prob string
}

// implementing the method Error on our ardError type
// because of the Error built-in interface? 
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//creating a function to test our errors
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}