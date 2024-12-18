package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func readInput(registerA, registerB, registerC *int, program *[]byte) {
	const (
		RegisterA int = 0
		RegisterB int = 1
		RegisterC int = 2
		Program   int = 3
	)

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	r := regexp.MustCompile(`:\s(.+$)`)
	if err != nil {
		log.Fatalf("Error compiling regexp: %v", err)
	}

	scanner := bufio.NewScanner(file)
	readMode := -1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		readMode++

		matches := r.FindStringSubmatch(line)

		switch readMode {
		case RegisterA:
			*registerA, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[1], err)
			}
		case RegisterB:
			*registerB, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[1], err)
			}
		case RegisterC:
			*registerC, err = strconv.Atoi(matches[1])
			if err != nil {
				log.Fatalf("Error converting item %s to int: %v", matches[1], err)
			}
		case Program:
			splitStr := strings.Split(matches[1], ",")
			for _, c := range splitStr {
				*program = append(*program, c[0]-'0')
			}
		}
	}
}

func getComboOperand(literalOperand byte, registerA, registerB, registerC int) int {
	switch literalOperand {
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	}

	return int(literalOperand)
}

func main() {
	startTime := time.Now()
	var output []string

	var registerA, registerB, registerC int
	var program []byte
	readInput(&registerA, &registerB, &registerC, &program)

	pointerIncrease := 2
	for i := 0; i < len(program)-1; i += pointerIncrease {
		pointerIncrease = 2

		instruction := program[i]
		operand := program[i+1]

		switch instruction {
		case 0:
			combo := getComboOperand(operand, registerA, registerB, registerC)
			registerA /= 1 << combo
		case 1:
			registerB = registerB ^ int(operand)
		case 2:
			combo := getComboOperand(operand, registerA, registerB, registerC)
			registerB = combo % 8
		case 3:
			if registerA != 0 {
				pointerIncrease = 0
				i = int(operand)
			}
		case 4:
			registerB = registerB ^ registerC
		case 5:
			combo := getComboOperand(operand, registerA, registerB, registerC)
			output = append(output, strconv.FormatInt(int64(combo%8), 10))
		case 6:
			combo := getComboOperand(operand, registerA, registerB, registerC)
			registerB = registerA / (1 << combo)
		case 7:
			combo := getComboOperand(operand, registerA, registerB, registerC)
			registerC = registerA / (1 << combo)
		}
	}

	log.Printf("The result is %s. Took: %d microseconds", strings.Join(output, ","), time.Since(startTime).Microseconds())
}
