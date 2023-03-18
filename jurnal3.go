package main
import "fmt"
func main() {
	var matakuliah, sks int
	var indeks string
	fmt.Scan(&matakuliah)
	jumNilai = 0
	jumSks = 0
	for i := 1, i <= matakuliah; i ++ {
		fmt.Scan(&indeks, &sks)
		for indeks != "A" && indeks != "B" && indeks != "C" && indeks != "D" && indeks != "E" || sks < 0 {
			fmt.Scan(&indeks, &sks)
		}
		jumSks += sks
		if indeks == "A"{
			jumNilai += 4*sks
		} else if indeks == "B" {
			jumNilai += 3*sks
		} else if indeks == "C" {
			jumNilai += 2*sks
		} else if indeks == "D" {
			jumNilai += 1*sks
		} else if indeks == "E" {
			jumNilai += 0*sks
		}
	}
	fmt.Println(float64(jumNilai)/float64(jumSks))		
}