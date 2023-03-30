package main

import "fmt"

func main() {
	// map adalah array dengan key name, seperti di PHP
	person := map[string]string{
		"name":    "galih rizki",
		"address": "tangerang",
	}

	person["age"] = "23"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	// deklarasi map tanpa isi
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Galih rizki"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
