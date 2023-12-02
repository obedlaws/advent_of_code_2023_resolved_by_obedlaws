package main

import (
	"log"
	"regexp"
	"strings"
)

var one = regexp.MustCompile("one")
var two = regexp.MustCompile("two")
var three = regexp.MustCompile("three")
var four = regexp.MustCompile("four")
var five = regexp.MustCompile("five")
var six = regexp.MustCompile("six")
var seven = regexp.MustCompile("seven")
var eight = regexp.MustCompile("eight")
var nine = regexp.MustCompile("nine")
var zero = regexp.MustCompile("zero")

type numberInfo struct {
	number       string
	initialIndex int
	finalIndex   int
}

func dayOnePartTwoFunc() {
	code := "2911threeninesdvxvheightwobm"
	// sliceOfValues := strings.Split(code, "three")
	// re := regexp.MustCompile("two")
	// numbers := regexp.MustCompile("1|2|3|4|5|6|7|8|9|0")

	if strings.Contains(code, "two") {
		res := regexp.MustCompile("two").FindAllStringIndex(code, -1)
		var newNumber = numberInfo{
			"two", res[0][0], res[0][1],
		}

		log.Println(newNumber)
	}

	// matches := re.FindAllStringIndex(code, -1)
	// allNumbers := numbers.FindAllStringIndex(code, -1)

	// log.Print(matches)
	// log.Println(allNumbers)
	// log.Print(string(code[matches[0][0]:matches[0][1]]))

	// for i, v := range sliceOfValues {
	// 	infoOfSlices[v] = i
	// }

	// res := ifNumberInString(code, infoOfSlices)
	// log.Print(infoOfSlices)
}

// func seperateLetters(code []string) []string {
// 	infoOfLetters :=

// 	for i, v := range code {

// 	}

// 	return []string{}

// }

// func ifNumberInString(code string) [][]int {
// 	matches := []numberInfo{}

// 	if strings.Contains(code, "one") {
// 		getOne := one.FindAllStringIndex(code, -1)
// 		matches = append(matches, numberInfo{
// 			"one",
// 			getOne[0][0],
// 			getOne[0][0],
// 		})
// 	}

// 	return matches

// }
