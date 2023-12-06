package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type characterCoordinates struct {
	Value string
	Type  string
}

func dayThreePartOne() (int, error) {
	listOfCodes, errGetCodes := getCodes()
	possibleNumbers := [][]string{}
	convertedNumbers := []int{}
	var errConvert error

	if errGetCodes != nil {
		return 0, errGetCodes
	}

	// 2. set coordinates of each of the characters
	mapOfCoordinates := createCoordinates(listOfCodes)
	// log.Println(mapOfCoordinates)
	for i := 1; i <= len(mapOfCoordinates); i++ {
		// log.Println(mapOfCoordinates[1])
		res := getNumbers(i, mapOfCoordinates)

		possibleNumbers = append(possibleNumbers, res...)

		// log.Println(possibleNumbers)
	}

	convertedNumbers, errConvert = convertToNumbers(possibleNumbers)
	if errConvert != nil {
		return 0, errConvert
	}

	// log.Print(len(convertedNumbers))

	// log.Println(convertedNumbers)

	// 3. findout which numbers are adjacent to
	sum := addAll(convertedNumbers)

	log.Println(sum)

	return sum, nil
}

func getCodes() ([]string, error) {
	res, err := os.ReadFile("./testInput.txt")
	if err != nil {
		return nil, err
	}

	slice := strings.Split(string(res), "\n")
	// for _, v := range slice {
	// 	log.Println(v)
	// }

	return slice, nil
}

func convertToNumbers(numbers [][]string) ([]int, error) {
	converted := []int{}

	for _, v := range numbers {
		var b bytes.Buffer
		for _, c := range v {
			b.WriteString(c)

		}
		n := b.String()

		r, err := strconv.Atoi(n)
		if err != nil {
			return []int{}, err
		}
		converted = append(converted, r)
	}

	return converted, nil
}

func addAll(num []int) int {
	sum := 0

	for _, v := range num {
		sum = sum + v
	}

	return sum

}

func getNumbers(row int, currentMap map[int]map[int]characterCoordinates) [][]string {
	number := [][]string{}

	for lat := 1; lat <= 140; lat++ {
		// log.Println("index of characters", lat)
		// log.Println(lat)
		if currentMap[row][lat].Type == "Number" {
			// log.Print("number found", row, lat)
			finalIndex, numberChars := checkNextDigits(lat, row, currentMap)
			results := checkSorrounding(row, lat, finalIndex, currentMap)

			if results == true {
				number = append(number, numberChars)
			}

			lat = finalIndex + 1
			// log.Println(lat)
		}
	}

	return number
}

func checkNextDigits(currentIndex int, currentRow int, currentMap map[int]map[int]characterCoordinates) (int, []string) {
	localIndex := currentIndex
	localRow := currentRow
	// log.Print(localIndex)
	number := []string{}

	for localIndex <= len(currentMap[localRow]) {
		// log.Println("finding the next number in row:", localRow)
		// log.Println(localIndex)
		// log.Println(currentMap[localRow][localIndex].Value)
		if currentMap[localRow][localIndex].Type == "Asterisk" || currentMap[localRow][localIndex].Type == "Dot" || currentMap[localRow][localIndex].Type == "Symbol" {
			// log.Println("found end of number: ", localIndex-1, number)
			// log.Println(localIndex) //index final
			// log.Println(number)
			return localIndex - 1, number
		}
		// log.Println("outside funciton where is dot or symbol", localIndex)
		if localIndex == len(currentMap[localRow]) {
			number = append(number, currentMap[localRow][localIndex].Value)
			return localIndex, number
		}
		number = append(number, currentMap[localRow][localIndex].Value)

		localIndex++
	}

	return 0, []string{}
}

func checkSorrounding(long int, firstIndex int, lastIndex int, currentMap map[int]map[int]characterCoordinates) bool {
	startingScanIndex := 0
	endingScanIndex := 0

	// log.Println(long)
	// log.Println("starting")
	if firstIndex == 1 {
		startingScanIndex = 1
	} else if firstIndex > 1 {
		startingScanIndex = firstIndex - 1
	}

	if lastIndex == len(currentMap[long]) {
		endingScanIndex = len(currentMap[long])
	} else if lastIndex < len(currentMap[long]) {
		endingScanIndex = lastIndex + 1
	}

	// log.Println(startingScanIndex)
	// log.Println(endingScanIndex)

	for i := startingScanIndex; i <= endingScanIndex; i++ {
		// log.Println("started for loop for validation...", i)
		if long > 1 {
			// log.Println("checked if row before is possible")

			if currentMap[long-1][i].Type == "Symbol" {
				// log.Println("symbol found in row before")
				// log.Println(currentMap[long-1][i].Type)
				return true
			}
		}

		//check below
		if long < len(currentMap) {
			// log.Println("checked if row after is possible")
			// log.Println(currentMap[long+1][i].Type)
			if currentMap[long+1][i].Type == "Symbol" {
				// log.Println("symbol found in row after")
				// log.Println(currentMap[long-1][i].Type)

				return true
			}
		}

	}

	// check sides

	if firstIndex != 1 {
		// log.Println("checked if indexto the right is possible")
		if currentMap[long][firstIndex-1].Type == "Symbol" {
			// log.Println("checked if is symbol")
			// log.Println(currentMap[long][firstIndex-1])

			return true
		}
	}

	if lastIndex != len(currentMap[long]) {
		// log.Println("checked if index to the left is possible")

		if currentMap[long][lastIndex+1].Type == "Symbol" {
			// log.Println("checked if is symbol")
			// log.Println(currentMap[long][firstIndex-1])
			return true
		}
	}

	return false
}

func createCoordinates(codes []string) map[int]map[int]characterCoordinates {
	mapOfCoordinates := make(map[int]map[int]characterCoordinates)

	for lon, v := range codes {
		slice := strings.Split(v, "")
		currentCharacter := make(map[int]characterCoordinates)
		for lat, c := range slice {
			currentCharacter[lat+1] = characterCoordinates{
				c, characterChecker(c),
			}
		}

		mapOfCoordinates[lon+1] = currentCharacter
	}

	return mapOfCoordinates
}

func characterChecker(character string) string {
	number := regexp.MustCompile("1|2|3|4|5|6|7|8|9|0").FindAllString(character, -1)
	if number != nil {
		return "Number"
	}

	if character == "*" {
		return "Asterisk"
	}

	if character == "." {
		return "Dot"
	}

	return "Symbol"
}
