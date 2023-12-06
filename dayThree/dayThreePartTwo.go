package main

var lines []string
var errGetCodes error

func dayThreePartTwoFunc() (int, error) {
	lines, errGetCodes = getCodes()
	if errGetCodes != nil {
		return 0, errGetCodes
	}

	characters := createCoordinates(lines)

	for row := 1; row <= 2; row++ {
		getNumbersOperations(row, characters)
	}

	return 0, nil
}

func getNumbersOperations(row int, coor map[int]map[int]characterCoordinates) {
	for line := 1; line <= len(coor[row]); line++ {
		if coor[row][line].Type == "Asterisk" {
			// log.Print(row-1, line-1, string(lines[row-1][line-1]))
			checkForNumbers(row-1, line-1)

		}
	}
}

func checkForNumbers(startingRow int, startingIndex int) {
	lineAbove := string(lines[startingRow-1][startingIndex-1 : startingIndex+2])
	currentLine := string(lines[startingRow][startingIndex-1 : startingIndex+2])
	lineBelow := string(lines[startingRow+1][startingIndex-1 : startingIndex+2])

}
