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
	_ = parsed

	symPositions := getLocations(parsed)

	total := 0

	for i, line := range parsed {
		minAdj := max(i-1, 0)
		maxAdj := min(i+2, len(parsed))
		nums := checkNums(line, symPositions[minAdj:maxAdj])
		for _, val := range nums {
			total += val
		}
	}

	return total
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	totalRatios := 0
	for n, line := range parsed {
		aboveLine := ""
		belowLine := ""

		if n-1 >= 0 {
			aboveLine = parsed[n-1]
		}
		if n+2 <= len(parsed) {
			belowLine = parsed[n+1]
		}

		totalRatios += checkGears(line, aboveLine, belowLine)
	}

	return totalRatios
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func getLocations(input []string) [][]int {
	symPositions := [][]int{}
	for _, line := range input {
		posSym := []int{}
		for i, char := range line {
			if !unicode.IsDigit(char) && char != '.' {
				posSym = append(posSym, i)
			}
		}
		symPositions = append(symPositions, posSym)
	}
	return symPositions
}

func checkNums(line string, symPositions [][]int) []int {
	output := []int{}
	onDigit := true
	numStr := ""
	includeNum := false

	for i, numVal := range line {

		// build full number
		if unicode.IsDigit(numVal) {
			numStr = numStr + string(numVal)

			// check if number is included
			for _, symLine := range symPositions {
				for _, symPos := range symLine {
					if symPos-i < 2 && symPos-i > -2 {
						includeNum = true
					}
				}
			}

		} else {
			onDigit = false
		}

		if !onDigit && numStr != "" {
			num := cast.ToInt(numStr)
			if includeNum {
				output = append(output, num)
				includeNum = false
			}
			numStr = ""
		}
		onDigit = true
	}

	// catch last num
	if numStr != "" && includeNum {
		num := cast.ToInt(numStr)
		output = append(output, num)
	}

	return output
}

func checkGears(currentRow string, aboveRow string, belowRow string) int {
	total := 0

	for i, val := range currentRow {
		gears := []int{}

		if val != '*' {
			continue
		}

		// search current row left
		baseNum := ""
		for _, n := range []int{1, 2, 3} {
			if i-n >= 0 && unicode.IsDigit(rune(currentRow[i-n])) {
				baseNum = string(currentRow[i-n]) + baseNum
				continue
			}
			break
		}
		if baseNum != "" {
			gear := cast.ToInt(baseNum)
			gears = append(gears, gear)
		}

		// search current row right
		baseNum = ""
		for _, n := range []int{1, 2, 3} {
			if i-n <= len(currentRow) && unicode.IsDigit(rune(currentRow[i+n])) {
				baseNum += string(currentRow[i+n])
				continue
			}
			break
		}
		if baseNum != "" {
			gear := cast.ToInt(baseNum)
			gears = append(gears, gear)
		}

		// search adjacent rows
		for _, row := range []string{aboveRow, belowRow} {
			if row == "" {
				continue
			}
			var skipNext bool
			for _, n := range []int{i - 1, i, i + 1} {
				if skipNext {
					if !unicode.IsDigit(rune(row[n])) {
						skipNext = false
					}
					continue
				}

				if unicode.IsDigit(rune(row[n])) {
					baseNum := string(row[n])
					// search left
					if n-1 >= 0 && unicode.IsDigit(rune(row[n-1])) {
						baseNum = string(row[n-1]) + baseNum
						if n-2 >= 0 && unicode.IsDigit(rune(row[n-2])) {
							baseNum = string(row[n-2]) + baseNum
						}
					}
					// search right
					if n+1 <= len(row) && unicode.IsDigit(rune(row[n+1])) {
						baseNum += string(row[n+1])
						if n+2 <= len(row) && unicode.IsDigit(rune(row[n+2])) {
							baseNum += string(row[n+2])
						}
					}

					skipNext = true
					gear := cast.ToInt(baseNum)
					gears = append(gears, gear)
				}
			}
		}

		if len(gears) > 1 {
			total += gears[0] * gears[1]
		}
	}

	return total
}
