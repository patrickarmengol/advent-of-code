package day14

import (
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

func tiltNorth(grid [][]rune) {
	for c := range grid[0] {
		for r := range grid {
			if grid[r][c] == 'O' {
				for cur_r := r; cur_r > 0 && grid[cur_r-1][c] == '.'; cur_r-- {
					grid[cur_r-1][c] = 'O'
					grid[cur_r][c] = '.'
				}
			}
		}
	}
}

func tiltWest(grid [][]rune) {
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 'O' {
				for cur_c := c; cur_c > 0 && grid[r][cur_c-1] == '.'; cur_c-- {
					grid[r][cur_c-1] = 'O'
					grid[r][cur_c] = '.'
				}
			}
		}
	}
}

func tiltSouth(grid [][]rune) {
	for c := range grid[0] {
		for r := len(grid) - 1; r >= 0; r-- {
			if grid[r][c] == 'O' {
				for cur_r := r; cur_r < len(grid)-1 && grid[cur_r+1][c] == '.'; cur_r++ {
					grid[cur_r+1][c] = 'O'
					grid[cur_r][c] = '.'
				}
			}
		}
	}
}

func tiltEast(grid [][]rune) {
	for r := range grid {
		for c := len(grid[0]) - 1; c >= 0; c-- {
			if grid[r][c] == 'O' {
				for cur_c := c; cur_c < len(grid[0])-1 && grid[r][cur_c+1] == '.'; cur_c++ {
					grid[r][cur_c+1] = 'O'
					grid[r][cur_c] = '.'
				}
			}
		}
	}
}

func Part1(input string) (int, error) {
	grid := util.Gridify(input)

	// util.PrintGrid(grid)

	tiltNorth(grid)

	// fmt.Println()
	// util.PrintGrid(grid)

	total := 0
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 'O' {
				total += len(grid) - r
			}
		}
	}

	return total, nil
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)

	// util.PrintGrid(grid)

	cycleNum := 1_000_000_000

	seenGrids := map[string]int{}
	history := map[int]string{}

	for c := 1; c <= cycleNum; c++ {
		tiltNorth(grid)
		tiltWest(grid)
		tiltSouth(grid)
		tiltEast(grid)
		gs := util.SprintGrid(grid)
		if cs, ok := seenGrids[gs]; ok {
			looplen := c - cs
			grid = util.Gridify(history[cs+((cycleNum-cs)%looplen)])
			break
		} else {
			seenGrids[gs] = c
			history[c] = gs
		}
	}

	// fmt.Println()
	// util.PrintGrid(grid)

	total := 0
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == 'O' {
				total += len(grid) - r
			}
		}
	}

	return total, nil
}
