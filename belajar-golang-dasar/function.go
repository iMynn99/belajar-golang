package main

import "fmt"

func sayHello() {
	fmt.Println("Hello!")
}

func hitungKotak() {
	var panjang int = 10
	var lebar int = 5
	var luas int = panjang * lebar
	fmt.Println("Luas kotak adalah :", luas)
}

func main() {
	sayHello()
	sayHello()
	hitungKotak()
}
