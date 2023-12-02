package main

import (
	"common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	gamePowerSum := 0
	lines := common.GetInputLines("/home/rbruenin/workspace/aoc2023/day2/input.txt")
	games := parseGames(lines)

	for _, game := range games {
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, round := range game.Rounds {
			if round.Red > maxRed {
				maxRed = round.Red
			}
			if round.Blue > maxBlue {
				maxBlue = round.Blue
			}
			if round.Green > maxGreen {
				maxGreen = round.Green
			}
		}
		gamePowerSum += maxRed * maxBlue * maxGreen
	}
	fmt.Println(gamePowerSum)
}

func parseGames(lines *[]string) []Game {
	var games []Game
	for _, line := range *lines {
		game := Game{}
		gameSplitString := strings.Split(line, ":")
		game.Id = getGameID(gameSplitString[0])
		rounds := strings.Split(gameSplitString[1], ";")
		for _, round := range rounds {
			game.Rounds = append(game.Rounds, getRoundFromString(round))
		}
		games = append(games, game)
	}
	return games

}
func getGameID(gameIDString string) int {
	r, _ := regexp.Compile("[0-9]+")
	matchedString := r.FindString(gameIDString)
	s, err := strconv.Atoi(matchedString)
	if err != nil {
		panic(err)
	}
	return s
}

func getRoundFromString(round string) Round {
	parsedRound := Round{}
	red := regexp.MustCompile(`(?P<redCount>\d+) red`)
	blue := regexp.MustCompile(`(?P<blueCount>\d+) blue`)
	green := regexp.MustCompile(`(?P<greenCount>\d+) green`)
	redMatch := red.FindStringSubmatch(round)
	blueMatch := blue.FindStringSubmatch(round)
	greenMatch := green.FindStringSubmatch(round)
	index := red.SubexpIndex("redCount")
	if redMatch != nil {
		parsedRound.Red, _ = strconv.Atoi(redMatch[index])
	} else {
		parsedRound.Red = 0
	}
	if blueMatch != nil {
		parsedRound.Blue, _ = strconv.Atoi(blueMatch[index])
	} else {
		parsedRound.Blue = 0
	}
	if greenMatch != nil {
		parsedRound.Green, _ = strconv.Atoi(greenMatch[index])
	} else {
		parsedRound.Green = 0
	}
	return parsedRound
}
