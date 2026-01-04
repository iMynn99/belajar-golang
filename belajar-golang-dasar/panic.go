package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("endApp")
	message := recover()
	fmt.Println("Terjadi panic:", message)
}

func runnApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi mengalami error")
	} else {
		fmt.Println("Aplikasi berjalan dengan baik")
	}

}

func main() {
	runnApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}

/**
Output:
endApp
Terjadi panic: Aplikasi mengalami error
*/
