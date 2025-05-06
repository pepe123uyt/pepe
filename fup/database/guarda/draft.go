package main

import (
	"fmt"
)

func main() {
	var wifi, login, admin bool

	fmt.Scan(&wifi)
	fmt.Scan(&login)
	fmt.Scan(&admin)

	if !wifi{
		fmt.Println("you must connect to wifi")
	} else if !login {
		fmt.Println("you need to login first")
	} else if !admin {
		fmt.Println("you must to login as admin")
	} else {
		fmt.Println("done")
	}
}
