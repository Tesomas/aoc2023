package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const INPUT_URL = "https://adventofcode.com/2023/day/1/input"

func main() {
	lines := getInputLines()
	var num = 0
	for _, line := range *lines {
		newNum, _ := findDigitsForLine(line)
		num += newNum
	}
	fmt.Print(num)

}

func findDigitsForLine(line string) (int, error) {
	r, _ := regexp.Compile("[0-9]")
	firstMatch := r.Find([]byte(line))
	reversedString := reverseString(line)
	secondMatch := r.Find([]byte(reversedString))
	return strconv.Atoi(string(append(firstMatch, secondMatch...)))

}

func reverseString(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func readInputFile() *os.File {
	data, err := os.OpenFile("./input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	return data

}

func getInputLines() *[]string {
	var lineList []string
	scanner := bufio.NewScanner(readInputFile())
	for scanner.Scan() {
		lineList = append(lineList, scanner.Text())
	}
	return &lineList
}
