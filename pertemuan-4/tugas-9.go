package main

import (
	"fmt"
)

func main() {
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	for i := 2; i < 7; i++ {
		fmt.Printf("%s ", kalimat[i])
	}
}
