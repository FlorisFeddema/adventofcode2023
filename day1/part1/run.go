package main

import (
	"adventofcode2023/helper"
	"log/slog"
	"strconv"
	"time"
)

func main() {
	result := 0
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	for _, line := range inputData {
		first := 0
		last := 0
		for _, char := range line {
			intChar, err := strconv.Atoi(string(char))
			if err == nil {
				if first == 0 {
					first = intChar
				}
				last = intChar
			}
		}
		result += first*10 + last
	}

	elapsed := time.Since(start)
	slog.Info("Result", "answer", result, "time", elapsed)
}
