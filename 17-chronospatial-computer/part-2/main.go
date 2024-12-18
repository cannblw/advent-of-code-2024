package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func readInput(registerB, registerC *int, program *[]byte) {
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
			continue
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

func areByteArraysEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true

}

func main() {
	startTime := time.Now()

	var program []byte
	var initialRegisterB, initialRegisterC int
	readInput(&initialRegisterB, &initialRegisterC, &program)

	currentRegisterA := -1

	// May I be forgiven for this
	for {
		pointerIncrease := 2
		var registerA, registerB, registerC int

		registerA, currentRegisterA = currentRegisterA+1, currentRegisterA+1
		var output []byte
		registerB = initialRegisterB
		registerC = initialRegisterC

		if registerA%10000000 == 0 {
			fmt.Println("Trying with register A =", registerA, "| seconds so far =", int(time.Since(startTime).Seconds()))
		}

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
				output = append(output, byte(combo%8))
			case 6:
				combo := getComboOperand(operand, registerA, registerB, registerC)
				registerB = registerA / (1 << combo)
			case 7:
				combo := getComboOperand(operand, registerA, registerB, registerC)
				registerC = registerA / (1 << combo)
			}
		}

		if areByteArraysEqual(program, output) {
			break
		}

	}

	log.Printf("The result is %d. Took: %d seconds", currentRegisterA, int(time.Since(startTime).Seconds()))
}
