package main
import "fmt"
func main() {
    var b, t int 
    fmt.Scan(&b, &t)

    trapesio := ((b + t) * 70) / 2
    media_nota := 160 * 70 / 2
    if trapesio > media_nota && b < t {
        fmt.Println("1")
    } else if trapesio > media_nota && b > t {
        fmt.Println("2")
    } else {
        fmt.Println("0") 
    }
}

