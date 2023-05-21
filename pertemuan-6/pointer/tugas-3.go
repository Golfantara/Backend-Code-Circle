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

func addToBuah(slice *[]string, elements string) {
	*slice = append(*slice, elements)
}

func main() {
	var buah = []string{}

	addToBuah(&buah, "Jeruk")
	addToBuah(&buah, "Semangka")
	addToBuah(&buah, "Mangga")
	addToBuah(&buah, "Strawberry")
	addToBuah(&buah, "Durian")
	addToBuah(&buah, "Manggis")
	addToBuah(&buah, "Alpukat")

	for i, item := range buah {
		fmt.Printf("%d. %s \n", i+1, item)
	}
}
