package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a cave.")
	case room == "lake":
		fmt.Println("The ice seems solid.")
	case room == "underwater":
		fmt.Println("The water is cold.")
	}
}
