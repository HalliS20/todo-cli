package functions

import (
	"fmt"
	"os"
	"strings"
)

func FileAsString(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %s", fileName)
		panic(err)
	}
	return string(data)
}

func FileAsLines(fileName string) []string {
	fileString := FileAsString(fileName)
	strings.Split(fileString, "\n")
	fileString = fileString[:len(fileString)-1]
	return strings.Split(fileString, "\n")
}
