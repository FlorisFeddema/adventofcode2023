package main

import (
	"adventofcode2023/helper"
	"fmt"
	"time"
)

func main() {
	result := 0
	inputData := helper.ReadFileToStringArray("input.txt")

	start := time.Now()

	// TODO: add code

	elapsed := time.Since(start)

	fmt.Printf("Result: %[1]d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}
