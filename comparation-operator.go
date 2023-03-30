package main

import "fmt"

func main() {
	galih1 := "galih"
	galih2 := "galih"

	var result bool = galih1 == galih2

	fmt.Println(galih1)
	fmt.Println(galih2)
	fmt.Println(result)

	value1 := 100
	value2 := 200
	fmt.Println("value1 < value2: ", value1 < value2)
	fmt.Println("value1 > value2: ", value1 > value2)
	fmt.Println("value1 == value2: ", value1 == value2)
	fmt.Println("value1 != value2: ", value1 != value2)
}
