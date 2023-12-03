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

	for x := 0; x <= len(inputData)-1; x++ {
		for y := 0; y <= len(inputData[x])-1; y++ {
			if strings.Contains("1234567890", string(inputData[x][y])) {
				value := string(inputData[x][y])
				digits := 1
				for y+digits <= len(inputData[x])-1 && strings.Contains("1234567890", string(inputData[x][y+digits])) {
					value += string(inputData[x][y+digits])
					digits++
				}
				if hasAdjacent(inputData, x, y, digits) {
					result += helper.StringToInt(value)
				}
				y += digits - 1
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}

func hasAdjacent(inputData []string, x int, y int, digits int) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < digits+1; j++ {
			if x+i < 0 || y+j < 0 || x+i >= len(inputData) || y+j >= len(inputData[x]) {
				continue
			}
			if inputData[x+i][y+j] != '.' && !strings.Contains("1234567890", string(inputData[x+i][y+j])) {
				return true
			}
		}
	}
	return false
}
