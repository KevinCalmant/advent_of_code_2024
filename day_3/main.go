package main

import (
	"advent_of_code_2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func extractAndSumMulValues(muls []string) int {
	var total int
	for _, sumMatch := range muls {
		numberRe := regexp.MustCompile(`\d+`)
		numberMatches := numberRe.FindAllString(sumMatch, -1)

		firstValue, _ := strconv.Atoi(numberMatches[0])
		secondValue, _ := strconv.Atoi(numberMatches[1])

		total += firstValue * secondValue
	}
	return total
}

func main() {
	lines, _ := utils.MultipleLines("input.txt")

	concatInput := strings.Join(lines, "")

	sumRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	sumMatches := sumRe.FindAllString(concatInput, -1)
	total := extractAndSumMulValues(sumMatches)

	fmt.Printf("Total: %d\n", total)

	dontDoRe := regexp.MustCompile(`don't\(\).*?do\(\)`)
	dontDoReplaced := dontDoRe.ReplaceAllString(concatInput, "")
	dontSumMatches := sumRe.FindAllString(dontDoReplaced, -1)
	totalFiltered := extractAndSumMulValues(dontSumMatches)

	fmt.Printf("Total filtered: %d\n", totalFiltered)
}
