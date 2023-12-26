package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	smallInputFile = "../input1.txt"
	inputFile      = "../input.txt"
)

var (
	useActualInput bool
	bagContent     map[string]int
)

func init() {
	useActualInput = true
	bagContent = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sum := 0
	inputFileToUse := smallInputFile
	if useActualInput == true {
		inputFileToUse = inputFile
	}
	f, err := os.Open(inputFileToUse)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var possibleGames []string
	for scanner.Scan() {
		lineStr := scanner.Text()
		lineSplits := strings.Split(lineStr, ": ")
		gameID, gameCubes := lineSplits[0], lineSplits[1]
		gameCubesSplits := strings.Split(gameCubes, "; ")
		gamePossible := true
		for _, cubes := range gameCubesSplits {
			cubeSplits := strings.Split(cubes, ", ")
			for _, cubeConf := range cubeSplits {
				cubeConfSplits := strings.Split(cubeConf, " ")
				cubeCount, _ := strconv.Atoi(cubeConfSplits[0])
				cubeColor := cubeConfSplits[1]
				if cubeCount > bagContent[cubeColor] {
					gamePossible = false
					break
				}
			}

			if gamePossible == false {
				break
			}
		}
		if gamePossible == true {
			possibleGames = append(possibleGames, gameID)
			gameIDSplits := strings.Split(gameID, " ")
			ID, _ := strconv.Atoi(gameIDSplits[1])
			sum += ID
		}
	}

	fmt.Println(possibleGames)
	fmt.Println(sum)
}
