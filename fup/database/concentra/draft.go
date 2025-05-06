package main
import "fmt"
func main() {
    var ini, fim int
    fmt.Scan(&ini, &fim)
    fmt.Print("[ ")
    fimb := fim
    for {        
        fmt.Printf("%v %v ", ini, fim)
        ini += 1 
        fim -= 1
        if ini > fimb{
            break 
        }
    }
        
        
    fmt.Println("]")
}
