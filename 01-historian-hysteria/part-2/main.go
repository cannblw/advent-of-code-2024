package main

import (
	"bufio"
	"log"
	"os"
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
	numbersAndCountB := make(map[int]int)

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
		numbersAndCountB[b] += 1
	}

	similarity := 0
	for _, item := range numbersA {
		similarity += item * numbersAndCountB[item]
	}

	log.Printf("The result is %d\n. Took: %d microseconds", similarity, time.Since(startTime).Microseconds())
}
