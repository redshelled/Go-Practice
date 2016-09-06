package main

import "fmt"

func main() {
	var room = "cave"

	if room == "cave" {
		fmt.Println("The room is big and empty.")
	} else if room == "entrance" {
		fmt.Println("The entrance is tall and wide.")
	} else {
		fmt.Println("The mountain is tall.")
	}
}
