package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "Eko" {
		return &notFoundError{"data not found"}
	}

	// ok

	return nil
}

func main() {
	err := SaveData("Budi", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error:", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown Error:", err.Error())
		// }

		switch e := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", e.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", e.Error())
		default:
			fmt.Println("Unknown Error:", e.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
