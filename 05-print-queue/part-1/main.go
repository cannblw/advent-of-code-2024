package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const ruleSeparator string = "|"

type Rule [2]int

func parseRule(rules []Rule, line string) []Rule {
	var rule Rule

	splitLine := strings.Split(line, ruleSeparator)

	firstRuleInt, err := strconv.Atoi(splitLine[0])
	if err != nil {
		log.Fatalf("Error converting item %s to int: %v", splitLine[0], err)
	}

	secondRuleInt, err := strconv.Atoi(splitLine[1])
	if err != nil {
		log.Fatalf("Error converting item %s to int: %v", splitLine[0], err)
	}

	rule[0] = firstRuleInt
	rule[1] = secondRuleInt

	return append(rules, rule)
}

func computeLine(rules []Rule, line string) (bool, int) {
	itemsStr := strings.Split(line, ",")
	items := make([]int, len(itemsStr))
	for i, _ := range itemsStr {
		c, err := strconv.Atoi(itemsStr[i])
		if err != nil {
			log.Fatalf("Error converting item %s to int: %v", itemsStr[i], err)
		}
		items[i] = c
	}

	for _, rule := range rules {
		a := rule[0]
		b := rule[1]

		aIdx := slices.Index(items, a)
		bIdx := slices.Index(items, b)

		if aIdx == -1 || bIdx == -1 {
			continue
		}

		if aIdx > bIdx {
			return false, 0
		}
	}

	return true, items[len(items)/2]
}

func main() {
	startTime := time.Now()
	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	readingRules := true
	rules := make([]Rule, 0)
	result := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			readingRules = false
			continue
		}

		if readingRules {
			rules = parseRule(rules, line)
		} else {
			isSorted, middleElement := computeLine(rules, line)
			if isSorted {
				result += middleElement
			}
		}
	}

	log.Printf("The result is %d\n. Took: %d microseconds", result, time.Since(startTime).Microseconds())
}
