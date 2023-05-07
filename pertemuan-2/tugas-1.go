package main

import (
	"fmt"
)

func main() {
	var angka1 int
	var angka2 int
	var input int8
	fmt.Println(" INI KALKULATOR SEDERHANA ")
	fmt.Println("1. Pertambahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. pembagian")
	fmt.Println("4. perkalian")
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Masukan Angka Awal")
		fmt.Scan(&angka1)
		fmt.Println("Masukan Angka Kedua")
		fmt.Scan(&angka2)
		fmt.Println("Angka 1 + Angka 2 = ", angka1+angka2)
	case 2:
		fmt.Println("Masukan Angka Awal")
		fmt.Scan(&angka1)
		fmt.Println("Masukan Angka Kedua")
		fmt.Scan(&angka2)
		fmt.Println("Angka 1 - Angka 2 = ", angka1-angka2)
	case 3:
		fmt.Println("Masukan Angka Awal")
		fmt.Scan(&angka1)
		fmt.Println("Masukan Angka Kedua")
		fmt.Scan(&angka2)
		fmt.Println("Angka 1 * Angka 2 = ", angka1*angka2)
	case 4:
		fmt.Println("Masukan Angka Awal")
		fmt.Scan(&angka1)
		fmt.Println("Masukan Angka Kedua")
		fmt.Scan(&angka2)
		fmt.Println("Angka 1 / Angka 2 = ", angka1/angka2)
	}

}
