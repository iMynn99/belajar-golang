package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHei(value HashName) {
	fmt.Println("Hello " + value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups() any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	person := Person{Name: "Eko"}
	sayHei(person)
	animal := Animal{Name: "Kucing"}
	sayHei(animal)

	var kosong any = Ups()
	fmt.Println(kosong)
}
