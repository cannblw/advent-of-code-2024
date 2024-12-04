package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	lines := strings.Split(string(file), "\n")

	result := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] != 'A' {
				continue
			}

			topLeft := lines[i-1][j-1]
			topRight := lines[i-1][j+1]
			bottomLeft := lines[i+1][j-1]
			bottomRight := lines[i+1][j+1]

			strokesMatch := 0
			if (topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M') {
				strokesMatch++
			}
			if (topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M') {
				strokesMatch++
			}

			// An X has 2 strokes, \ and /
			if strokesMatch == 2 {
				result++
			}
		}
	}

	log.Printf("The result is %d\n. Took: %d microseconds", result, time.Since(startTime).Microseconds())
}
