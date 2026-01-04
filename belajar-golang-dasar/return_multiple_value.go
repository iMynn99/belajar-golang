package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
