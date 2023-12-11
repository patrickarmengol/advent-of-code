package day11

import (
	"fmt"
	"slices"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

type galaxy struct {
	row int
	col int
}

func findEmptyRowsAndCols(grid [][]rune) ([]int, []int) {
	// check rows
	emptyRows := []int{}
	for r := range grid {
		isEmpty := true
		for _, point := range grid[r] {
			if point != '.' {
				isEmpty = false
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, r)
		}
	}

	// check columns
	emptyCols := []int{}
	for c := range grid[0] {
		isEmpty := true
		for r := range grid {
			point := grid[r][c]
			if point != '.' {
				isEmpty = false
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, c)
		}
	}

	return emptyRows, emptyCols
}

func findGalaxies(grid [][]rune) []galaxy {
	// iterate through grid for '#' characters
	galaxies := []galaxy{}
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] == '#' {
				galaxies = append(galaxies, galaxy{r, c})
			}
		}
	}
	return galaxies
}

func distance(a, b galaxy, emptyRows, emptyCols []int, expandRate int) int {
	// i don't remember the name for this method (maybe cartesian distance?; not euclidean distance)
	// abs(y1-y2) + abs(x1-x2)
	srow, brow := a.row, b.row
	if srow > brow {
		srow, brow = brow, srow
	}
	scol, bcol := a.col, b.col
	if scol > bcol {
		scol, bcol = bcol, scol
	}

	diffRow := (brow - srow)
	diffCol := (bcol - scol)

	// if we pass over expanded space increment by expansion rate
	for i := srow + 1; i < brow; i++ {
		if slices.Contains(emptyRows, i) {
			diffRow += expandRate - 1
		}
	}
	for i := scol + 1; i < bcol; i++ {
		if slices.Contains(emptyCols, i) {
			diffCol += expandRate - 1
		}
	}

	return diffRow + diffCol
}

func printGrid(grid [][]rune) {
	for r := range grid {
		for c := range grid[0] {
			fmt.Printf("%c", grid[r][c])
		}
		fmt.Println()
	}
	fmt.Println()
}

func Part1(input string) (int, error) {
	grid := util.Gridify(input)

	emptyRows, emptyCols := findEmptyRowsAndCols(grid)

	galaxies := findGalaxies(grid)

	total := 0
	// iterate through combinations
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += distance(galaxies[i], galaxies[j], emptyRows, emptyCols, 2)
		}
	}
	return total, nil
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)

	emptyRows, emptyCols := findEmptyRowsAndCols(grid)

	galaxies := findGalaxies(grid)

	total := 0
	// iterate through combinations
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += distance(galaxies[i], galaxies[j], emptyRows, emptyCols, 1000000)
		}
	}

	return total, nil
}
