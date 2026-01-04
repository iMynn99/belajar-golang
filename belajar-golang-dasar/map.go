package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println(person["Salah"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
