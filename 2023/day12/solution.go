package day12

import (
	"fmt"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

var dpWays = map[string]int{}

func findWays(springs []rune, groups []int, si, gi, clen int) int {
	// dp cache key is serialized state
	key := fmt.Sprintf("%q %v %d %d %d", springs, groups, si, gi, clen)

	// check dp for existing
	if result, ok := dpWays[key]; ok {
		// fmt.Printf("found for %s -> %d\n", key, result)
		return result
	}

	// if reached end
	if si == len(springs) {
		if gi == len(groups) && clen == 0 {
			// finished with groups and current empty
			return 1
		} else if gi == len(groups)-1 && clen == groups[gi] {
			// on last group and current fills group
			return 1
		} else {
			// groups not satisfied
			return 0
		}
	}

	// backtrack recursion with next character
	ways := 0
	nextChar := springs[si]
	// step '.'
	if nextChar == '.' || nextChar == '?' {
		if clen == 0 {
			ways += findWays(springs, groups, si+1, gi, 0)
		} else if clen > 0 && gi < len(groups) && groups[gi] == clen {
			ways += findWays(springs, groups, si+1, gi+1, 0)
		}
	}
	// step '#'
	if nextChar == '#' || nextChar == '?' {
		ways += findWays(springs, groups, si+1, gi, clen+1)
	}

	// update dp cache
	dpWays[key] = ways

	return ways
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	total := 0
	for _, line := range lines {
		fs := strings.Fields(line)
		springs := []rune(fs[0])
		groups, err := util.AtoiSlice(strings.Split(fs[1], ","))
		if err != nil {
			return 0, err
		}
		total += findWays(springs, groups, 0, 0, 0)
	}
	return total, nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	total := 0
	for _, line := range lines {
		fs := strings.Fields(line)

		springs := []rune(fs[0])
		eSprings := []rune{}
		for i := 0; i < 4; i++ {
			eSprings = append(eSprings, springs...)
			eSprings = append(eSprings, '?')
		}
		eSprings = append(eSprings, springs...)

		groups, err := util.AtoiSlice(strings.Split(fs[1], ","))
		if err != nil {
			return 0, err
		}
		eGroups := []int{}
		for i := 0; i < 5; i++ {
			eGroups = append(eGroups, groups...)
		}

		total += findWays(eSprings, eGroups, 0, 0, 0)
	}
	return total, nil
}
