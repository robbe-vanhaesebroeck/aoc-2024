package main

import (
	"aoc-2024/common"
	"fmt"
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

func getDiffMaps(report []int) (map[int]int, map[int]int) {
	posDiffMap := make(map[int]int)
	negDiffMap := make(map[int]int)

	for idx, value := range report[:len(report) - 1] {
		diff := report[idx+1] - value 

		if diff >= 0 {
			posDiffMap[diff]++
		}

		if diff < 0 {
			negDiffMap[diff]++
		}
	}

	return posDiffMap, negDiffMap
}

func isValid(report []int) bool {
	posDiffMap, negDiffMap := getDiffMaps(report)

	// Only one of these should be non empty
	if len(posDiffMap) > 0 && len(negDiffMap) > 0 {
		return false
	}

	for k := range posDiffMap {
		if k == 0 || k > 3 {
			return false
		}
	}

	for k := range negDiffMap {
		if k == 0 || k < -3 {
			return false
		}
	}

	return true
}


func part1(fileName string) int {
	reports := readLines(fileName)
	safeReports := 0

	for _, report := range reports {
		if isValid(report) {
			safeReports++
		} 
	}

	return safeReports
}

func part2(fileName string) int {
	reports := readLines(fileName)
	safeReports := 0

	for _, report := range reports {
		validSubReports := 0

		if isValid(report) {
			safeReports++
			continue
		}

		for idx := 0; idx < len(report); idx++ {
			copiedReport := make([]int, len(report))
			copy(copiedReport, report)
			erasedReport := append(copiedReport[:idx], copiedReport[idx+1:]...)
			if isValid(erasedReport) {
				validSubReports++
			}
		}

		if validSubReports > 0 {
			safeReports++
		}
	}

	return safeReports
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