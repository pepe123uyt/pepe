package main
import "fmt"

func main() {
	var c1, c2 string
	fmt.Scan(&c1, &c2)

	if c1 == c2 {
		fmt.Println("empate")
	} else if (c1 == "R" && c2 == "S") || 
	          (c1 == "S" && c2 == "P") || 
	          (c1 == "P" && c2 == "R") {
		fmt.Println("jog1")
	} else {
		fmt.Println("jog2")

	
}
}