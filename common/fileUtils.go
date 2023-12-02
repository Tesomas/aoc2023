package common

import (
	"bufio"
	"os"
)

func readInputFile(filePath string) *os.File {
	data, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	return data

}

func GetInputLines(filePath string) *[]string {
	var lineList []string
	scanner := bufio.NewScanner(readInputFile(filePath))
	for scanner.Scan() {
		lineList = append(lineList, scanner.Text())
	}
	return &lineList
}
