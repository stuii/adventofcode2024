package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var input string
var inputLines []string

func main() {
	start := time.Now()

	loadInput()
	parseInput()
	solve()

	elapsed := time.Since(start)
	fmt.Println()
	fmt.Printf("Execution Time: %s", elapsed)
}

func loadInput() {
	body, err := os.ReadFile("input.txt")
	if err == nil {
		input = string(body)
	}
}

func parseInput() {
	lines := strings.Split(input, "\r\n")
	inputLines = append(inputLines, lines...)
}

func solve() {
	var similarity int = 0
	var leftValues []int
	var rightValues []int
	for _, elem := range inputLines {
		lineValue := strings.Split(elem, "   ")
		leftInt, _ := strconv.Atoi(lineValue[0])
		rightInt, _ := strconv.Atoi(lineValue[1])

		leftValues = append(leftValues, leftInt)
		rightValues = append(rightValues, rightInt)
	}

	sort.Slice(leftValues, func(a, b int) bool {
		return leftValues[a] < leftValues[b]
	})

	sort.Slice(rightValues, func(a, b int) bool {
		return rightValues[a] < rightValues[b]
	})

	for i := 0; i < len(leftValues); i++ {
		similarity += countOccurrences(rightValues, leftValues[i]) * leftValues[i]

	}
	fmt.Print(similarity)
}

func countOccurrences(haystack []int, needle int) int {
	count := 0
	for _, value := range haystack {
		if value == needle {
			count++
		}
	}
	return count
}
