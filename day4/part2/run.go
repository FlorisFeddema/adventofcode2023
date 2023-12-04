package main

import (
	"adventofcode2023/helper"
	"fmt"
	"strings"
	"time"
)

func main() {
	result := 0
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	amountOfCards := []int{}
	for range inputData {
		amountOfCards = append(amountOfCards, 1)
	}

	for i, card := range inputData {
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

		for j := 1; j <= correct; j++ {
			amountOfCards[i+j] += amountOfCards[i]
		}
	}

	for _, amount := range amountOfCards {
		result += amount
	}

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}
