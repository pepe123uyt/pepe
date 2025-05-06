package main
import "fmt"
func main() {
	var p, d1, d2 int
	fmt.Scan(&p, &d1, &d2)

	soma := d1 + d2

	if p == 0 && soma % 2 == 0 {
		fmt.Println("0")
	} else if p == 1 && soma % 2 != 0 {
		fmt.Println("0")
	}	else {
		fmt.Println("1")
	}
}
