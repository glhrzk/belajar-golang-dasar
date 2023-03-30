package main

import "fmt"

func main() {
	// array tanpa inisiasi
	var names [3]string
	names[0] = "galih"
	names[1] = "rizki"
	names[2] = "galih rizki"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array dengan deklarasi
	var values = [3]int{
		90,
		51,
		90,
	}

	fmt.Println(values)

	// count array
	fmt.Println(len(names))
	fmt.Println(len(values))
}
