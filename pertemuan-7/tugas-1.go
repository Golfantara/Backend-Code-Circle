package main

import "fmt"

func tanggal(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

func main() {
	defer tanggal("golang Backend Development", 2021)
	fmt.Println("test")
	fmt.Println("testest")
}
