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
)

func init() {
	useActualInput = true
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
	for scanner.Scan() {
		minCubeContent := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		lineStr := scanner.Text()
		lineSplits := strings.Split(lineStr, ": ")
		_, gameCubes := lineSplits[0], lineSplits[1]
		gameCubesSplits := strings.Split(gameCubes, "; ")
		for _, cubes := range gameCubesSplits {
			cubeSplits := strings.Split(cubes, ", ")

			for _, cubeConf := range cubeSplits {
				cubeConfSplits := strings.Split(cubeConf, " ")
				cubeCount, _ := strconv.Atoi(cubeConfSplits[0])
				cubeColor := cubeConfSplits[1]
				if minCubeContent[cubeColor] < cubeCount {
					minCubeContent[cubeColor] = cubeCount
				}
			}
		}
		sum += minCubeContent["red"] * minCubeContent["green"] * minCubeContent["blue"]
	}

	fmt.Println(sum)
}
