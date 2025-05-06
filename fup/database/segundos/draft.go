package main

import "fmt"

func main() {
	var total int
	fmt.Scan(&total)

	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60

	fmt.Printf("%d:%d:%d\n", h, m, s)

}
