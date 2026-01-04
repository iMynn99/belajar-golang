package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "1111111111"

	var contoh = "22222222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKTP)
}
