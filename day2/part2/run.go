package main

import (
	"adventofcode2023/helper"
	"log/slog"
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
	slog.Info("Result", "answer", result, "time", elapsed)
}

func checkScoreValid(scores []string) int {
	red, green, blue := 0, 0, 0
	for _, score := range scores {
		amounts := strings.Split(score, ",")
		for _, amount := range amounts {
			red = getAmountForColor(amount, "red", red)
			green = getAmountForColor(amount, "green", green)
			blue = getAmountForColor(amount, "blue", blue)
		}
	}
	return red * green * blue
}

func getAmountForColor(input string, color string, currentValue int) int {
	if strings.Contains(input, color) {
		amountNumber := helper.StringToInt(strings.Split(input, " ")[1])
		if amountNumber > currentValue {
			currentValue = amountNumber
		}
	}
	return currentValue
}
