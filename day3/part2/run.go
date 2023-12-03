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
			if inputData[x][y] == '*' {
				values := hasAdjacentValues(inputData, x, y)
				if len(values) == 2 {
					ratio := values[0] * values[1]
					result += ratio
				}
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}

func hasAdjacentValues(inputData []string, x int, y int) []int {
	var values []int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if x+i < 0 || y+j < 0 || x+i >= len(inputData) || y+j >= len(inputData[x]) {
				continue
			}
			if strings.Contains("1234567890", string(inputData[x+i][y+j])) {
				value := string(inputData[x+i][y+j])
				digitsN := 1
				for y+j-digitsN >= 0 && strings.Contains("1234567890", string(inputData[x+i][y+j-digitsN])) {
					value = string(inputData[x+i][y+j-digitsN]) + value
					digitsN++
				}
				digitsP := 1
				for y+j+digitsP <= len(inputData[x+i])-1 && strings.Contains("1234567890", string(inputData[x+i][y+j+digitsP])) {
					value += string(inputData[x+i][y+j+digitsP])
					digitsP++
				}
				values = append(values, helper.StringToInt(value))
				j += digitsP - 1
			}
		}
	}
	return values
}
