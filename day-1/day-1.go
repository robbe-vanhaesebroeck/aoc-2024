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
	firstNumbers, secondNumbers := readLines(fileName)

	sort.Ints(firstNumbers)
	sort.Ints(secondNumbers)

	sum := 0

	for idx, firstNumber := range firstNumbers {
		secondNumber := secondNumbers[idx]

		sum += int(math.Abs(float64(secondNumber - firstNumber)))
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
}
