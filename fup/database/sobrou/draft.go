package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var va, vb, vc float64
	fmt.Scan(&va, &vb, &vc)
	var dinheiro float64
	fmt.Scan(&dinheiro)

	total := float64(a)*va + float64(b)*vb + float64(c)*vc

	fmt.Printf("%.2f\n", dinheiro-total)

}