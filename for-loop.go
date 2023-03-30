package main

import "fmt"

func main() {
	value := 1

	for value < 10 {
		fmt.Println("belum terpenuhi", value)
		value++
	}
	fmt.Println("terpenuhi", value)

	// for dengan statement
	for i := 1; i < 10; i++ {
		fmt.Println("belum terpenuhi", i)
	}

	names := []string{"galih", "budi", "ujang", "komar"}
	person := map[string]string{
		"name":    "galih",
		"address": "tangerang",
	}

	// for manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range atau bisa disebut dengan foreach, with key and value
	for index, name := range names {
		fmt.Println("index ke-", index, " = ", name)
	}

	// for range only value
	for _, name := range names {
		fmt.Println(name)
	}

	for i, key := range person {
		fmt.Println(i, key)
	}
}
