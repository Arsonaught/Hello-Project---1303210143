package main

import "fmt"

func main() {
	var membership string
	var A, B, C, D, E int
	var cashback, diskon float64
	fmt.Scan(&membership, &A, &B, &C, &D, &E)
	cashback = 0
	diskon = 0
	if A%2 == 0 && B%2 == 0 && C%2 == 0 && D%2 == 0 && E%2 == 0 {
		cashback = 3.1 * float64(A+B+C)
	} else if A%2 != 0 && B%2 != 0 && C%2 != 0 && D%2 != 0 && E%2 != 0 {
		diskon = 1.7 * float64(C+D+E)
	} else {
		diskon = 1.7 * float64(C+D+E)
		cashback = 3.1 * float64(A+B+C)
	}
	if membership == "yes" {
		diskon *= 1.15
		cashback *= 1.15
	}
	if cashback > 35 {
		cashback = 35
	}
	if diskon > 50 {
		diskon = 50
	}
	fmt.Print(cashback, diskon)

}
