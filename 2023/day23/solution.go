package day23

import (
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
	"github.com/patrickarmengol/megs/deque/lldeque"
	"github.com/patrickarmengol/megs/set/hashset"
)

type coord struct {
	row int
	col int
}

var adjacencies = map[rune]coord{
	'N': {-1, 0},
	'S': {1, 0},
	'W': {0, -1},
	'E': {0, 1},
}

func getNeighborsFromPath(grid [][]rune, path []coord) []coord {
	// current position is at end of path
	curCoord := path[len(path)-1]
	curTile := grid[curCoord.row][curCoord.col]

	// find neighbors from current
	nextCoords := []coord{}
	// next
	switch curTile {
	case '>':
		n := coord{curCoord.row + adjacencies['E'].row, curCoord.col + adjacencies['E'].col}
		if n.row < 0 || n.row > len(grid)-1 || n.col < 0 || n.col > len(grid[0])-1 {
			return nextCoords
		}
		nextCoords = append(nextCoords, n)
	case '<':
		n := coord{curCoord.row + adjacencies['W'].row, curCoord.col + adjacencies['W'].col}
		if n.row < 0 || n.row > len(grid)-1 || n.col < 0 || n.col > len(grid[0])-1 {
			return nextCoords
		}
		nextCoords = append(nextCoords, n)
	case 'v':
		n := coord{curCoord.row + adjacencies['S'].row, curCoord.col + adjacencies['S'].col}
		if n.row < 0 || n.row > len(grid)-1 || n.col < 0 || n.col > len(grid[0])-1 {
			return nextCoords
		}
		nextCoords = append(nextCoords, n)
	case '^':
		n := coord{curCoord.row + adjacencies['N'].row, curCoord.col + adjacencies['N'].col}
		if n.row < 0 || n.row > len(grid)-1 || n.col < 0 || n.col > len(grid[0])-1 {
			return nextCoords
		}
		nextCoords = append(nextCoords, n)
	case '.', 'S':
		for d, adj := range adjacencies {
			n := coord{curCoord.row + adj.row, curCoord.col + adj.col}
			if n.row < 0 || n.row > len(grid)-1 || n.col < 0 || n.col > len(grid[0])-1 {
				continue
			}
			nTile := grid[n.row][n.col]
			// don't walk into trees
			if nTile == '#' {
				continue
			}
			// don't backtrack
			if len(path) > 1 {
				lastCoord := path[len(path)-2]
				if lastCoord == n {
					continue
				}
			}
			// don't go uphill
			if nTile == '<' && d == 'E' {
				continue
			}
			if nTile == '>' && d == 'W' {
				continue
			}
			if nTile == 'v' && d == 'N' {
				continue
			}
			if nTile == '^' && d == 'S' {
				continue
			}
			nextCoords = append(nextCoords, n)
		}
	}

	return nextCoords
}

func Part1(input string) (int, error) {
	grid := util.Gridify(input)
	// assuming a hardcoded start end by inspection of input
	// also can see it seems to be a DAG
	// each node is preceded by arrow(s) and has arrow(s) coming out
	// the number of '.' between nodes can be used as weight
	start := coord{0, 1}
	end := coord{len(grid) - 1, len(grid[0]) - 2}
	grid[start.row][start.col] = 'S'
	grid[end.row][end.col] = 'E'

	coordHashFunc := func(c coord) string { return fmt.Sprintf("<%d,%d>", c.row, c.col) }
	g := graph.New(coordHashFunc, graph.Directed(), graph.Acyclic(), graph.Weighted())

	g.AddVertex(start)
	g.AddVertex(end)

	// queue holds paths
	q := lldeque.New[[]coord]()
	q.PushBack([]coord{start})

	for q.Len() != 0 {
		// fmt.Println(q.String())
		curPath := q.PopFront()
		curCoord := curPath[len(curPath)-1]
		curTile := grid[curCoord.row][curCoord.col]

		if curTile == 'E' {
			// last node at start of path
			lastNode := curPath[0]
			// new node at end of path
			newNode := curCoord
			// weight is len of path - 1
			weight := len(curPath) - 1
			// vertex already exists in graph
			// add edge
			g.AddEdge(coordHashFunc(lastNode), coordHashFunc(newNode), graph.EdgeWeight(weight), graph.EdgeAttribute("label", fmt.Sprintf("%d", weight)))
			continue
		}

		nextCoords := getNeighborsFromPath(grid, curPath)
		// fmt.Println("next", nextCoords)

		// check if new node

		curIsDot := curTile == '.'
		prevIsArrow := false
		if len(curPath) > 1 {
			prevCoord := curPath[len(curPath)-2]
			prevTile := grid[prevCoord.row][prevCoord.col]
			prevIsArrow = strings.ContainsRune("><v^", prevTile)
		}
		nextIsArrow := true
		for _, n := range nextCoords {
			nTile := grid[n.row][n.col]
			if !strings.ContainsRune("><v^", nTile) {
				nextIsArrow = false
				break
			}
		}
		if curIsDot && prevIsArrow && nextIsArrow {
			// fmt.Println("found new node", curCoord)
			// found new node
			// last node at start of path
			lastNode := curPath[0]
			// new node at end of path
			newNode := curCoord
			// weight is len of path - 1
			weight := len(curPath) - 1
			// add node
			g.AddVertex(newNode)
			// add edge
			g.AddEdge(coordHashFunc(lastNode), coordHashFunc(newNode), graph.EdgeWeight(weight), graph.EdgeAttribute("label", fmt.Sprintf("%d", weight)))
			// reset path
			curPath = []coord{newNode}
		}

		// push neighbors to queue
		for _, n := range nextCoords {
			nextPath := []coord{}
			nextPath = append(nextPath, curPath...)
			nextPath = append(nextPath, n)
			q.PushBack(nextPath)
		}
	}

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)

	order, _ := graph.TopologicalSort(g)
	fmt.Println(order)

	dist := map[string]int{}
	dist[order[0]] = 0

	am, err := g.AdjacencyMap()
	if err != nil {
		panic("asdf")
	}

	for _, node := range order {
		for successor, v := range am[node] {
			weight := v.Properties.Weight
			dist[successor] = max(dist[successor], dist[node]+weight)
		}
	}

	fmt.Println(dist)

	// util.PrintGrid(grid)

	return dist[coordHashFunc(end)], nil
}

func Part2(input string) (int, error) {
	grid := util.Gridify(input)
	// assuming a hardcoded start end by inspection of input
	// also can see it seems to be a DAG
	// each node is preceded by arrow(s) and has arrow(s) coming out
	// the number of '.' between nodes can be used as weight
	start := coord{0, 1}
	end := coord{len(grid) - 1, len(grid[0]) - 2}
	grid[start.row][start.col] = 'S'
	grid[end.row][end.col] = 'E'

	coordHashFunc := func(c coord) string { return fmt.Sprintf("<%d,%d>", c.row, c.col) }
	g := graph.New(coordHashFunc, graph.Weighted())

	g.AddVertex(start)
	g.AddVertex(end)

	// queue holds paths
	q := lldeque.New[[]coord]()
	q.PushBack([]coord{start})

	for q.Len() != 0 {
		// fmt.Println(q.String())
		curPath := q.PopFront()
		curCoord := curPath[len(curPath)-1]
		curTile := grid[curCoord.row][curCoord.col]

		if curTile == 'E' {
			// last node at start of path
			lastNode := curPath[0]
			// new node at end of path
			newNode := curCoord
			// weight is len of path - 1
			weight := len(curPath) - 1
			// vertex already exists in graph
			// add edge
			g.AddEdge(coordHashFunc(lastNode), coordHashFunc(newNode), graph.EdgeWeight(weight), graph.EdgeAttribute("label", fmt.Sprintf("%d", weight)))
			continue
		}

		nextCoords := getNeighborsFromPath(grid, curPath)
		// fmt.Println("next", nextCoords)

		// check if new node

		curIsDot := curTile == '.'
		prevIsArrow := false
		if len(curPath) > 1 {
			prevCoord := curPath[len(curPath)-2]
			prevTile := grid[prevCoord.row][prevCoord.col]
			prevIsArrow = strings.ContainsRune("><v^", prevTile)
		}
		nextIsArrow := true
		for _, n := range nextCoords {
			nTile := grid[n.row][n.col]
			if !strings.ContainsRune("><v^", nTile) {
				nextIsArrow = false
				break
			}
		}
		if curIsDot && prevIsArrow && nextIsArrow {
			// fmt.Println("found new node", curCoord)
			// found new node
			// last node at start of path
			lastNode := curPath[0]
			// new node at end of path
			newNode := curCoord
			// weight is len of path - 1
			weight := len(curPath) - 1
			// add node
			g.AddVertex(newNode)
			// add edge
			g.AddEdge(coordHashFunc(lastNode), coordHashFunc(newNode), graph.EdgeWeight(weight), graph.EdgeAttribute("label", fmt.Sprintf("%d", weight)))
			// reset path
			curPath = []coord{newNode}
		}

		// push neighbors to queue
		for _, n := range nextCoords {
			nextPath := []coord{}
			nextPath = append(nextPath, curPath...)
			nextPath = append(nextPath, n)
			q.PushBack(nextPath)
		}
	}

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)

	dist := map[string]int{}

	am, err := g.AdjacencyMap()
	if err != nil {
		panic("asdf")
	}

	getLongestPath(am, coordHashFunc(start), 0, hashset.New[string](), dist)

	fmt.Println(dist)

	// util.PrintGrid(grid)

	return dist[coordHashFunc(end)], nil
}

func getLongestPath(am map[string]map[string]graph.Edge[string], node string, curSum int, visited *hashset.Set[string], distances map[string]int) {
	if visited.Has(node) {
		return
	}
	visited.Add(node)

	if distances[node] < curSum {
		distances[node] = curSum
	}

	for successor, v := range am[node] {
		weight := v.Properties.Weight
		getLongestPath(am, successor, curSum+weight, visited, distances)
	}

	visited.Remove(node)
}
