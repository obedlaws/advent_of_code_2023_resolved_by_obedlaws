package main

import "log"

func main() {
	dayTwoRes, err := dayTwoPartOneFunc()
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Sum of all ids in the games that combination was possible: %v", dayTwoRes)

	dayTwoPartTwoRes, errPartTwo := dayTwoPartTwoFunc()
	if errPartTwo != nil {
		log.Panic(err)
	}

	log.Printf("Sum of all the games multiplications: %v", dayTwoPartTwoRes)

}
