package main
import "fmt"
import "math"
func main() {
    var xa, ya, xb, yb float64
    fmt.Scan(&xa, &ya, &xb, &yb)

    dista := math.Sqrt(((xa-xb)*(xa-xb)) + ((ya-yb)*(ya-yb)))

    fmt.Printf("%.2f\n", dista)

}
