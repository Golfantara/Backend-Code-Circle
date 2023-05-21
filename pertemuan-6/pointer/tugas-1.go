package main

import (
	"fmt"
	"math"
)

func hitungLuasKelilingLingkaran(jariJari float64, luas *float64, keliling *float64) {
	*luas = math.Pi * jariJari * jariJari
	*keliling = 2 * math.Pi * jariJari
}

func main() {
	var jariJari float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	var luasLingkaran float64
	var kelilingLingkaran float64

	hitungLuasKelilingLingkaran(jariJari, &luasLingkaran, &kelilingLingkaran)

	fmt.Printf("Luas lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling lingkaran: %.2f\n", kelilingLingkaran)
}
