package main

import "fmt"

func main() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 165*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 24*365/687)
	fmt.Println("-------------------------")
	fmt.Printf("%-15v |%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v |%4v\n", "Virgin Galactic", 100)
	fmt.Println("-------------------------")
}
