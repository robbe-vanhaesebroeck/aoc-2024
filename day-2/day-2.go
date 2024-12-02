package main

import (
	"aoc-2024/common"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readLines(fileName string) [][]int{
	data, err := os.ReadFile(fileName)
	common.CheckError(err)

	parsedData := string(data)
	reportLines := strings.Split(parsedData, "\n")

	reports := make([][]int, len(reportLines))

	for idx, line := range reportLines {
		values := strings.Split(line, " ")

		report := make([]int, len(values))
		for valIdx, value := range values {
			report[valIdx], err = strconv.Atoi(value)
			common.CheckError(err)
		}

		reports[idx] = report
	}

	return reports
}

func part1(fileName string) int {
	reports := readLines(fileName)
	safeReports := 0

	for idx, report := range reports {
		diff := make([]int, len(report) - 1)
		
		// Calculate the difference between sequential values in the report
		for idx, value := range report[:len(report) - 1] {
			nextValue := report[idx + 1]
			diff[idx] = nextValue - value
		}

		sign := math.Signbit(float64(diff[0]))
		safe := true

		for _, value := range diff {
			// If the sign is different than the previous value, the report is not safe
			if sign != math.Signbit(float64(value)) {
				safe = false
				break
			}

			// If the absolute difference between both values is greater than 3 or smaller than 1, the report is not safe
			if math.Abs(float64(value)) > 3 || math.Abs(float64(value)) < 1 {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	return safeReports
}

func part2(fileName string) int {
	return 0
}

func main() {
	fileName := "day-2/input.txt"

	fmt.Println("Part 1")
	result1 := part1(fileName)
	fmt.Println("Result:", result1)

	fmt.Println("Part 2")
	result2 := part2(fileName)
	fmt.Println("Result:", result2)
}