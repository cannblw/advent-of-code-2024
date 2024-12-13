package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/mat"
)

const readModeCount int = 3

const (
	ButtonA int = 0
	ButtonB int = 1
	Prize   int = 2
)

const tokensA = 3
const tokensB = 1

func isInteger(f float64) bool {
	return math.Abs(f-math.Round(f)) < 1e-9
}

func main() {
	startTime := time.Now()

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	r := regexp.MustCompile(`(\d+).*?(\d+)`)
	if err != nil {
		log.Fatalf("Error compiling regexp: %v", err)
	}

	result := 0

	scanner := bufio.NewScanner(file)
	readMode := -1

	var a1, a2, b1, b2, p1, p2 int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		readMode = (readMode + 1) % readModeCount

		matches := r.FindStringSubmatch(line)

		switch readMode {
		case ButtonA:
			a1, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[1], err)
			}
			a2, err = strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[2], err)
			}

		case ButtonB:
			b1, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[1], err)
			}

			b2, err = strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[2], err)
			}

		case Prize:
			p1, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[1], err)
			}
			p2, err = strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[2], err)
			}

			A := mat.NewDense(2, 2, []float64{float64(a1), float64(b1), float64(a2), float64(b2)})
			b := mat.NewVecDense(2, []float64{float64(p1), float64(p2)})

			var x mat.VecDense
			err := x.SolveVec(A, b)
			if err != nil {
				continue
			}

			solutionA := x.AtVec(0)
			solutionB := x.AtVec(1)

			if isInteger(solutionA) && isInteger(solutionB) {
				a := int(math.Round(solutionA))
				b := int(math.Round(solutionB))

				if a >= 0 && b >= 0 {
					result += a*tokensA + b*tokensB
				}
			}
		}
	}

	log.Printf("The result is %d. Took: %d microseconds", result, time.Since(startTime).Microseconds())
}
