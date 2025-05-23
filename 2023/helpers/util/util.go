package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

func Make2DGrid[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func Fill2DGrid(grid [][]rune, fill rune) {
	for r := range grid {
		for c := range grid[0] {
			grid[r][c] = fill
		}
	}
}

func Gridify(s string) [][]rune {
	lines := parse.Lines(s)

	grid := Make2DGrid[rune](len(lines), len(lines[0]))
	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
		}
	}

	return grid
}

func PrintGrid(grid [][]rune) {
	for r := range grid {
		for c := range grid[0] {
			fmt.Printf("%c", grid[r][c])
		}
		fmt.Println()
	}
}

func SprintGrid(grid [][]rune) string {
	chars := []rune{}
	for r := range grid {
		for c := range grid[0] {
			chars = append(chars, grid[r][c])
		}
		chars = append(chars, '\n')
	}
	return string(chars)
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := Set[T]{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}

func (s Set[T]) Contains(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Members() []T {
	result := make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Members())
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	result := NewSet(s.Members()...)
	result.Add(s2.Members()...)
	return result
}

func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	result := NewSet[T]()
	for _, v := range s.Members() {
		if s2.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func AtoiSlice(vals []string) ([]int, error) {
	result := make([]int, 0, len(vals))
	for _, va := range vals {
		va = strings.TrimSpace(va)
		vi, err := strconv.Atoi(va)
		if err != nil {
			return nil, err
		}
		result = append(result, vi)
	}
	return result, nil
}
