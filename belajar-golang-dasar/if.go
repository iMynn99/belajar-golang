package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hello, Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name terlalu panjang")
	} else {
		fmt.Println("Name sudah benar")
	}
}
