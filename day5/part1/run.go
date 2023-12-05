package main

import (
	"adventofcode2023/helper"
	"fmt"
	"strings"
	"time"
)

func main() {
	result := 2147483647
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	seeds := strings.Split(strings.TrimPrefix(inputData[0], "seeds: "), " ")

	var mappings [][]string
	mappings = append(mappings, []string{})

	mappingCount := 0
	for i := 3; i < len(inputData); i++ {
		if inputData[i] == "" {
			i++
			mappingCount++
			mappings = append(mappings, []string{})
			continue
		}
		mappings[mappingCount] = append(mappings[mappingCount], inputData[i])
	}

	for _, seed := range seeds {
		value := helper.StringToInt(seed)
		for _, mapping := range mappings {
			for _, typeRange := range mapping {
				rangeValues := strings.Split(typeRange, " ")
				destinationRange := helper.StringToInt(rangeValues[0])
				sourceRange := helper.StringToInt(rangeValues[1])
				rangeLength := helper.StringToInt(rangeValues[2])
				if value >= sourceRange && value < sourceRange+rangeLength {
					value = destinationRange - sourceRange + value
					break
				}
			}
		}
		if value < result {
			result = value
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}
