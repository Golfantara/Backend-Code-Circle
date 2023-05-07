package main

import (
	"fmt"
)

func main() {
	var panjangPersegiPanjang int
	var lebarPersegiPanjang int
	var alasSegitiga int
	var tinggiSegitiga int

	fmt.Println("masukan panjang")
	fmt.Scan(&panjangPersegiPanjang)
	fmt.Println("masukan lebar")
	fmt.Scan(&lebarPersegiPanjang)
	hasilPersegiPanjang := 2 * (panjangPersegiPanjang + lebarPersegiPanjang)
	fmt.Println("hasil", +hasilPersegiPanjang)
	fmt.Println("masukan alas segitiga")
	fmt.Scan(&alasSegitiga)
	fmt.Println("masukan tinggi segitiga")
	fmt.Scan(&tinggiSegitiga)
	hasilLuasSegitiga := 0.5 * float32(alasSegitiga) * float32(tinggiSegitiga)
	fmt.Println("hasil Luas Segitiga : ", +hasilLuasSegitiga)
}
