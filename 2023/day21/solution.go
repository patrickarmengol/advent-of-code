package day21

import (
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
	"github.com/patrickarmengol/megs/deque/lldeque"
	"github.com/patrickarmengol/megs/set/hashset"
)

type vertex struct {
	row int
	col int
}

func findStart(g [][]rune) vertex {
	for r := range g {
		for c := range g[0] {
			if g[r][c] == 'S' {
				return vertex{r, c}
			}
		}
	}
	panic("couldn't find start")
}

type state struct {
	v   vertex
	nsl int
}

func getNeighbors(g [][]rune, v vertex, inf bool) []vertex {
	adjacencies := map[rune]vertex{
		'N': {-1, 0},
		'S': {1, 0},
		'W': {0, -1},
		'E': {0, 1},
	}

	neighbors := []vertex{}
	for _, adj := range adjacencies {
		n := vertex{v.row + adj.row, v.col + adj.col}
		if inf {
			nw := vertex{modulo(n.row, len(g)), modulo(n.col, len(g[0]))}
			if g[nw.row][nw.col] == '#' {
				continue
			}
		} else {
			if n.row < 0 || n.row >= len(g) || n.col < 0 || n.col >= len(g[0]) || g[n.row][n.col] == '#' {
				continue
			}
		}
		neighbors = append(neighbors, n)
	}

	return neighbors
}

func modulo(a, b int) int {
	return ((a % b) + b) % b
}

func Part1(input string) (int, error) {
	grid := util.Gridify(input)

	numSteps := 64

	start := state{findStart(grid), numSteps}
	reached := hashset.New[vertex]()
	ends := hashset.New[vertex]()

	frontier := lldeque.Of(start)

	for !frontier.IsEmpty() {
		cs := frontier.PopFront()
		if cs.nsl%2 == 0 {
			ends.Add(cs.v)
			if cs.nsl == 0 {
				continue
			}
		}
		for _, n := range getNeighbors(grid, cs.v, false) {
			if !reached.Has(n) {
				frontier.PushBack(state{n, cs.nsl - 1})
				reached.Add(n)
			}
		}
	}

	// fmt.Println(ends)

	// for _, ee := range ends.Members() {
	// 	grid[ee.row][ee.col] = 'O'
	// }
	//
	// util.PrintGrid(grid)

	return ends.Len(), nil
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)

	// numSteps is 65 + 131*x = 26501365
	// x=202300

	// find the coefficients using polynomial regression
	// need 3 points for degree 2 polynomial

	p := []int{}
	for i := 0; i < 3; i++ {
		numSteps := 65 + (131 * i)

		start := state{findStart(grid), numSteps}
		reached := hashset.New[vertex]()
		ends := hashset.New[vertex]()

		frontier := lldeque.Of(start)

		for !frontier.IsEmpty() {
			cs := frontier.PopFront()
			if cs.nsl%2 == 0 {
				ends.Add(cs.v)
			}
			if cs.nsl == 0 {
				continue
			}
			for _, n := range getNeighbors(grid, cs.v, true) {
				if !reached.Has(n) {
					frontier.PushBack(state{n, cs.nsl - 1})
					reached.Add(n)
				}
			}
		}
		p = append(p, ends.Len())
	}

	// with points for x ranging 0..2
	//
	// p[0] = a(0)^2 + b(0) + c = c
	// p[1] = a(1)^2 + b(1) + c = a + b + c -> a + b = p[1] - p[0]
	// p[2] = a(2)^2 + b(2) + c = 4a + 2b + c -> 4a + 2b = p[2] - p[0]
	//
	// (4a + 2b) - (2a + 2b) = (p[2] - p[0]) - (2p[1] - 2p[0])
	// 2a = p[2] - 2p[1] + p[0]
	//
	// a = (p[2] - 2p[1] + p[0])/2
	// b = p[1] - p[0] - a
	// c = p[0]

	a := (p[2] - (2 * p[1]) + p[0]) / 2
	b := p[1] - p[0] - a
	c := p[0]

	// quadratic fit is 15024 x^2 + 15114 x + 3814

	x := 202300
	res := (a * x * x) + (b * x) + c

	return res, nil
}
