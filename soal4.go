package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var A, B, C, D int
	fmt.Scanln(&A)
	fmt.Scanln(&B)

	rand.Seed(int64(A))
	C = rand.Intn(4) + 1
	fmt.Println(C)
	D = rand.Intn(4) + 1
	fmt.Println(D)

	if B == D && !(C == D) {
		fmt.Print("Pemenang Adalah Anda")
	} else if C == D && !(B == D) {
		fmt.Print("Pemenang Adalah Pak Pras")
	} else if C == D && B == D {
		fmt.Print("Seri")
	} else {
		fmt.Print("Tidak Ada Pemenang")
	}
}
