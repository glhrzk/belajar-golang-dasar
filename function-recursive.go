package main

import "fmt"

func manual(value int) int {
	count := 1
	for i := value; i > 0; i-- {
		count *= i
	}
	return count
}

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1)
	}
}

func main() {

	value1 := manual(5)
	value2 := 5 * 4 * 3 * 2 * 1
	otomatis := recursive(5)

	fmt.Println("value1: ", value1)
	fmt.Println("value2: ", value2)
	fmt.Println("otomatis: ", otomatis)
}
