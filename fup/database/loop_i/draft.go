package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    for a := a; a < b; a += 1{
        fmt.Println(a)
    }
}
