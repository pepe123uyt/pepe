package main
import "fmt"
func main() {
var x, y, z int
fmt.Scan(&x, &y, &z)

if x != y && y == z {
    fmt.Println("2")
}   else if x == y && y != z {
    fmt.Println("2")
}   else if z != y && y != x && z != x {
    fmt.Println("0")
}   else if x == z && x != y{
    fmt.Println("2")
}   else {
    fmt.Println("3")
}




}
