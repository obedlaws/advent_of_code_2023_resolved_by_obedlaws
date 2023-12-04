package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type gameDraw struct {
	GreenCubes int
	BlueCubes  int
	RedCubes   int
}

type gameInfo struct {
	GameNumber int
	IsPossible bool
	Draws      []gameDraw
}

func dayTwoPartOneFunc() (int, error) {
	listOfGames, errGetGames := getGames()
	listOfGamesInfo := []gameInfo{}
	if errGetGames != nil {
		return 0, errGetGames
	}

	for i, v := range listOfGames {
		game := getGameInformation(v, i+1)
		listOfGamesInfo = append(listOfGamesInfo, game)
	}

	for i, v := range listOfGamesInfo {
		res := checkIfGame(&v)
		v.IsPossible = res
		listOfGamesInfo[i] = v
	}

	n := sumOfAllIds(listOfGamesInfo)
	return n, nil
}

func getGames() ([]string, error) {
	content, err := os.ReadFile("./dayTwoInput.txt")
	if err != nil {
		return []string{}, err
	}

	codes := strings.Split(string(content), "\n")
	return codes, nil

}

func getGameInformation(game string, gameId int) gameInfo {
	currentGame := gameInfo{}

	firstSplit := strings.Split(game, ":")
	secondSplit := strings.Split(firstSplit[1], ";")

	// log.Println(firstSplit)
	// log.Println(secondSplit[0])

	noCommaSlice := [][]string{}
	for _, v := range secondSplit {
		thirdSplit := strings.Split(v, ",")
		noCommaSlice = append(noCommaSlice, thirdSplit)
	}
	for _, v := range noCommaSlice {
		// log.Printf("current slice in noCommasSlice: %s", v)
		currentDraw := gameDraw{}

		for _, v2 := range v {
			// log.Printf("current item inside of the slice: %s", v2)
			if strings.Contains(v2, "blue") {
				res := regexp.MustCompile("blue").FindAllStringIndex(v2, -1)
				number := numberConverter(v2[:res[0][0]])
				// log.Printf("current number converted: %v", number)
				currentDraw.BlueCubes = number
			}
			if strings.Contains(v2, "red") {
				res := regexp.MustCompile("red").FindAllStringIndex(v2, -1)
				number := numberConverter(v2[:res[0][0]])
				// log.Printf("current number converted: %v", number)
				currentDraw.RedCubes = number
			}
			if strings.Contains(v2, "green") {
				res := regexp.MustCompile("green").FindAllStringIndex(v2, -1)
				number := numberConverter(v2[:res[0][0]])
				// log.Printf("current number converted: %v", number)
				currentDraw.GreenCubes = number
			}
		}
		currentGame.Draws = append(currentGame.Draws, currentDraw)
	}

	currentGame.GameNumber = gameId
	currentGame.IsPossible = true

	// log.Println(currentGame)

	return currentGame
}

func checkIfGame(game *gameInfo) bool {

	for _, v := range game.Draws {
		if v.BlueCubes > 14 {
			game.IsPossible = false
			return false
		} else if v.GreenCubes > 13 {
			game.IsPossible = false
			return false
		} else if v.RedCubes > 12 {
			game.IsPossible = false
			return false
		}
	}

	return true
}

func sumOfAllIds(allGames []gameInfo) int {
	allIds := []int{}
	number := 0
	for _, v := range allGames {
		if v.IsPossible == true {
			allIds = append(allIds, v.GameNumber)
		}
	}

	for _, v := range allIds {
		number = number + v
	}

	return number
}

func numberConverter(number string) int {
	results := regexp.MustCompile(" ").ReplaceAllLiteralString(number, "")
	// log.Println(results)

	n, err := strconv.Atoi(results)
	if err != nil {
		log.Panic(err)
	}

	// log.Println(n)
	return n
}
