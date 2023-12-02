package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type numberInfo struct {
	number       string
	initialIndex int
}

func dayOnePartTwoFunc() (int, error) {
	listOfCodes, errCodes := getCodes()
	listOfAllNumbers := []int{}

	if errCodes != nil {
		return 0, errCodes
	}

	for _, v := range listOfCodes {

		log.Println(v)
		listOfCharacters := characterIdentifier(v)

		firstNumberString := firstIndexChecker(listOfCharacters)
		lastNumberString := lastIndexChecker(listOfCharacters)

		convFirstNumber := numberConverter(firstNumberString.number)
		convLastNumber := numberConverter(lastNumberString.number)

		stringNumber := fmt.Sprintf("%s%s", convFirstNumber, convLastNumber)
		intNumber, errConverter := strconv.Atoi(stringNumber)
		if errConverter != nil {
			return 0, errConverter
		}
		log.Println(convFirstNumber, convLastNumber)
		listOfAllNumbers = append(listOfAllNumbers, int(intNumber))

	}

	sum := addAllNumbers(listOfAllNumbers)

	return sum, nil

}

func characterIdentifier(code string) []numberInfo {
	listOfNumberInfo := []numberInfo{}
	list := [][]int{}
	regexList := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, v := range regexList {
		res := regexp.MustCompile(v).FindAllStringIndex(code, -1)
		if len(res) != 0 {
			for _, v := range res {
				list = append(list, v)
			}
		}
	}

	for _, v := range list {
		listOfNumberInfo = append(listOfNumberInfo, numberInfo{
			code[v[0]:v[1]], v[0],
		})
	}

	return listOfNumberInfo
}

func numberConverter(number string) string {
	selectedNumber := ""

	switch number {
	case "one":
		selectedNumber = "1"
	case "two":
		selectedNumber = "2"
	case "three":
		selectedNumber = "3"
	case "four":
		selectedNumber = "4"
	case "five":
		selectedNumber = "5"
	case "six":
		selectedNumber = "6"
	case "seven":
		selectedNumber = "7"
	case "eight":
		selectedNumber = "8"
	case "nine":
		selectedNumber = "9"
	case "zero":
		selectedNumber = "0"
	default:
		selectedNumber = number
	}

	return selectedNumber
}

func lastIndexChecker(arr []numberInfo) numberInfo {
	lastIndex := arr[0].initialIndex
	indexOfLastNumber := 0

	for i, v := range arr {
		if lastIndex < v.initialIndex {
			indexOfLastNumber = i
			lastIndex = v.initialIndex
		}
	}

	return arr[indexOfLastNumber]
}

func firstIndexChecker(arr []numberInfo) numberInfo {
	firstIndex := arr[0].initialIndex
	indexOfFirstNumber := 0

	for i, v := range arr {
		if firstIndex > v.initialIndex {
			indexOfFirstNumber = i
			firstIndex = v.initialIndex
		}
	}

	return arr[indexOfFirstNumber]
}
