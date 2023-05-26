package main

import (
	"fmt"
)

func main() {
	var prefix interface{} = "hasil penjumlahan dari\n"
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Menggunakan type assertion untuk mengonversi variabel ke tipe yang sesuai
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	// Menghitung jumlah angka
	jumlah := 0
	for _, angka := range angkaPertama {
		jumlah += angka
	}
	for _, angka := range angkaKedua {
		jumlah += angka
	}

	// Menghasilkan output
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
