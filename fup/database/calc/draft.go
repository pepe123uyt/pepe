package main
import "fmt"
func soma( x, y int)int {
    resultado := x + y
    return resultado
}
func sub(x, y int)int {
    resultado := x - y
    return resultado
}
func multi(x, y int)int {
    resultado := x * y
    return resultado
}
func divi(x, y int)int {
    resultado := x / y
    return resultado
    }
func main() {

    var x, y int
    fmt.Scan(&x, &y)
    var operador string
    fmt.Scan(&operador)

    switch operador {
    case "+":
        fmt.Print("%d\n", soma(x, y))
    
    case "-":
        fmt.Printf("%d\n", sub(x, y))
    
    case "*":
        fmt.Printf("%d\n", multi(x, y))
    
    case "/":
        fmt.Printf("%d\n", divi(x, y))
    



    }

 
}
