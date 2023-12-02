package main

import (
	"adventofcode2023/helper"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputData := helper.ReadFileToStringArray("input.txt")
	result := 0
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	start := time.Now()

	for _, line := range inputData {
		first := 0
		last := 0
		for i, char := range line {
			intChar, err := strconv.Atoi(string(char))
			if err == nil {
				if first == 0 {
					first = intChar
				}
				last = intChar
			} else {
				restOfLine := line[i:]
				for i, number := range numbers {
					if strings.HasPrefix(restOfLine, number) {
						if first == 0 {
							first = i + 1
						}
						last = i + 1
						break
					}
				}
			}
		}
		result += first*10 + last
	}

	elapsed := time.Since(start)
	slog.Info("Result", "answer", result, "time", elapsed)
}
