package main

import "fmt"

func main() {

	name := "galih rizkiansyah"

	// long statement
	switch name {
	case "bunda":
		fmt.Println("Hello galih")
		fmt.Println("Hello galih")
	case "joko":
		fmt.Println("Hello joko")
		fmt.Println("Hello joko")
	default:
		fmt.Println("Hi")
	}

	// short statement
	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah sesuai")
	}

	// tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang, ", len(name))
	case length > 5:
		fmt.Println("Nama lumayan panjang, ", len(name))
	default:
		fmt.Println("Nama sudah benar,", len(name))

	}
}
