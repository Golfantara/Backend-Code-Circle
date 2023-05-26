package main

import (
	"fmt"
)

func main() {
	var prefix interface{} = "hasil penjumlahan dari\n"
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)
	jumlah := 0
	for _, angka := range angkaPertama {
		jumlah += angka
	}
	for _, angka := range angkaKedua {
		jumlah += angka
	}

	output := fmt.Sprintf("%s", prefix)
	for i, angka := range angkaPertama {
		output += fmt.Sprintf("%d", angka)
		if i < len(angkaPertama)-1 || len(angkaKedua) > 0 {
			output += " + "
		}
	}
	for i, angka := range angkaKedua {
		output += fmt.Sprintf("%d", angka)
		if i < len(angkaKedua)-1 {
			output += " + "
		}
	}
	output += fmt.Sprintf(" = %d", jumlah)

	fmt.Println(output)
}
