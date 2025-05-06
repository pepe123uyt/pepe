package main
import "fmt"
func main() {

var num int
fmt.Scan(&num)

switch {
case num == 3:
    fmt.Printf("SIM\n")
case num == 5:
    fmt.Printf("SIM\n")
default:
    fmt.Printf("NAO\n")

}
}
