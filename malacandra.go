package main

import "fmt"

func main() {
	var distance = 56000000 // km
	var timeTravel = 28     // days
	const secDayConvert = 86400

	fmt.Println(distance/(timeTravel*secDayConvert), "km/s")
}
