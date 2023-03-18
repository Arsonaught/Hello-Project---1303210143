package main

import "fmt"

func main() {
	var b, prima int
	fmt.Scan(&b)
	prima = 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			prima += 1
		}
	}
	fmt.Print(prima)
	if prima == 2 {
		fmt.Print("\nPrima : true")
	} else {
		fmt.Print("\nPrima : false")
	}
}
