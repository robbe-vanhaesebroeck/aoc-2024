package main

import (
	"aoc-2024/common"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readLines(fileName string) string {
	data, err := os.ReadFile(fileName)
	common.CheckError(err)

	return string(data)
}

var mulRegex = regexp.MustCompile(`mul\((?P<firstNum>\d{1,3}),(?P<secondNum>\d{1,3})\)`)
var mulRegexPart2 = regexp.MustCompile(`mul\((?<firstNum>\d{1,3}),(?P<secondNum>\d{1,3})\)|do\(\)|don\'t\(\)`)

func part1(fileName string) int {
	instructions := readLines(fileName)

	
	mulInstructions := mulRegex.FindAllStringSubmatch(instructions, -1)

	firstNumIdx := mulRegex.SubexpIndex("firstNum")
	secondNumIdx := mulRegex.SubexpIndex("secondNum")
	
	sum := 0

	for _, group := range mulInstructions {
		firstNum, err := strconv.Atoi(group[firstNumIdx])
		common.CheckError(err)
		secondNum, err := strconv.Atoi(group[secondNumIdx])
		common.CheckError(err)

		sum += firstNum * secondNum
	}

	return sum
}

func part2(fileName string) int {
	instructions := readLines(fileName)

	allInstructions := mulRegexPart2.FindAllStringSubmatch(instructions, -1)

	firstNumIdx := mulRegexPart2.SubexpIndex("firstNum")
	secondNumIdx := mulRegexPart2.SubexpIndex("secondNum")

	enabled := true
	sum := 0

	for _, group := range allInstructions {
		if (group[0] == "do()") {
			enabled = true
		} else if (group[0] == "don't()") {
			enabled = false
		} else if (enabled) {
			firstNum, err := strconv.Atoi(group[firstNumIdx])
			common.CheckError(err)
			secondNum, err := strconv.Atoi(group[secondNumIdx])
			common.CheckError(err)

			sum += firstNum * secondNum
		}
	}
	return sum
}

func main() {
	fileName := "day-3/input.txt"

	fmt.Println("Part 1")
	result1 := part1(fileName)
	fmt.Println(result1)

	fmt.Println("Part 2")
	result2 := part2(fileName)
	fmt.Println(result2)

}