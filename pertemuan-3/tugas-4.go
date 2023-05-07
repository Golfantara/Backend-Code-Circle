package main

import "fmt"

func main() {
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("D")
	} else if nilaiJohn < 50 {
		fmt.Println("E")
	} else {
		fmt.Println("wrong option")
	}

	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("D")
	} else if nilaiDoe < 50 {
		fmt.Println("E")
	} else {
		fmt.Println("wrong option")
	}
}
