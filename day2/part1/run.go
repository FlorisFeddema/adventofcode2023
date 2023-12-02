package main

import (
	"adventofcode2023/helper"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

func main() {
	result := 0
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	for i, game := range inputData {
		sets := strings.Split(game, ":")[1]
		scores := strings.Split(sets, ";")
		if checkScoreValid(scores) {
			result += i + 1
		}
	}

	elapsed := time.Since(start)
	slog.Info("Result", "answer", result, "time", elapsed)
}

func checkScoreValid(scores []string) bool {
	for _, score := range scores {
		amounts := strings.Split(score, ",")
		for _, amount := range amounts {
			var maxAmount int
			if strings.Contains(amount, "red") {
				maxAmount = 12
			} else if strings.Contains(amount, "green") {
				maxAmount = 13
			} else if strings.Contains(amount, "blue") {
				maxAmount = 14
			}
			amountNumber, _ := strconv.Atoi(strings.Split(amount, " ")[1])
			if amountNumber > maxAmount {
				return false
			}
		}
	}
	return true
}
