package main
import "fmt"
func main() {
    var salario float64
    fmt.Scan(&salario) 
    
    if salario <= 1000 {
        fmt.Printf("%.2f\n", salario * 1.2)
    } else if salario <= 1500 {
        fmt.Printf("%.2f\n", salario * 1.15)
    } else if salario <= 2000 {
        fmt.Printf("%.2f\n", salario * 1.1)
    } else if salario > 2000 {
        fmt.Printf("%.2f\n", salario * 1.05)
    }
}
