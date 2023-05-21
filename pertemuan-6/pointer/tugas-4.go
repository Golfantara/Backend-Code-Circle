/*
buatlah function dengan nama tambahDataFilm untuk menambahkan data map[string]string ke
slice dataFilm, lalu tampilkan datanya sesuai output yang di inginkan
var dataFilm = []map[string]string{}
tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
tambahDataFilm("spiderman", "2 jam", "action", "2004",
&dataFilm)
tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
// isi dengan jawaban anda untuk menampilkan data
berikut ini output soal 4(untuk urutan title atau genre muncul duluan itu tidak mesti persis):
*/

/*
buatlah function yang menambahkan data di bawah ini ke variabel buah (wajib menggunakan
pointer):
● Jeruk
● Semangka
● Mangga
● Strawberry
● Durian
● Manggis
● Alpukat
lalu tampilkan dengan loop dan beri angka di depannya sehingga menghasilkan output seperti
ini:
1. Jeruk
2. Semangka
3. Mangga
4. Strawberry
5. Durian
6. Manggis
7. Alpukat

*/

package main

import "fmt"

func addToFilm(judul, durasi, category, release string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"judul":    judul,
		"durasi":   durasi,
		"category": category,
		"release":  release,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	var dataFilm = []map[string]string{}

	addToFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	addToFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	addToFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	addToFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d. %s \n", i+1, item)
	}
}
