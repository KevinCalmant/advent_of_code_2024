package main

import (
	"advent_of_code_2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getReportStatus(reports []string) string {
	const (
		safe   = "SAFE"
		unsafe = "UNSAFE"
	)

	var tempLevelDiff int

	for index := 0; index < len(reports)-1; index++ {
		currentLevel, err1 := strconv.Atoi(reports[index])
		nextLevel, err2 := strconv.Atoi(reports[index+1])

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input in reports")
			return unsafe
		}

		diff := nextLevel - currentLevel
		absoluteDiff := math.Abs(float64(diff))

		if absoluteDiff > 3 || absoluteDiff < 1 {
			return unsafe
		}
		if (tempLevelDiff > 0 && diff < 0) || (tempLevelDiff < 0 && diff > 0) {
			return unsafe
		}

		tempLevelDiff = diff
	}

	return safe
}

func main() {
	lists, err := utils.MultipleLines("./input.txt")
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	var safe int
	var safeDampenered int
	for _, line := range lists {
		reports := strings.Fields(line)
		if getReportStatus(reports) == "SAFE" {
			safe++
		} else {
			for i := 0; i < len(reports); i++ {
				dampeneredReports := utils.DeleteElement(reports, i)
				if getReportStatus(dampeneredReports) == "SAFE" {
					safeDampenered++
					break
				}
			}
		}
	}
	fmt.Printf("Safe reports: %d\n", safe)
	fmt.Printf("Safe dampanered reports: %d", safeDampenered+safe)
}
