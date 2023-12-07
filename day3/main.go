package main

import (
	"common"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := common.GetInputLines("/home/rbruenin/workspace/aoc2023/day3/input.txt")
	grid := make([][]int, 140)
	for i := range grid {
		grid[i] = make([]int, 140)
	}
	markSymbolRegionsInGrid(&grid, lines)
	for _, line := range grid {
		fmt.Println(line)
	}
	fmt.Println(getSumOfLines(&grid, lines))
}

func getSumOfLines(grid *[][]int, lines *[]string) int {
	sum := 0
	for index, line := range *lines {
		regexp := regexp.MustCompile(`[0-9]+`)
		matches := regexp.FindAllStringIndex(line, -1)
		if matches != nil {
			for _, match := range matches {
				matchShouldCount := false

				for x := match[0]; x < match[1]; x++ {
					if (*grid)[index][x] == 1 {
						matchShouldCount = true
						break
					}
				}

				if matchShouldCount {
					num, err := strconv.Atoi(line[match[0]:match[1]])
					if err != nil {
						panic(err)
					}
					sum += num
				}
			}

		}
		fmt.Println(index, sum)
	}
	return sum
}

func markSymbolRegionsInGrid(grid *[][]int, lines *[]string) {
	for index, line := range *lines {
		regexp := regexp.MustCompile(`[\*]`)
		matches := regexp.FindAllStringIndex(line, -1)

		if matches != nil {
			for _, match := range matches {
				if match[0] > 0 {
					(*grid)[index+1][match[0]-1] = 1
					(*grid)[index][match[0]-1] = 1
					if index > 0 {
						(*grid)[index-1][match[0]-1] = 1
					}
				}

				if index > 0 {
					(*grid)[index-1][match[0]] = 1
					(*grid)[index-1][match[0]+1] = 1

				}
				(*grid)[index][match[0]] = 1
				(*grid)[index][match[0]+1] = 1
				(*grid)[index+1][match[0]] = 1
				(*grid)[index+1][match[0]+1] = 1

			}
		}
	}
}

// convert this to a lambda
