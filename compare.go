package main

import "fmt"

func main() {
	fmt.Println("There is a sign here that reads 'No Minors'.")

	var age = 39
	var minor = age < 18

	fmt.Printf("At the age of %v, am I a minor? %v\n", age, minor)
}
