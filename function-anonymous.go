package main

import "fmt"

type Blocklist func(fillter string) bool

func registerUser(name string, blocklist Blocklist) string {
	if blocklist(name) {
		return "You are blocked ***"
	} else {
		return "Welcome " + name
	}
}

func main() {
	blockList := registerUser("anjing", func(fillter string) bool {
		return fillter == "anjing"
	})
	fmt.Println(blockList)
}
