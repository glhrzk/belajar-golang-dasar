package main

import "fmt"

func main() {
	// type declaration adalah kemampuan membuat alias dari type variable
	type noKTP string
	type married bool

	var ktpGalih noKTP = "16123123123"
	fmt.Println(ktpGalih)
}
