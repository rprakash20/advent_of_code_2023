package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	input_file = "input.txt"
)

var (
	digitMap map[string]int
)

func init() {
	digitMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findAllMatches(str string) []int {
	var matches []int
	lenStr := len(str)
	for i := 0; i < lenStr; i++ {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {

			if i+3 <= lenStr {
				if str[i:i+3] == "one" || str[i:i+3] == "two" || str[i:i+3] == "six" {
					matches = append(matches, digitMap[str[i:i+3]])
				}
			}
			if i+4 <= lenStr {
				if str[i:i+4] == "four" || str[i:i+4] == "five" || str[i:i+4] == "nine" {
					matches = append(matches, digitMap[str[i:i+4]])
				}
			}
			if i+5 <= lenStr {
				if str[i:i+5] == "three" || str[i:i+5] == "seven" || str[i:i+5] == "eight" {
					matches = append(matches, digitMap[str[i:i+5]])
				}
			}
		} else {
			matches = append(matches, num)
		}
	}
	return matches
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
		numList := findAllMatches(lineStr)
		numLine := numList[0]*10 + numList[len(numList)-1]
		sum += numLine
	}

	fmt.Println(sum)
}
