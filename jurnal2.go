package main

import "fmt"

func main() {
	var berat, kg, gr, harga_kg, harga_gr, total int

	fmt.Print("berat parsel (gram): ")
	fmt.Scan(&berat)
	kg = berat / 1000
	gr = berat % 1000
	fmt.Println("detail berat:", kg, "kg +", gr, "gr")
	harga_kg = kg * 10000
	if gr >= 500 {
		harga_gr = gr * 5
	} else {
		harga_gr = gr * 15
	}
	fmt.Println("detail biaya: Rp.", harga_kg, "+ Rp.", harga_gr)
	if kg > 10 {
		total = harga_kg
	} else {
		total = harga_kg + harga_gr
	}
	fmt.Println("total biaya: Rp.", total)
}
