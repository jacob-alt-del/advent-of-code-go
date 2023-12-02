package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"advent-of-code-go/pkg/cast"

	"github.com/atotto/clipboard"
)

//go:embed input.txt
var input string
var nonAlphanumericRegex = regexp.MustCompile(`[^0-9 ]+`)

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
	maxr := 12
	maxg := 13
	maxb := 14
	count := 0

	parsed := parseInput(input)
	mp := getMap(parsed)

	for i, game := range mp {
		r, g, b := calcMax(game)
		if r > maxr || g > maxg || b > maxb {
			continue
		}
		count += i
	}

	return count
}

func part2(input string) int {
	power := 0

	parsed := parseInput(input)
	mp := getMap(parsed)

	for _, game := range mp {
		r, g, b := calcMax(game)
		power += r * g * b
	}

	return power
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func getMap(input []string) map[int][]string {
	mp := make(map[int][]string)
	for i, val := range input {
		game := strings.Split(val, ":")[1]
		game = strings.ReplaceAll(game, " ", "")
		mp[i+1] = strings.Split(game, ";")
	}
	return mp
}

func calcMax(game []string) (int, int, int) {
	maxr := 0
	maxg := 0
	maxb := 0
	for _, r := range game {
		cubes := strings.Split(r, ",")
		for _, col := range cubes {
			str := nonAlphanumericRegex.ReplaceAllString(col, "")
			num := cast.ToInt(str)
			if strings.Contains(col, "red") && num > maxr {
				maxr = num
			}
			if strings.Contains(col, "green") && num > maxg {
				maxg = num
			}
			if strings.Contains(col, "blue") && num > maxb {
				maxb = num
			}
		}
	}
	return maxr, maxg, maxb
}
