package main

import (
	"adventofcode2023/helper"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	result := 0
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	for _, card := range inputData {
		cardData := strings.Split(strings.Split(card, ":")[1], "|")
		winning := strings.Split(strings.TrimSpace(cardData[0]), " ")
		mine := strings.Split(strings.TrimSpace(cardData[1]), " ")

		correct := 0
		for _, number := range mine {
			if number == "" {
				continue
			}
			for _, winningNumber := range winning {
				if winningNumber == "" {
					continue
				}
				if strings.TrimSpace(number) == strings.TrimSpace(winningNumber) {
					correct++
				}
			}
		}

		if correct != 0 {
			points := math.Pow(2, float64(correct-1))
			result += int(points)
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}
