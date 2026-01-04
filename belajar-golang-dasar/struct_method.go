package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	eko := Customer{
		Name:    "Eko",
		Address: "Indonesia",
		Age:     30,
	}
	eko.sayHello("Budi")
	eko.sayHello("Joko")
	eko.sayHello("Kurniawan")
}
