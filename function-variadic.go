package main

import "fmt"

// sama seperti di php

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	count := sumAll(50, 10, 20, 10, 10)
	fmt.Println(count)

	// variadic function di slice
	sliceCount := []int{10, 10, 20, 20, 30, 30, 40, 40, 50, 50}
	count = sumAll(sliceCount...)
	fmt.Println(count)
}
