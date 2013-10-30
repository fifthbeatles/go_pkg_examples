package main

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func NewMyError(s string) error {
	return MyError{time.Now(), s}
}

func main() {
	err := errors.New("Created by errors.New")
	if err != nil {
		fmt.Println(err)
	}
	err = NewMyError("MyError with When and What")
	if err != nil {
		fmt.Println(err)
	}
}
