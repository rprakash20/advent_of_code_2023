package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Processing int

const (
	smallInputFile = "../input1.txt"
	inputFile      = "../input.txt"
)

const (
	PROCESSING_UNSPECIFIED Processing = iota
	SEED_TO_SOIL
	SOIL_TO_FERT
	FERT_TO_WATER
	WATER_TO_LIGHT
	LIGHT_TO_TEMP
	TEMP_TO_HUMID
	HUMID_TO_LOC
)

func (p Processing) String() string {
	return [...]string{"processing_unspecified",
		"seed_to_soil", "soil_to_fert", "fert_to_water", "water_to_light", "light_to_temp", "temp_to_humid", "humid_to_loc",
	}[p]
}

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

func getResNum(dataMap map[string][]int, num int) (resNum int) {
	for i := 0; i < len(dataMap["src"]); i++ {
		if num >= dataMap["src"][i] &&
			num < dataMap["src"][i]+dataMap["ranges"][i] {
			resNum = dataMap["dest"][i] + num - dataMap["src"][i]
			break
		}
	}
	if resNum == 0 {
		resNum = num
	}
	return resNum
}

func main() {
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
	start := false
	processing := PROCESSING_UNSPECIFIED
	almanac := make(map[string]map[string][]int)
	var lowestLocNum int
	var seedsList []string

	for scanner.Scan() {
		lineStr := scanner.Text()

		if strings.HasPrefix(lineStr, "seeds:") {
			seedsSplit := strings.Split(lineStr, ": ")
			seedsStr := seedsSplit[1]
			seedsList = strings.Fields(seedsStr)
		}

		if strings.HasPrefix(lineStr, "seed-to-soil map:") {
			start = true
			processing = SEED_TO_SOIL
		} else if strings.HasPrefix(lineStr, "soil-to-fertilizer map:") {
			start = true
			processing = SOIL_TO_FERT
		} else if strings.HasPrefix(lineStr, "fertilizer-to-water map:") {
			processing = FERT_TO_WATER
			start = true
		} else if strings.HasPrefix(lineStr, "water-to-light map:") {
			processing = WATER_TO_LIGHT
			start = true
		} else if strings.HasPrefix(lineStr, "light-to-temperature map:") {
			processing = LIGHT_TO_TEMP
			start = true
		} else if strings.HasPrefix(lineStr, "temperature-to-humidity map:") {
			processing = TEMP_TO_HUMID
			start = true
		} else if strings.HasPrefix(lineStr, "humidity-to-location map:") {
			processing = HUMID_TO_LOC
			start = true
		} else if lineStr == "" && start == true {
			start = false
		} else if processing != PROCESSING_UNSPECIFIED {
			lineStrSplits := strings.Fields(lineStr)
			dest, _ := strconv.Atoi(lineStrSplits[0])
			source, _ := strconv.Atoi(lineStrSplits[1])
			ranges, _ := strconv.Atoi(lineStrSplits[2])

			if _, ok := almanac[processing.String()]; !ok {
				almanac[processing.String()] = make(map[string][]int)
			}
			almanac[processing.String()]["src"] = append(almanac[processing.String()]["src"], source)
			almanac[processing.String()]["dest"] = append(almanac[processing.String()]["dest"], dest)
			almanac[processing.String()]["ranges"] = append(almanac[processing.String()]["ranges"], ranges)
		}
	}

	//fmt.Println(almanac)

	for _, seed := range seedsList {
		seedNum, _ := strconv.Atoi(seed)
		var (
			soilNum  int
			fertNum  int
			waterNum int
			lightNum int
			tempNum  int
			humidNum int
			locNum   int
		)

		soilNum = getResNum(almanac[SEED_TO_SOIL.String()], seedNum)

		fertNum = getResNum(almanac[SOIL_TO_FERT.String()], soilNum)

		waterNum = getResNum(almanac[FERT_TO_WATER.String()], fertNum)

		lightNum = getResNum(almanac[WATER_TO_LIGHT.String()], waterNum)

		tempNum = getResNum(almanac[LIGHT_TO_TEMP.String()], lightNum)

		humidNum = getResNum(almanac[TEMP_TO_HUMID.String()], tempNum)

		locNum = getResNum(almanac[HUMID_TO_LOC.String()], humidNum)

		fmt.Println(seedNum, soilNum, fertNum, waterNum, lightNum, tempNum, humidNum, locNum)

		if lowestLocNum == 0 || locNum < lowestLocNum {
			lowestLocNum = locNum
		}
	}

	fmt.Println("Answer: ", lowestLocNum)
}
