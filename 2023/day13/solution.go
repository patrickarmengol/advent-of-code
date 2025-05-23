package day13

import (
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

func Part1(input string) (int, error) {
	sections := strings.Split(input, "\n\n")

	totalRows := 0
	totalCols := 0

	for _, sec := range sections {
		grid := util.Gridify(sec)
		// util.PrintGrid(grid)
		// fmt.Println()

		length := len(grid)
		width := len(grid[0])

		// assuming one instance of symmetry row-wise OR col-wise

		for r := 0; r < length-1; r++ {
			eq := true
			dr := 0
			for {
				up := r - dr
				down := r + 1 + dr
				if up < 0 || down >= length {
					break
				}
				for c := 0; c < width; c++ {
					if grid[up][c] != grid[down][c] {
						eq = false
						break
					}
				}
				dr++
			}
			if eq {
				totalRows += r + 1
				break
			}
		}

		for c := 0; c < width-1; c++ {
			eq := true
			dc := 0
			for {
				left := c - dc
				right := c + 1 + dc
				if left < 0 || right >= width {
					break
				}
				for r := 0; r < length; r++ {
					if grid[r][left] != grid[r][right] {
						eq = false
						break
					}
				}
				dc++
			}
			if eq {
				totalCols += c + 1
				break
			}
		}

	}

	return totalCols + (100 * totalRows), nil
}

func Part2(input string) (int, error) {
	sections := strings.Split(input, "\n\n")

	totalRows := 0
	totalCols := 0

	for _, sec := range sections {
		grid := util.Gridify(sec)
		// util.PrintGrid(grid)
		// fmt.Println()

		length := len(grid)
		width := len(grid[0])

		// assuming one instance of symmetry row-wise OR col-wise

		for r := 0; r < length-1; r++ {
			smudges := 0
			dr := 0
			for {
				up := r - dr
				down := r + 1 + dr
				if up < 0 || down >= length {
					break
				}
				for c := 0; c < width; c++ {
					if grid[up][c] != grid[down][c] {
						smudges += 1
					}
				}
				if smudges > 1 {
					break
				}
				dr++
			}
			if smudges == 1 {
				totalRows += r + 1
				break
			}
		}

		for c := 0; c < width-1; c++ {
			smudges := 0
			dc := 0
			for {
				left := c - dc
				right := c + 1 + dc
				if left < 0 || right >= width {
					break
				}
				for r := 0; r < length; r++ {
					if grid[r][left] != grid[r][right] {
						smudges += 1
					}
				}
				if smudges > 1 {
					break
				}
				dc++
			}
			if smudges == 1 {
				totalCols += c + 1
				break
			}
		}

	}

	return totalCols + (100 * totalRows), nil
}
