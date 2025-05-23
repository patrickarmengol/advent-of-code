package day17

import (
	"container/heap"
	"math"
	"slices"
	"strconv"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

type vertex struct {
	row int
	col int
}

type state struct {
	vert     vertex
	curDir   string
	dirCount int
}

type result struct {
	dist int
	st   state
}

type resultheap []result

func (h resultheap) Len() int           { return len(h) }
func (h resultheap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h resultheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *resultheap) Push(x any) {
	*h = append(*h, x.(result))
}

func (h *resultheap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var adjacencies = map[string]vertex{
	"n": {-1, 0},
	"s": {1, 0},
	"w": {0, -1},
	"e": {0, 1},
}

func dijkstra(g [][]int, start vertex, streakMin int, streakMax int) map[state]int {
	// initialize minheap
	resheap := &resultheap{result{0, state{start, "", 0}}}
	heap.Init(resheap)

	// init a map to track distances
	distances := map[state]int{}

	// while items in the heap
	for resheap.Len() > 0 {
		// pop from heap
		res := heap.Pop(resheap).(result)

		// add valid neighbors
		for newDir, adj := range adjacencies {
			neighbor := vertex{res.st.vert.row + adj.row, res.st.vert.col + adj.col}

			// check not opposite direction
			if slices.Contains([]string{"ns", "sn", "we", "ew"}, res.st.curDir+newDir) {
				continue
			}

			// check in bounds
			if neighbor.row < 0 || neighbor.row >= len(g) || neighbor.col < 0 || neighbor.col >= len(g[0]) {
				continue
			}

			// check new direction streak count
			var newCount int
			if newDir == res.st.curDir {
				newCount = res.st.dirCount + 1
			} else {
				newCount = 1
			}
			// gte min before changing directions (except start)
			if res.st.dirCount < streakMin && newCount == 1 && res.st.dirCount != 0 {
				continue
			}
			// gte min at end block
			if res.st.dirCount < streakMin && (neighbor == vertex{len(g) - 1, len(g[0]) - 1}) {
				continue
			}
			// lte max
			if newCount > streakMax {
				continue
			}

			// neighbor is valid, delare new state
			newDist := res.dist + g[neighbor.row][neighbor.col]
			newSt := state{neighbor, newDir, newCount}

			// check distance not already logged for same state
			if _, ok := distances[newSt]; ok {
				continue
			}

			// update distances and push to heap
			distances[newSt] = newDist
			heap.Push(resheap, result{newDist, newSt})
		}
	}
	return distances
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	// populate weights graph
	weights := util.Make2DGrid[int](len(lines), len(lines[0]))
	for r := range lines {
		for c := range lines[0] {
			w, err := strconv.Atoi(string(lines[r][c]))
			if err != nil {
				panic("couldn't convert num in input to int")
			}
			weights[r][c] = w
		}
	}

	// calc min dist found at end node
	end := vertex{len(weights) - 1, len(weights[0]) - 1}
	m := math.MaxInt
	for state, dist := range dijkstra(weights, vertex{0, 0}, 0, 3) {
		if state.vert == end {
			m = min(m, dist)
		}
	}

	return m, nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	// populate weights graph
	weights := util.Make2DGrid[int](len(lines), len(lines[0]))
	for r := range lines {
		for c := range lines[0] {
			w, err := strconv.Atoi(string(lines[r][c]))
			if err != nil {
				panic("couldn't convert num in input to int")
			}
			weights[r][c] = w
		}
	}

	// calc min dist found at end node
	end := vertex{len(weights) - 1, len(weights[0]) - 1}
	m := math.MaxInt
	for state, dist := range dijkstra(weights, vertex{0, 0}, 4, 10) {
		if state.vert == end {
			m = min(m, dist)
		}
	}

	return m, nil
}
