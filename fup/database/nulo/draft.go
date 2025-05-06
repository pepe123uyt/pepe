package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)

    switch {
    case num > 0:
        fmt.Printf("+\n")
    case num < 0:
        fmt.Printf("-\n")
    default:
        fmt.Printf("0\n")
    }
}
