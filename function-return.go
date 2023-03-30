package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	name := "galih"

	fmt.Println(getHello(name))
}
