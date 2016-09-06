package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := rand.Intn(2016) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31

	if year%400 == 0 || year%4 == 0 && year%100 == 0 {
		switch month {
		case 2:
			daysInMonth = 29
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
	} else if year%400 != 0 || year%4 != 0 && year%100 != 0 {
		switch month {
		case 2:
			daysInMonth = 28
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}