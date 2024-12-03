package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Size of the maximum construct. In this case, don't()
const chunkSize int = 7

// From exercise, mul(X,Y), where X and Y are each 1-3 digit numbers.
const parametersAndClosingParenthesisMaxLength int = 8

const mulFn string = "mul("

func main() {
	startTime := time.Now()

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	input := string(file)

	enabled := true
	lastAppliedBound := -1
	result := 0

	for i := chunkSize; i <= len(input); i++ {
		chunk := input[i-chunkSize : i]

		if enabled && strings.Contains(chunk, "don't()") {
			enabled = false
		} else if !enabled && strings.Contains(chunk, "do()") {
			enabled = true
		}

		if !enabled {
			continue
		}

		mulIndex := strings.Index(chunk, mulFn)

		if mulIndex == -1 {
			continue
		}

		nextCharsLeftBound := i - chunkSize + mulIndex + len(mulFn)
		nextCharsRightBound := nextCharsLeftBound + parametersAndClosingParenthesisMaxLength

		if nextCharsRightBound > len(input) {
			nextCharsRightBound = len(input)
		}

		nextChars := input[nextCharsLeftBound:nextCharsRightBound]

		closingParenthesisIndex := strings.Index(nextChars, ")")
		if closingParenthesisIndex == -1 {
			continue
		}

		if lastAppliedBound == nextCharsLeftBound {
			continue
		}
		lastAppliedBound = nextCharsLeftBound

		paramsStr := nextChars[0:closingParenthesisIndex]
		params := strings.Split(paramsStr, ",")

		l, err := strconv.Atoi(params[0])
		if err != nil {
			continue
		}

		r, err := strconv.Atoi(params[1])
		if err != nil {
			continue
		}

		if l < 0 || l > 999 || r < 0 || r > 999 {
			continue
		}

		result += l * r
	}

	log.Printf("The result is %d\n. Took: %d microseconds", result, time.Since(startTime).Microseconds())
}
