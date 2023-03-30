package main

import "fmt"

func main() {
	// deklarasi variable dengan type data
	var name string
	name = "Galih rizki"
	fmt.Println(name)

	// deklarasi variable tanpa type data
	var friendName = "clover"
	fmt.Println(friendName)

	// penulisan cepat dari variable
	country := "Indonesia"
	fmt.Println(country)

	// multi variable
	var (
		firstName = "galih"
		lastName  = "rizki"
	)
	fmt.Println(firstName, lastName)

}
