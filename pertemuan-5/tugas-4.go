package main

import (
	"fmt"
)

func main() {
	var dataFilm []map[string]string

	addDataFilm := func(title, duration, genre, year string) {
		film := map[string]string{
			"title":    title,
			"duration": duration,
			"genre":    genre,
			"year":     year,
		}

		dataFilm = append(dataFilm, film)
	}

	addDataFilm("GOTG", "3 Jam", "Action", "2023")
	addDataFilm("kodrat", "3 Jam", "Horror", "2023")
	addDataFilm("KKN", "3 Jam", "Horror", "2022")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
