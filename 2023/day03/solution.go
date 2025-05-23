package day03

import (
	"strconv"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

var numerics = util.NewSet[rune]([]rune("0123456789")...)

func Part1(input string) (int, error) {
	grid := util.Gridify(input)

	total := 0
	allParts := util.NewSet[part]()
	for r := range grid {
		for c := range grid[r] {
			if !numerics.Contains(grid[r][c]) && grid[r][c] != '.' {
				allParts = allParts.Union(findNeighborParts(grid, r, c))
			}
		}
	}
	for part := range allParts {
		total += part.val
	}

	return total, nil
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)

	total := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '*' {
				neighbors := findNeighborParts(grid, r, c).Members()
				if len(neighbors) == 2 {
					total += neighbors[0].val * neighbors[1].val
				}
			}
		}
	}

	return total, nil
}

func findNeighborParts(grid [][]rune, symR, symC int) util.Set[part] {
	parts := util.NewSet[part]()

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			r := symR + dr
			c := symC + dc
			if numerics.Contains(grid[r][c]) {
				n := []rune{grid[r][c]}
				nc := c
				for i := c - 1; i >= 0; i-- {
					if numerics.Contains(grid[r][i]) {
						n = append([]rune{grid[r][i]}, n...)
						nc = i
					} else {
						break
					}
				}
				for i := c + 1; i < len(grid[r]); i++ {
					if numerics.Contains(grid[r][i]) {
						n = append(n, grid[r][i])
					} else {
						break
					}
				}
				nn, err := strconv.Atoi(string(n))
				if err != nil {
					panic("neighbor contains non numerical runes")
				}
				parts.Add(part{val: nn, row: r, col: nc})
			}
		}
	}
	return parts
}

type part struct {
	val int
	row int
	col int
}
