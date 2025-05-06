package main
import "fmt"
func main() {
    
    var cel float64
    fmt.Scan(&cel)
    
    fahr :=  cel * 1.8 + 32
    
    fmt.Printf("%.6f\n", fahr)


}
