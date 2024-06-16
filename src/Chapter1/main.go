package main

import (
	"fmt"
)

func main() {
	fmt.Println("Random message...")
	randomMessgae()

	fmt.Println()

	fmt.Println("Print stars...")
	printStars()

	fmt.Println()

	fmt.Println("Defining vars...")
	medicalForm()

	fmt.Println()

	fmt.Println("Multiple return values...")
	debug, logLevel, startupTime := getConfig()
	fmt.Println(debug, logLevel, startupTime)

	fmt.Println()

	fmt.Println("Function with pointer argument...")
	var count int
	add5Value(count)
	fmt.Println("add5Value post: ", count)
	add5Point(&count)
	fmt.Println("add5Value post: ", count)

	fmt.Println()
	fmt.Println("swapping values...")
	a, b := 10, 20
	fmt.Printf("before swap, a: %v, b: %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap, a: %v, b: %v\n", a, b)

	fmt.Println()
	fmt.Println("after fixing message bug...")
	messageBug()
}
