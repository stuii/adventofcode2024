package main

import (
	"ioutil"
)

var input string

func main() {
	loadInput()

}

func loadInput() {
	body, err := ioutil.ReadFile("./input.txt")
	if err == nil {
		input = body
	}
}
