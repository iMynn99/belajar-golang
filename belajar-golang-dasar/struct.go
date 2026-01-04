package main

import "fmt"

type Customers struct {
	Name, Addres string
	Age          int
}

func main() {
	customer := Customers{}
	fmt.Println(customer)

	customer.Name = "Eko"
	customer.Addres = "Indonesia"
	customer.Age = 30
	fmt.Println(customer)

	customer2 := Customers{
		Name:   "Budi",
		Addres: "Indonesia",
		Age:    25,
	}
	fmt.Println(customer2)

	customer3 := Customers{"Joko", "Indonesia", 35}
	fmt.Println(customer3)
}
