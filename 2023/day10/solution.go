package day10

import (
	"regexp"
	"slices"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

type node struct {
	row int
	col int
}

var toCheckSides = map[rune][]string{
	'F': {"right", "down"},
	'-': {"left", "right"},
	'7': {"left", "down"},
	'|': {"up", "down"},
	'J': {"up", "left"},
	'L': {"up", "right"},
	'S': {"left", "right", "up", "down"},
}

var validNeighbors = map[string][]rune{
	"left":  {'-', 'F', 'L'},
	"right": {'-', '7', 'J'},
	"up":    {'|', 'F', '7'},
	"down":  {'|', 'L', 'J'},
}

func findNeighbors(grid [][]rune, vertex node) []node {
	// get bounds of grid
	length := len(grid)
	width := len(grid[0])

	// get char at vertex
	char := grid[vertex.row][vertex.col]

	// define potential adjacentCells
	adjacentCells := map[string]node{
		"left":  {vertex.row, vertex.col - 1},
		"right": {vertex.row, vertex.col + 1},
		"up":    {vertex.row - 1, vertex.col},
		"down":  {vertex.row + 1, vertex.col},
	}

	result := []node{}

	validSides := []string{}
	for _, side := range toCheckSides[char] {
		a := adjacentCells[side]
		if a.row >= 0 && a.row < length && a.col >= 0 && a.col < width {
			if slices.Contains(validNeighbors[side], grid[a.row][a.col]) {
				validSides = append(validSides, side)
				result = append(result, a)
			}
		}
	}
	if char == 'S' {
		slices.Sort(validSides)
		sig := strings.Join(validSides, "")
		sm := map[string]rune{
			"downleft":  '7',
			"downright": 'F',
			"downup":    '|',
			"leftright": '-',
			"leftup":    'J',
			"rightup":   'L',
		}
		grid[vertex.row][vertex.col] = sm[sig]
	}

	return result
}

func findStart(grid [][]rune) node {
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'S' {
				return node{r, c}
			}
		}
	}
	panic("couldn't find S")
}

func Part1(input string) (int, error) {
	// initialize grid
	grid := util.Gridify(input)

	// initialize distance map
	dist := map[node]int{}
	// initialize visited map
	visited := map[node]bool{}

	// find start
	s := findStart(grid)
	visited[s] = true

	// initialize queue and enqueue start
	q := []node{s}

	d := 0

	// while q is not empty
	for len(q) != 0 {
		// dequeue vertex from queue
		v := q[0]
		q = q[1:]

		if dist[v] == d+1 {
			d++
		}

		// fmt.Printf("from vertex %v %c\n", v, grid[v.row][v.col])
		// iterate through neighbors
		for _, n := range findNeighbors(grid, v) {
			if !visited[n] {
				// fmt.Printf("visiting %v %c %d\n", n, grid[n.row][n.col], d+1)
				visited[n] = true
				dist[n] = d + 1
				q = append(q, n)
			}
		}
	}

	// fmt.Println(dist)
	md := 0
	for _, v := range dist {
		if v > md {
			md = v
		}
	}
	return md, nil
}

var asdf = map[*regexp.Regexp]string{
	regexp.MustCompile(`L-*7`): "|",
	regexp.MustCompile(`L-*J`): "||",
	regexp.MustCompile(`F-*7`): "||",
	regexp.MustCompile(`F-*J`): "|",
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)

	mainloop := map[node]bool{}

	s := findStart(grid)
	mainloop[s] = true

	q := []node{s}

	for len(q) != 0 {
		v := q[0]
		q = q[1:]

		for _, n := range findNeighbors(grid, v) {
			if !mainloop[n] {
				mainloop[n] = true
				q = append(q, n)
				// fmt.Println("visiting", n)
			}
		}
	}

	// replace non-mainloop pipes with empty cells
	for r := range grid {
		for c := range grid[r] {
			if !mainloop[node{r, c}] {
				grid[r][c] = '.'
			}
		}
	}

	fill := 0
	for _, row := range grid {
		srow := string(row)
		for rx, rep := range asdf {
			srow = rx.ReplaceAllString(srow, rep)
		}
		cross := false
		for _, char := range srow {
			if strings.ContainsRune("L7FJ|", char) {
				cross = !cross
			} else if char == '.' && cross {
				fill += 1
			}
		}
	}

	// // replace non-mainloop pipes with empty cells
	// for r := range grid {
	// 	for c := range grid[r] {
	// 		fmt.Printf("%c", grid[r][c])
	// 	}
	// 	fmt.Println()
	// }
	//
	return fill, nil
}
