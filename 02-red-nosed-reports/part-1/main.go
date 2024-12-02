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

	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)

		secondItem, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Error converting item %s to int: %v", tokens[1], err)
		}

		firstItem, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatalf("Error converting item %s to int: %v", tokens[0], err)
		}

		increasing := secondItem > firstItem

		i := 1

		for ; i < len(tokens); i++ {
			secondItem, err := strconv.Atoi(tokens[i])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", tokens[i], err)
			}

			firstItem, err := strconv.Atoi(tokens[i-1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", tokens[i-1], err)
			}

			if secondItem == firstItem {
				break
			}

			if secondItem > firstItem && (!increasing || secondItem-firstItem > 3) {
				break
			}

			if secondItem < firstItem && (increasing || firstItem-secondItem > 3) {
				break
			}
		}

		if i == len(tokens) {
			safeCount++
		}
	}

	log.Printf("The result is %d\n. Took: %d microseconds", safeCount, time.Since(startTime).Microseconds())
}
