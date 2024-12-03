package main

import (
	"advent_of_code_2024/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lists, err := utils.MultipleLines("./input.txt")
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	var leftSide []float64
	var rightSide []float64

	for _, line := range lists {
		temp := strings.Fields(line)
		left, _ := strconv.ParseFloat(temp[0], 64)
		right, _ := strconv.ParseFloat(temp[1], 64)
		leftSide = append(leftSide, left)
		rightSide = append(rightSide, right)
	}

	slices.Sort(leftSide)
	slices.Sort(rightSide)

	var distance float64
	var similarityScore float64
	for index, leftValue := range leftSide {
		// Part 1
		distance += math.Abs(leftValue - rightSide[index])

		// Part 2
		count := utils.Count(rightSide, func(x float64) bool {
			return x == leftValue
		})
		similarityScore += leftValue * float64(count)
	}

	fmt.Printf("Distance: %.0f\n", distance)
	fmt.Printf("Similarity score: %.0f\n", similarityScore)
}
