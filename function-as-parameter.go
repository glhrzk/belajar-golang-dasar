package main

import "fmt"

// type declaration
type Filter func(string2 string) string

// parsing function to function
func sayHelloWithFiller(name string, filter Filter) string {
	filteredName := filter(name)
	return "Hello " + filteredName
}

func filter(filterName string) string {
	if filterName == "anjing" {
		return "***"
	} else {
		return filterName
	}
}

func main() {

	filter := sayHelloWithFiller("anjing", filter)
	fmt.Println(filter)
}
