package day16

import (
	"slices"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

type coord struct {
	row int
	col int
}

type beam struct {
	coord
	enteredDir rune
}

var dm = map[rune]coord{
	'n': {-1, 0},
	's': {1, 0},
	'w': {0, -1},
	'e': {0, 1},
}

var rm = map[rune]map[rune][]rune{
	'.': {
		'n': {'n'},
		's': {'s'},
		'w': {'w'},
		'e': {'e'},
	},
	'/': {
		'n': {'e'},
		's': {'w'},
		'w': {'s'},
		'e': {'n'},
	},
	'\\': {
		'n': {'w'},
		's': {'e'},
		'w': {'n'},
		'e': {'s'},
	},
	'|': {
		'n': {'n'},
		's': {'s'},
		'w': {'n', 's'},
		'e': {'n', 's'},
	},
	'-': {
		'n': {'w', 'e'},
		's': {'w', 'e'},
		'w': {'w'},
		'e': {'e'},
	},
}

func Part1(input string) (int, error) {
	grid := util.Gridify(input)
	length := len(grid)
	width := len(grid[0])

	// map tiles to entered directions to track energized
	energized := map[coord][]rune{}

	beams := []beam{{coord{0, 0}, 'e'}}

	for len(beams) > 0 {
		// pop from remaining beams
		cur := beams[0]
		beams = beams[1:]

		// if out of bounds, skip
		if cur.coord.row < 0 || cur.coord.row >= length || cur.coord.col < 0 || cur.coord.col >= width {
			continue
		}

		// if alread entered from dir, skip
		if slices.Contains(energized[cur.coord], cur.enteredDir) {
			continue
		}
		// otherwise add to energized
		energized[cur.coord] = append(energized[cur.coord], cur.enteredDir)

		// find next tiles depending on current tile and entered dir
		tile := grid[cur.coord.row][cur.coord.col]
		for _, nextdir := range rm[tile][cur.enteredDir] {
			// add to queue
			beams = append(beams, beam{coord{cur.coord.row + dm[nextdir].row, cur.coord.col + dm[nextdir].col}, nextdir})
		}
	}
	return len(energized), nil
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)
	length := len(grid)
	width := len(grid[0])

	starts := []beam{}
	for r := range grid {
		starts = append(starts, beam{coord{r, 0}, 'e'})
		starts = append(starts, beam{coord{r, width - 1}, 'w'})
	}
	for c := range grid[0] {
		starts = append(starts, beam{coord{0, c}, 's'})
		starts = append(starts, beam{coord{length - 1, c}, 'n'})
	}

	largest := 0
	for _, start := range starts {
		beams := []beam{start}
		// map tiles to entered directions to track energized
		energized := map[coord][]rune{}

		for len(beams) > 0 {
			// pop from remaining beams
			cur := beams[0]
			beams = beams[1:]

			// if out of bounds, skip
			if cur.coord.row < 0 || cur.coord.row >= length || cur.coord.col < 0 || cur.coord.col >= width {
				continue
			}

			// if alread entered from dir, skip
			if slices.Contains(energized[cur.coord], cur.enteredDir) {
				continue
			}
			// otherwise add to energized
			energized[cur.coord] = append(energized[cur.coord], cur.enteredDir)

			// find next tiles depending on current tile and entered dir
			tile := grid[cur.coord.row][cur.coord.col]
			for _, nextdir := range rm[tile][cur.enteredDir] {
				// add to queue
				beams = append(beams, beam{coord{cur.coord.row + dm[nextdir].row, cur.coord.col + dm[nextdir].col}, nextdir})
			}
		}
		if len(energized) > largest {
			largest = len(energized)
		}
	}

	return largest, nil
}
