package main
import "fmt"
func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if (a > b && a < c) || (a < b && a > c)  {
		fmt.Printf("%d\n", a)
	} else if (b > a && b < c) || (b < a && b > c)  {
		fmt.Printf("%d\n", b)
	} else {
		fmt.Printf("%d\n", c)
	}
}
