package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		lineStr := scanner.Text()
		lineSplits := strings.Split(lineStr, ": ")
		lotteryNums := strings.Split(lineSplits[1], " | ")
		winningNums := strings.Fields(lotteryNums[0])
		drawNums := strings.Fields(lotteryNums[1])
		points := 0
		drawNumsMap := make(map[string]int)
		for _, drawNum := range drawNums {
			drawNumsMap[drawNum] = 1
		}

		for _, winningNum := range winningNums {
			if _, ok := drawNumsMap[winningNum]; ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}

	fmt.Println(sum)
}
