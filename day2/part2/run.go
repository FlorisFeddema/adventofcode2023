package main

import (
	"adventofcode2023/helper"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	result := 0
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	for _, game := range inputData {
		sets := strings.Split(game, ":")[1]
		scores := strings.Split(sets, ";")
		total := checkScoreValid(scores)
		result += total
	}

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}

func checkScoreValid(scores []string) int {
	red, green, blue := 0, 0, 0
	for _, score := range scores {
		amounts := strings.Split(score, ",")
		for _, amount := range amounts {
			if strings.Contains(amount, "red") {
				amountNumber, _ := strconv.Atoi(strings.Split(amount, " ")[1])
				if amountNumber > red {
					red = amountNumber
				}
			} else if strings.Contains(amount, "green") {
				amountNumber, _ := strconv.Atoi(strings.Split(amount, " ")[1])
				if amountNumber > green {
					green = amountNumber
				}
			} else if strings.Contains(amount, "blue") {
				amountNumber, _ := strconv.Atoi(strings.Split(amount, " ")[1])
				if amountNumber > blue {
					blue = amountNumber
				}
			}
		}
	}
	return red * green * blue
}
