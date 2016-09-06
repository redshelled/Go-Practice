package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 92
	var guess = 0

	for {
		var randNum = rand.Intn(100)

		if randNum == number {
			fmt.Print("Number guessed is ")
			fmt.Print(randNum)
			break
		} else {
			guess++
			fmt.Print("I am on guess number ")
			fmt.Print(guess)
			fmt.Println(" and still guesssing...")
		}
	}

}
