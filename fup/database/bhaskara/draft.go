package main
import "fmt"
import "math"
func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)

    delta := (b * b) - (4 *a * c)

    if delta > 0 {
        fmt.Printf("%.2f\n", (-b + math.Sqrt(delta)) / (2 * a))
        fmt.Printf("%.2f\n", (-b - math.Sqrt(delta)) / (2 * a))
    } else if delta == 0 && b == 0 {
        fmt.Printf("%.2f\n", b / (2 * a))
    } else if delta == 0 {
        fmt.Printf("%.2f\n", -b / (2 * a))
    } else {
        fmt.Println("nao ha raiz real")
    }
}
