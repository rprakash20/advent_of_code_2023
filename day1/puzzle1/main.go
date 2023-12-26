package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	input_file = "input.txt"
)

var (
	digitRegex *regexp.Regexp
)

func init() {
	digitRegex = regexp.MustCompile("[0-9]")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sum := 0

	f, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineStr := scanner.Text()
		numList := digitRegex.FindAllString(lineStr, -1)
		firstNum, _ := strconv.Atoi(numList[0])
		lastNum, _ := strconv.Atoi(numList[len(numList)-1])
		numLine := firstNum*10 + lastNum
		sum += numLine
	}

	fmt.Println(sum)
}
