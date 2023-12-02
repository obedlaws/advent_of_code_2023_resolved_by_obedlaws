package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dayOneFunc() (int, error) {
	stringsOfCodes, errCodes := getCodes()
	convertedNumbers := []int{}

	if errCodes != nil {
		return 0, errCodes
	}

	for _, v := range stringsOfCodes {

		localNumbers := []string{}
		array := strings.Split(v, "")
		for _, s := range array {
			if checkIfNumber(s) {
				localNumbers = append(localNumbers, s)
			}
		}
		realNumber, errNumberConverter := getFinalNumber(localNumbers)
		if errNumberConverter != nil {
			return 0, errNumberConverter
		}

		convertedNumbers = append(convertedNumbers, realNumber)
	}

	totalSum := addAllNumbers(convertedNumbers)

	return totalSum, nil

}

func getCodes() ([]string, error) {
	content, err := os.ReadFile("dayOneCodes.txt")
	if err != nil {
		return []string{}, err
	}

	codes := strings.Split(string(content), "\n")
	return codes, nil

}

func checkIfNumber(inputNumber string) bool {
	numbersSlice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	for _, n := range numbersSlice {
		if inputNumber == n {
			return true
		}
	}

	return false
}

func getFinalNumber(sliceOfNumbers []string) (int, error) {

	firstNumber := sliceOfNumbers[0]
	secondNumber := sliceOfNumbers[len(sliceOfNumbers)-1]

	stringNumber := fmt.Sprintf("%s%s", firstNumber, secondNumber)

	number, err := strconv.Atoi(stringNumber)
	if err != nil {
		return 0, errors.New("failed to convert string into number")
	}

	return int(number), nil
}

func addAllNumbers(sliceOfCodes []int) int {
	totalSum := 0

	for _, v := range sliceOfCodes {
		totalSum = totalSum + v
	}

	return totalSum
}
