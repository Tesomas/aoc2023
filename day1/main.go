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
		newNum, err := findDigitsForLine(line)
		if err != nil {
			panic(err)
		}
		fmt.Println(newNum)
		num += newNum
	}
	fmt.Println(num)

}

func findDigitsForLine(line string) (int, error) {
	line = replaceSpelledOutDigits(line)
	r, _ := regexp.Compile("[0-9]")
	firstMatch := r.Find([]byte(line))
	reversedString := reverseString(line)
	secondMatch := r.Find([]byte(reversedString))
	return strconv.Atoi(string(append(firstMatch, secondMatch...)))
}

func replaceSingleDigit(line string, matchString string, replacementString string) string {
	regex, _ := regexp.Compile(matchString)
	return regex.ReplaceAllString(line, replacementString)
}

func replaceSpelledOutDigits(line string) string {
	line = replaceSingleDigit(line, "one", "on1e")
	line = replaceSingleDigit(line, "two", "tw2o")
	line = replaceSingleDigit(line, "three", "thr3ee")
	line = replaceSingleDigit(line, "four", "fo4ur")
	line = replaceSingleDigit(line, "five", "fi5ve")
	line = replaceSingleDigit(line, "six", "si6x")
	line = replaceSingleDigit(line, "seven", "sev7en")
	line = replaceSingleDigit(line, "eight", "ei8ght")
	line = replaceSingleDigit(line, "nine", "ni9ne")
	return line
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
