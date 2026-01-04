package main

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	println(a, b, c)
}
