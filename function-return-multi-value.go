package main

import "fmt"

func getFullName() (string, string) {
	return "galih", "rizki"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// ignore salah satu return value
	personFirstName, _ := getFullName()
	fmt.Println(personFirstName)

}
