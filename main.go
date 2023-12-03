package main

import (
	dayone "advent/dayOne"
	"log"
)

func main() {

	//Day One - Part 1
	dayOneResults, errDayOne := dayone.DayOneFunc()
	if errDayOne != nil {
		log.Panic(errDayOne)
	}

	log.Printf("Day One - Part One: Get the sum of all codes to calibrate the Trebuchet?!: %v", dayOneResults)

	//Day One - Part 2
	dayOnePartTwoResults, errDayOnePartTwo := dayone.DayOnePartTwoFunc()
	if errDayOnePartTwo != nil {
		log.Panic(errDayOnePartTwo)
	}

	log.Printf("Day One - Part Two: Get the sum of all codes but from the letters. Calibrate BETTER: %v", dayOnePartTwoResults)
}
