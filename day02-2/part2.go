package main

import (
	"fmt"
	"os"
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
	var safeCount int = 0
	for _, elem := range inputLines {
		lineValue := strings.Split(elem, " ")
		safe := isSafe(lineValue, false)

		if !safe {
			for i := 0; i < len(lineValue); i++ {
				newSlice := make([]string, len(lineValue))
				copy(newSlice, lineValue)
				remove(newSlice, i)
				newSlice2 := make([]string, len(newSlice)-1)
				copy(newSlice2, newSlice)
				safe = isSafe(newSlice2, true)
				if safe {
					break
				}
			}
		}

		if safe {
			safeCount++
		}

	}

	fmt.Print(safeCount)
}

func getDirection(lastDigit int, newDigit int) string {
	if lastDigit > newDigit {
		return "-"
	}
	if lastDigit < newDigit {
		return "+"
	}
	return "o"
}

func abs(a int, b int) int {
	var c int = a - b
	if c < 0 {
		c *= -1
	}
	return c
}

func isSafe(lineValue []string, forceNoIgnore bool) bool {
	var lastDigit int = -1
	var direction string = "x"
	var unsafeCount int = 0
	if forceNoIgnore {
		unsafeCount = 1
	}
	for _, val := range lineValue {
		var ignore bool = false

		intVal, _ := strconv.Atoi(val)

		if lastDigit != -1 {
			newDirection := getDirection(lastDigit, intVal)
			diff := abs(lastDigit, intVal)
			if newDirection == "o" || diff < 1 || diff > 3 {
				ignore = true
			} else {
				if direction != "x" {
					if newDirection != direction || diff < 1 || diff > 3 {
						ignore = true
					}
				}
				if !ignore {
					direction = newDirection
				}
			}
		}
		if !ignore {
			lastDigit = intVal
		}
		if ignore {
			unsafeCount++
		}
	}
	return unsafeCount <= 1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
