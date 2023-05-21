/*
Tulislah sebuah function dengan nama introduce. pastikan semua parameter pada function
introduce di gunakan semuanya (wajib menggunakan pointer)
var sentence string
introduce(&sentence, "John", "laki-laki", "penulis", "30")
fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang
berusia 30 tahun"
introduce(&sentence, "Sarah", "perempuan", "model", "28")
fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia
28 tahun"
*/

package main

import "fmt"

type porto struct {
	name, gender, hobi, age string
}

func introduce(Porto *porto) {
}

func main() {
	var sentance = porto{"lutfi", "pria", "ngoding", "50"}

	fmt.Printf("pak %s adalah seorang %s yang berusia %s tahun \n", sentance.name, sentance.hobi, sentance.age)

	var sentance2 *porto = &sentance

	*sentance2 = porto{"rohendo", "pria", "ngoding", "50"}
	fmt.Printf("pak %s adalah jelma %s nu berusia %s tahun", sentance2.name, sentance2.hobi, sentance2.age)
}
