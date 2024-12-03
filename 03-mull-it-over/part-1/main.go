package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const mulStart string = "mul("

// From exercise, mul(X,Y), where X and Y are each 1-3 digit numbers.
const parametersAndClosingParenthesisMaxLength int = 8

func main() {
	startTime := time.Now()

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	input := string(file)

	result := 0

	for i := len(mulStart); i <= len(input); i++ {
		fn := input[i-len(mulStart) : i]

		if fn != mulStart {
			continue
		}

		rightBound := i + parametersAndClosingParenthesisMaxLength + 2
		if rightBound > len(input) {
			rightBound = len(input)
		}

		nextChars := input[i:rightBound]

		closingParenthesisIndex := strings.Index(nextChars, ")")
		if closingParenthesisIndex == -1 {
			continue
		}

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
