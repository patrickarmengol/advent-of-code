package day08

import (
	
	"slices"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

type node struct {
	left  string
	right string
}

func parseNetwork(lines []string) map[string]node {
	network := map[string]node{}
	for _, line := range lines {
		fs := strings.Split(line, " = ")
		index := fs[0]
		dsts := strings.Split(strings.Trim(fs[1], "()"), ", ")
		left := dsts[0]
		right := dsts[1]
		network[index] = node{left, right}
	}
	return network
}

func allEndWithZ(ns []string) bool {
	for _, n := range ns {
		if !strings.HasSuffix(n, "Z") {
			return false
		}
	}
	return true
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	steps := lines[0]
	network := parseNetwork(lines[2:])

	cur := "AAA"
	step := 0
	for cur != "ZZZ" {
		direction := steps[step%len(steps)]
		if direction == 'L' {
			cur = network[cur].left
		} else {
			cur = network[cur].right
		}
		step += 1
	}
	return step, nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmMulti(ns ...int) int {
	result := ns[0]
	for _, n := range ns[1:] {
		result = lcm(result, n)
	}
	return result
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	steps := lines[0]
	network := parseNetwork(lines[2:])

	curs := []string{}
	for k := range network {
		if strings.HasSuffix(k, "A") {
			curs = append(curs, k)
		}
	}

	loopLengths := []int{}
	step := 0
	for len(loopLengths) < len(curs) {
		direction := steps[step%len(steps)]
		for i, c := range curs {
			if direction == 'L' {
				curs[i] = network[c].left
			} else {
				curs[i] = network[c].right
			}
		}
		step += 1
		toDelete := []int{}
		for i, c := range curs {
			if strings.HasSuffix(c, "Z") {
				loopLengths = append(loopLengths, step)
				toDelete = append(toDelete, i)
			}
		}
		for _, d := range toDelete {
			curs = slices.Delete(curs, d, d)
		}
	}

	return lcmMulti(loopLengths...), nil
}
