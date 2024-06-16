package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Even odd...")
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is even")
	}

	fmt.Println()

	fmt.Println("Fizzbuzz...")
	for i := 0; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}

	fmt.Println()
	dayBorn := time.Monday
	fmt.Println("Born day...")
	bornDay(dayBorn)

	fmt.Println()
	fmt.Println("Range over map...")
	rangeOverMap()

	fmt.Println()
	fmt.Println("Bubble sort...")
	arr := []int{3, 5, 2, 6, 8, 1}
	fmt.Println("Array before sorting: ", arr)
	bubbleSort(arr)
	fmt.Println("Array after sorting: ", arr)
}
