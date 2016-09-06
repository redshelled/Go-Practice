package main

import "fmt"

func main() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("You need to head further up the mountains.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live your life.")
	} else {
		fmt.Println("Didnt quite get that...")
	}
}
