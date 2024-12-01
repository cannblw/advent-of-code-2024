package main

import (
	"bufio"
	"cmp"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	defer file.Close()

	numbersA := make([]int, 0)
	numbersB := make([]int, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		a, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatalf("Error converting item %s to int: %v", tokens[0], err)
		}

		b, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Error converting item %s to int: %v", tokens[0], err)
		}

		numbersA = append(numbersA, a)
		numbersB = append(numbersB, b)
	}

	slices.SortStableFunc(numbersA, cmp.Compare[int])
	slices.SortStableFunc(numbersB, cmp.Compare[int])

	distance := 0

	for i := 0; i < len(numbersA); i++ {
		distance += int(math.Abs(float64(numbersB[i] - numbersA[i])))
	}

	log.Printf("The result is %d\n. Took: %d microseconds", distance, time.Since(startTime).Microseconds())
}
