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
	useActualInput  bool
	adjacentIndexes [][]int
	engineSchematic [][]string
)

func init() {
	// change this to false to run this program
	// with smaller input
	useActualInput = true
	engineSchematic = make([][]string, 0)
	adjacentIndexes = [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
		{1, -1},
		{-1, -1},
		{-1, 1},
		{1, 1},
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isAGear(i int, j int, engineSchematic [][]string) (isGear bool, gearCoordinates string) {
	for _, coOrdinates := range adjacentIndexes {
		x := i + coOrdinates[0]
		y := j + coOrdinates[1]
		if x >= 0 && y >= 0 && x < len(engineSchematic) && y < len(engineSchematic[0]) {
			if engineSchematic[x][y] == "*" {
				isGear = true
				gearCoordinates = strconv.Itoa(x) + strconv.Itoa(y)
				return isGear, gearCoordinates
			}
		}
	}

	return false, ""
}

func updateGearMap(coordinates string, num int, gearMap map[string][]int) {
	_, ok := gearMap[coordinates]
	if !ok {
		gearMap[coordinates] = make([]int, 0)
	}
	gearMap[coordinates] = append(gearMap[coordinates], num)
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
		lineStr := scanner.Text()
		lineStrSplits := strings.Split(lineStr, "")
		var row []string
		row = append(row, lineStrSplits...)
		engineSchematic = append(engineSchematic, row)
	}

	var gearMap map[string][]int
	gearMap = make(map[string][]int)
	for i := 0; i < len(engineSchematic); i++ {
		row := engineSchematic[i]
		var currentNum int
		isGear := false
		var gearCoordinates string
		for j := 0; j < len(row); j++ {
			digit, err := strconv.Atoi(row[j])
			if err == nil {
				if isGear == false {
					isGear, gearCoordinates = isAGear(i, j, engineSchematic)
				}
				currentNum = currentNum*10 + digit

				// Do not miss the number existing at the end of the row
				if isGear == true && j == len(row)-1 {
					updateGearMap(gearCoordinates, currentNum, gearMap)
				}
			} else {
				if currentNum != 0 && isGear {
					updateGearMap(gearCoordinates, currentNum, gearMap)
				}
				currentNum = 0
				isGear = false
			}
		}
	}
	//fmt.Println("gearMap:", gearMap)

	for _, v := range gearMap {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	fmt.Println("Answer:", sum)
}
