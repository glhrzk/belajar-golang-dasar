package main

import "fmt"

func main() {

	name := "ujang"

	if name == "galih" {
		fmt.Println("Hello galih")
	} else if name == "ujang" {
		fmt.Println("Woiy ujang")
	} else {
		println("Hi, boleh kenalan?")
	}

	// long statement
	shortState := len(name)
	if shortState > 6 {
		fmt.Println("nama sudah benar")
	} else {
		fmt.Println("nama terlalu panjang")
	}

	// short statement
	if longState := len(name); longState > 6 {
		fmt.Println("nama sudah benar")
	} else {
		fmt.Println("nama terlalu panjang")
	}
}
