package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absen = 81

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsen = absen > 80

	var lulus = lulusNilaiAkhir && lulusAbsen
	fmt.Println(lulus)
}
