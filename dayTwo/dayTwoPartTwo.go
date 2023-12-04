package main

func dayTwoPartTwoFunc() (int, error) {
	// 1. get the information of each games
	listOfGames, errGetGames := getGames()
	listOfGamesInfo := []gameInfo{}
	listOfResults := []int{}

	if errGetGames != nil {
		return 0, errGetGames
	}

	for i, v := range listOfGames {
		game := getGameInformation(v, i+1)
		listOfGamesInfo = append(listOfGamesInfo, game)
	}
	// 2. find the least amount of each block in each games and also multiply it and add it to
	// struct

	for _, v := range listOfGamesInfo {
		res := getMultplication(v.Draws)

		listOfResults = append(listOfResults, res)
	}

	res := 0
	for _, v := range listOfResults {
		res = res + v
	}

	// 3. add up all the games multiplications

	return res, nil
}

func getMultplication(currentGame []gameDraw) int {
	leastBlue := 0
	leastRed := 0
	leastGreen := 0

	for _, v := range currentGame {
		if leastBlue < v.BlueCubes {
			leastBlue = v.BlueCubes
		}
		if leastRed < v.RedCubes {
			leastRed = v.RedCubes
		}
		if leastGreen < v.GreenCubes {
			leastGreen = v.GreenCubes
		}
	}

	return leastBlue * leastRed * leastGreen
}
