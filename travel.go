package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km
	var speed = 16            // km/s

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	fmt.Println(distance/speed/86400, "days")

	distance = 57600000
	fmt.Println(distance/speed/86400, "days")
}
