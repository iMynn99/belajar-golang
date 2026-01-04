package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Validation Errors")
	notFoundError   = errors.New("Data Not Found")
)

func GetById(id string) error {
	if id == "" {
		return validationError
	}
	if id != "Eko" {
		return notFoundError
	}
	return nil
}

func main() {
	err := GetById("Eko")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Not Found Errors")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
