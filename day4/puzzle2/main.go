package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cardDetails struct {
	drawNums        []string
	winningNums     []string
	cardNum         int
	scratchCardsNum int
}

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
	lotteryCardMap := make(map[string]*cardDetails)
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
		card := strings.Fields(lineSplits[0])
		currentCardNum, _ := strconv.Atoi(card[1])
		lotteryNums := strings.Split(lineSplits[1], " | ")
		winningNums := strings.Fields(lotteryNums[0])
		drawNums := strings.Fields(lotteryNums[1])
		lotteryCardMap[fmt.Sprint(currentCardNum)] = &cardDetails{
			winningNums:     winningNums,
			drawNums:        drawNums,
			cardNum:         currentCardNum,
			scratchCardsNum: 1,
		}
	}

	totalCardsCount := len(lotteryCardMap)
	for num := 1; num <= totalCardsCount; num++ {
		details := lotteryCardMap[fmt.Sprint(num)]
		drawNumsMap := make(map[string]int)
		count := 0
		for _, drawNum := range details.drawNums {
			drawNumsMap[drawNum] = 1
		}

		for _, winningNum := range details.winningNums {
			if _, ok := drawNumsMap[winningNum]; ok {
				count++
			}
		}

		for j := 0; j < details.scratchCardsNum; j++ {
			for i := 1; i <= count && i <= totalCardsCount; i++ {
				cardNum := details.cardNum + i
				lotteryCardMap[fmt.Sprint(cardNum)].scratchCardsNum++
			}
		}
	}

	for num := 1; num <= totalCardsCount; num++ {
		sum += lotteryCardMap[fmt.Sprint(num)].scratchCardsNum
	}
	fmt.Println(sum)
}
