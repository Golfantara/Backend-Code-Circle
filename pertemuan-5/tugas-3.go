package main

import "fmt"

func fruits(name string, jenisBuah ...string) {
	fmt.Print("hallo nama saya " + name)
	fmt.Print("buah kesukaan saya ")
	for _, buah := range jenisBuah {
		fmt.Print(buah)
	}
}

func main() {
	fruits("dejan ", "mangga ", "apel ", "durian ")
	fmt.Print()
}
