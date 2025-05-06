package main
import "fmt"
func maior(x int, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func main() {
var x, y int
fmt.Scan(&x, &y)

fmt.Printf("%d\n", maior(x, y))
}
