package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"

	"advent-of-code-go/pkg/cast"

	"github.com/atotto/clipboard"
)

//go:embed input.txt
var input string

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

	var output int
	nums := []int{}

	for _, line := range parsed {
		var first int = 0
		var last int = 0

		for i := 0; i < len(line); i++ {
			val := line[i]
			if unicode.IsDigit(rune(val)) {
				if first == 0 {
					first = cast.ToInt(string(val))
				}
				last = cast.ToInt(string(val))
			}
		}

		nums = append(nums, first*10+last)
	}

	for _, val := range nums {
		output += val
	}

	return output
}

func part2(input string) int {
	parsed := parseInput(input)

	var output int
	nums := []int{}

	for _, line := range parsed {
		numList := getList(line)
		first := numList[0]
		last := numList[len(numList)-1]
		nums = append(nums, first*10+last)
	}

	for _, val := range nums {
		output += val
	}

	return output
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func getList(input string) []int {
	list := []int{}

	for i := 0; i < len(input); i++ {
		num := input[i]

		if unicode.IsDigit(rune(num)) {
			list = append(list, cast.ToInt(string(num)))
			continue
		}

		numbers := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}

		// check 5 letter nums
		if len(input)-i > 4 {
			n, ok := numbers[string(input[i:i+5])]
			if ok {
				list = append(list, n)
				continue
			}
		}

		// check 4 letter nums
		if len(input)-i > 3 {
			n, ok := numbers[string(input[i:i+4])]
			if ok {
				list = append(list, n)
				continue
			}
		}

		// check 3 letter nums
		if len(input)-i > 2 {
			n, ok := numbers[string(input[i:i+3])]
			if ok {
				list = append(list, n)
				continue
			}
		}

	}

	return list
}
