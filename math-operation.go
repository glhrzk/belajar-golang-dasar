package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := a * b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// augmented assignment
	i := 10
	fmt.Println(i)
	i += 10
	fmt.Println(i)

	// Unary operator
	i++
	fmt.Println(i)

	negative := -100
	positive := 100

	fmt.Println(negative)
	fmt.Println(positive)
}
