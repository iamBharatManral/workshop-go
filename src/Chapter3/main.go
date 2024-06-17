package main

import "fmt"

func main() {
	password := "Hello123$"
	if passwordChecker(password) {
		fmt.Println("password is valid")
	} else {
		fmt.Println("password is not valid")
	}
}
