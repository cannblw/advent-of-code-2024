package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	const target string = "XMAS"

	file, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	lines := strings.Split(string(file), "\n")

	result := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			item := lines[i][j]

			if item != target[0] {
				continue
			}

			// u and v act like a unit vector in the direction of the solution
			for u := -1; u <= 1; u++ {
				for v := -1; v <= 1; v++ {
					if u == 0 && v == 0 {
						continue
					}

					for w := 1; w < len(target); w++ {
						nextVerticalIdx := i + u*w
						nextHorizontalIdx := j + v*w

						if nextVerticalIdx < 0 || nextHorizontalIdx < 0 || nextVerticalIdx >= len(lines) || nextHorizontalIdx >= len(lines[i]) {
							break
						}

						if lines[nextVerticalIdx][nextHorizontalIdx] != target[w] {
							break
						} else if w == len(target)-1 {
							result++
							break
						}
					}
				}
			}
		}
	}

	log.Printf("The result is %d\n. Took: %d microseconds", result, time.Since(startTime).Microseconds())
}
