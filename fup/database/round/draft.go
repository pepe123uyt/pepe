package main

import (
	"fmt"
	"math"
)

func main() {
	var operador rune
	var x float64

	fmt.Scanf("%c", &operador)
	fmt.Scan(&x)

	if operador == 'c' {
		fmt.Printf("%.0f\n", math.Ceil(x))
	} else if operador == 'f' {
		fmt.Printf("%.0f\n", math.Floor(x))
	} else if operador == 'r' {
		fmt.Printf("%.0f\n", math.Round(x))
	} else {
	
	}
}
