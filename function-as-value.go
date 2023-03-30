package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	name := getGoodBye
	fmt.Println(name("clover"))
}
