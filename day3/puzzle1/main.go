package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	digitDotRegex   *regexp.Regexp
)

func init() {
	// change this to false to run this program
	// with smaller input
	useActualInput = true
	engineSchematic = make([][]string, 0)
	digitDotRegex = regexp.MustCompile(`[0-9]`)
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

func hasAdjacentSymbols(i int, j int, engineSchematic [][]string) bool {
	hasSymbols := false
	for _, coOrdinates := range adjacentIndexes {
		x := i + coOrdinates[0]
		y := j + coOrdinates[1]
		if x >= 0 && y >= 0 && x < len(engineSchematic) && y < len(engineSchematic[0]) {
			matched := digitDotRegex.Match([]byte(engineSchematic[x][y]))
			if matched == false && engineSchematic[x][y] != "." {
				hasSymbols = true
				break
			}
		}
	}

	if hasSymbols {
		return true
	}
	return false
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

	for i := 0; i < len(engineSchematic); i++ {
		row := engineSchematic[i]
		var numsInRow []int
		var currentNum int
		hasSymbols := false
		for j := 0; j < len(row); j++ {
			digit, err := strconv.Atoi(row[j])
			if err == nil {
				if hasSymbols == false {
					hasSymbols = hasAdjacentSymbols(i, j, engineSchematic)
				}
				currentNum = currentNum*10 + digit
				// Do not miss the number existing at the end of the row
				if hasSymbols == true && j == len(row)-1 {
					numsInRow = append(numsInRow, currentNum)
				}
			} else {
				if currentNum != 0 && hasSymbols {
					numsInRow = append(numsInRow, currentNum)
				}
				currentNum = 0
				hasSymbols = false
			}
		}

		for _, num := range numsInRow {
			sum += num
		}
	}
	fmt.Println("Answer:", sum)
}
