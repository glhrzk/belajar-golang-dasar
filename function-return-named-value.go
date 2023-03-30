package main

import "fmt"

// melakukan return value secara sekaligus
func getFullName2() (firstName, lastName string) {
	firstName = "galih"
	lastName = "rizki"
	return
}

func main() {
	a, b := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
}
