package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)
switch {
    case num % 7 == 0:
    fmt.Printf("SIM\n")
    default: {
    fmt.Printf("NAO\n")
}
    }

}