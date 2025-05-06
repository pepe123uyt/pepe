package main
import "fmt"
func main() {
    var m, a ,b int
    fmt.Scan(&m, &a, &b)

    c := m - (a + b)
    if c > b && c > a {
        fmt.Println(c)
    } else if b > a && b > c {
        fmt.Println(b)
    } else {
        fmt.Println(a)
    }
}
