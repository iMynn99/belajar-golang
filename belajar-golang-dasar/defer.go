package main

func logging() {
	println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	println("Run Application")
}

func main() {
	runApplication()
}
