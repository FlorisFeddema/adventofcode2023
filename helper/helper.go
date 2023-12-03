package helper

import (
	"os"
	"strconv"
	"strings"
)

func ReadFileToStringArray(fileName string) []string {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(fileData))
	return strings.Split(data, "\n")
}

func StringToInt(input string) int {
	tmp, _ := strconv.Atoi(input)
	return tmp
}
