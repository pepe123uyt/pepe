package main
import "fmt"
func main() {
    var inf, sup int
    fmt.Scan(&inf, &sup)
    var soma int
    if inf > sup {
        fmt.Println("invalido")
        return
    }
    for i := inf; i < sup + 1; i++ {
        if i % 2 == 0 {
            soma += 1
        }
    }
}
