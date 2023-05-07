package main

import (
	"fmt"
)

func main() {
	var tanggal = 17
	var bulan = 8
	var tahun = 1945

	switch bulan {
	case 8:
		fmt.Println(tanggal, "agustus", tahun)
	}
}
