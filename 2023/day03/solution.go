package day03

import (
	_ "embed"
	"regexp"
	"slices"
	"strconv"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

var numRX = regexp.MustCompile(`\d+`)

func Part1(input string) (string, error) {
	lines := parse.Lines(input)

	// populate a 2d grid
	grid := util.Make2D[rune](len(lines), len(lines[0]))
	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
		}
	}

	total := 0
	for i, line := range lines {
		for _, m := range numRX.FindAllStringIndex(line, -1) {
			if symbolNeighbor(grid, i, m[0], m[1]) {
				num, err := strconv.Atoi(line[m[0]:m[1]])
				if err != nil {
					panic("number regex pattern broke")
				}
				total += num
			}
		}
	}

	return strconv.Itoa(total), nil
}

func symbolNeighbor(grid [][]rune, lineIndex int, numStart int, numEnd int) bool {
	gLength := len(grid)
	gWidth := len(grid[0])

	for i := lineIndex - 1; i < lineIndex+2; i++ {
		for j := numStart - 1; j < numEnd+1; j++ {
			if i < 0 || i >= gLength || j < 0 || j >= gWidth {
				continue
			}
			if !slices.Contains([]rune("0123456789."), grid[i][j]) {
				return true
			}
		}
	}
	return false
}

func Part2(input string) (string, error) {
	lines := parse.Lines(input)

	gearMap := map[gear][]part{}

	grid := util.Make2D[rune](len(lines), len(lines[0]))
	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
		}
	}

	for i, line := range lines {
		for _, m := range numRX.FindAllStringIndex(line, -1) {
			gearNs := gearNeighbors(grid, i, m[0], m[1])
			for _, gearN := range gearNs {
				gearMap[gearN] = append(gearMap[gearN], part{row: i, startCol: m[0], endCol: m[1]})
			}
		}
	}

	total := 0
	for _, parts := range gearMap {
		if len(parts) == 2 {
			part1 := parts[0]
			part2 := parts[1]
			part1Val, err := strconv.Atoi(lines[part1.row][part1.startCol:part1.endCol])
			if err != nil {
				panic("invalid num regex")
			}
			part2Val, err := strconv.Atoi(lines[part2.row][part2.startCol:part2.endCol])
			if err != nil {
				panic("invalid num regex")
			}
			product := part1Val * part2Val
			total += product
		}
	}

	return strconv.Itoa(total), nil
}

func gearNeighbors(grid [][]rune, lineIndex int, numStart int, numEnd int) []gear {
	gLength := len(grid)
	gWidth := len(grid[0])

	gears := []gear{}

	for i := lineIndex - 1; i < lineIndex+2; i++ {
		for j := numStart - 1; j < numEnd+1; j++ {
			if i < 0 || i >= gLength || j < 0 || j >= gWidth {
				continue
			}
			if grid[i][j] == '*' {
				gears = append(gears, gear{row: i, col: j})
			}
		}
	}
	return gears
}

type gear struct {
	row int
	col int
}

type part struct {
	row      int
	startCol int
	endCol   int
}
