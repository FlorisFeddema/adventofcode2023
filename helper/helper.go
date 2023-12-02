package helper

import (
	"os"
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
