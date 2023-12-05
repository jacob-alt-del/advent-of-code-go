package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/atotto/clipboard"
)

//go:embed input.txt
var input string

type game struct {
	count       int
	winningNums []string
	playerNums  []string
}

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)

	gameMap := getMap(parsed)

	total := 0

	for _, game := range gameMap {
		wins := getWins(game)
		total += getScore(wins)
	}

	return total
}

func part2(input string) int {
	parsed := parseInput(input)

	tickets := 0

	gameMap := getMap(parsed)

	for i := 1; i <= len(parsed); i++ {
		wins := getWins(gameMap[i])
		if wins > 0 {
			gameMap = updateCounts(i, wins, gameMap)
		}
		tickets += gameMap[i].count
	}

	return tickets
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func getWins(input game) int {
	wins := 0

	for _, wn := range input.winningNums {
		for _, pn := range input.playerNums {
			if pn == wn {
				wins++
			}
		}
	}

	return wins
}

func getScore(wins int) int {
	score := 0

	if wins == 1 {
		score = 1
	} else if wins > 1 {
		score = int(math.Pow(2, float64(wins-1)))
	}

	return score
}

func getMap(input []string) map[int]game {
	output := map[int]game{}
	for i, val := range input {
		nums := strings.Split(strings.Split(val, ":")[1], "|")

		winningNumsStr := strings.Trim(nums[0], " ")
		playerNumsStr := strings.Trim(nums[1], " ")

		winningNumsStr = strings.ReplaceAll(winningNumsStr, "  ", " ")
		playerNumsStr = strings.ReplaceAll(playerNumsStr, "  ", " ")

		winningNums := strings.Split(winningNumsStr, " ")
		playerNums := strings.Split(playerNumsStr, " ")

		output[i+1] = game{
			count:       1,
			winningNums: winningNums,
			playerNums:  playerNums,
		}
	}

	return output
}

func updateCounts(prev int, wins int, gameMap map[int]game) map[int]game {
	for i := prev + 1; i <= prev+wins; i++ {
		s := gameMap[i]
		s.count += gameMap[prev].count
		gameMap[i] = s
	}
	return gameMap
}
