package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/crazyvinvin/advent-of-code/2024/shared"
)

func main() {
	startTime := time.Now()
	input := shared.ReadInput("../day4")

	elapsed := time.Since(startTime)
	fmt.Printf("Reading input took: %s\n", elapsed)

	startTimeTranslateInput := time.Now()

	lines := createLinesFromInput(input)

	elapsed = time.Since(startTimeTranslateInput)
	fmt.Printf("createLinesFromInput took: %s\n", elapsed)

	solutionStartTime := time.Now()

	var options [][]string
	options = append(options, getHorizontalOptions(lines)...)
	options = append(options, getVerticalOptions(lines)...)
	options = append(options, getDiagonalOptions(lines)...)
	options = append(options, getRevertedSlices(options)...)

	fmt.Println("Total Options: ", len(options))
	fmt.Println("Total XMAS: ", countXMAS(options))

	elapsed = time.Since(solutionStartTime)
	fmt.Printf("Solution took: %s\n", elapsed)

	elapsed = time.Since(startTime)
	fmt.Printf("Total script took: %s\n", elapsed)
}

func countXMAS(stringSlices [][]string) int {
	numberOfXMAS := 0
	for _, slice := range stringSlices {
		word := strings.Join(slice, "")
		if word == "XMAS" {
			numberOfXMAS++
		}
	}
	return numberOfXMAS
}

func getRevertedSlices(_slices [][]string) [][]string {
	result := make([][]string, len(_slices))

	for i, slice := range _slices {
		result[i] = make([]string, len(slice))
		copy(result[i], slice)
		slices.Reverse(result[i])
	}

	fmt.Println("Reverted options: ", len(result))
	return result
}

func getHorizontalOptions(lines [][]string) [][]string {
	var options [][]string
	for _, horizontalLine := range lines {
		for i := 0; i < len(horizontalLine)-3; i++ {
			options = append(options, []string{
				horizontalLine[i],
				horizontalLine[i+1],
				horizontalLine[i+2],
				horizontalLine[i+3],
			})
		}
	}
	fmt.Println("Found horizontal options: ", len(options))
	return options
}

func getVerticalOptions(lines [][]string) [][]string {
	var options [][]string
	for horizontalIndex := 0; horizontalIndex < len(lines[0]); horizontalIndex++ {
		for verticalIndex := 0; verticalIndex < len(lines)-3; verticalIndex++ {
			options = append(options, []string{
				lines[verticalIndex][horizontalIndex],
				lines[verticalIndex+1][horizontalIndex],
				lines[verticalIndex+2][horizontalIndex],
				lines[verticalIndex+3][horizontalIndex],
			})
		}
	}
	fmt.Println("Found vertical options: ", len(options))
	return options
}

func getDiagonalOptions(lines [][]string) [][]string {
	var options [][]string
	for horizontalIndex := 0; horizontalIndex < len(lines[0])-3; horizontalIndex++ {
		for verticalIndex := 0; verticalIndex < len(lines)-3; verticalIndex++ {
			options = append(options, []string{
				lines[verticalIndex][horizontalIndex],
				lines[verticalIndex+1][horizontalIndex+1],
				lines[verticalIndex+2][horizontalIndex+2],
				lines[verticalIndex+3][horizontalIndex+3],
			})
		}
	}

	for horizontalIndex := 3; horizontalIndex < len(lines[0]); horizontalIndex++ {
		for verticalIndex := 0; verticalIndex < len(lines)-3; verticalIndex++ {
			options = append(options, []string{
				lines[verticalIndex][horizontalIndex],
				lines[verticalIndex+1][horizontalIndex-1],
				lines[verticalIndex+2][horizontalIndex-2],
				lines[verticalIndex+3][horizontalIndex-3],
			})
		}
	}

	fmt.Println("Found diagonal options: ", len(options))
	return options
}

func createLinesFromInput(input string) [][]string {
	linesAsStrings := strings.Split(input, "\r\n")

	var lines [][]string

	for _, str := range linesAsStrings {
		line := []string{}
		for _, char := range str {
			line = append(line, string(char))
		}
		lines = append(lines, line)
	}

	return lines
}
