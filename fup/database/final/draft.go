package main
import "fmt"
func main() {
    var nota1, nota2, nota_final int
    fmt.Scan(&nota1, &nota2)

    media := (nota1 + nota2) / 2
    
    if media >= 7{
        fmt.Println("aprovado")
    } else if media < 4{
        fmt.Println("reprovado")
    } else if media >= 4 && media < 7 {
        fmt.Scan(&nota_final)
        if nota_final > 5 {
            fmt.Println("aprovado na final")
        } else {
            fmt.Println("reprovado na final")
        }
    }
}



