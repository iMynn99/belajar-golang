package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	case "Kurniawan":
		fmt.Println("Hello Kurniawan")
	default:
		fmt.Println("Hello, Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name terlalu panjang")
	case false:
		fmt.Println("Name sudah benar")
	}

	name = "Khannedy"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name terlalu panjang")
	case length > 5:
		fmt.Println("Name lumayan panjang")
	default:
		fmt.Println("Name sudah benar")
	}

}
