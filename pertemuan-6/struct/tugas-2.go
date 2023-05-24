package main

import "fmt"

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s *segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (p *persegi) luas() int {
	return (p.sisi * p.sisi)
}

func (pp *persegiPanjang) luas() int {
	return (pp.panjang * pp.lebar)
}

func main() {
	hSegitiga := &segitiga{alas: 4, tinggi: 3}
	hPersegi := &persegi{sisi: 5}
	hPersegiPanjang := &persegiPanjang{panjang: 5, lebar: 3}

	fmt.Println("luas segitiga : ", hSegitiga.luas())
	fmt.Println("luas persegi : ", hPersegi.luas())
	fmt.Println("luas persegi panjang :", hPersegiPanjang.luas())

}
