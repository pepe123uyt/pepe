package main
import "fmt"

func abs(a int, b int) int {
    if a < b {
        return (a - b) * (-1)
    }
    if a > b {
        return (a - b)
    }

    return 0
}

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    fmt.Printf("%d\n", abs (a, b))

}
