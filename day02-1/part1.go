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
		var lastDigit int = -1
		var direction string = "x"
		var safe bool = true
		for _, val := range lineValue {
			intVal, _ := strconv.Atoi(val)
			if lastDigit != -1 {
				newDirection := getDirection(lastDigit, intVal)
				diff := abs(lastDigit, intVal)
				if newDirection == "o" || diff < 1 || diff > 3 {
					safe = false
				} else {
					if direction != "x" {
						if newDirection != direction || diff < 1 || diff > 3 {
							safe = false
						}
					}
					if safe {
						direction = newDirection
					}
				}
			}
			if safe {
				lastDigit = intVal
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
