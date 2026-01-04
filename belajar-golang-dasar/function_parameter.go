package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Nama :", firstName, lastName)
}

func main() {
	sayHelloTo("Eko", "Kurniawan")
	sayHelloTo("Budi", "Nugraha")
}
