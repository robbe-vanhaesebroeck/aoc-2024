package main

import (
	"aoc-2024/common"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(fileName string) ([]int, []int) {
	data, err := os.ReadFile(fileName)

	common.CheckError(err)

	parsedData := string(data)
	lines := strings.Split(parsedData, "\n")

	numbers1 := make([]int, len(lines))
	numbers2 := make([]int, len(lines))

	for idx, line := range lines {
		if line == "" {
			continue
		}

		lineValues := strings.Split(line, "   ")

		firstNumber, err := strconv.Atoi(lineValues[0])
		common.CheckError(err)
		numbers1[idx] = firstNumber

		secondNumber, err := strconv.Atoi(lineValues[1])
		common.CheckError(err)
		numbers2[idx] = secondNumber
	}

	return numbers1, numbers2
}

func part1(fileName string) int {
	firstLocationIds, secondLocationIds := readLines(fileName)

	sort.Ints(firstLocationIds)
	sort.Ints(secondLocationIds)

	sum := 0

	for idx, firstNumber := range firstLocationIds {
		secondNumber := secondLocationIds[idx]

		sum += int(math.Abs(float64(secondNumber - firstNumber)))
	}

	return sum
}

func countOccurrences(occurrences []int) map[int]int {
	count := make(map[int]int)

	for _, occurrence := range occurrences {
		count[occurrence] += 1
	}

	return count
}

func part2(fileName string) int {
	locationIds, occurrences := readLines(fileName)

	occurrenceMap := countOccurrences(occurrences)

	sum:= 0

	for _, locationId := range locationIds {
		sum += locationId * occurrenceMap[locationId]	
	}

	return sum
}

func main() {
	fileName := "day-1/input.txt"

	fmt.Println("Part 1")
	result1 := part1(fileName)
	fmt.Println("Result:", result1)

	// Part 1

	fmt.Println("Part 2")
	result2 := part2(fileName)
	fmt.Println("Result:", result2)
}
