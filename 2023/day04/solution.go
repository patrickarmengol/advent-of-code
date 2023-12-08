package day04

import (
	
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

func Part1(input string) (int, error) {
	cards := parse.Lines(input)
	total := 0
	for _, card := range cards {
		numbers := strings.Split(strings.Split(card, ": ")[1], " | ")
		winningNums := util.NewSet(strings.Fields(numbers[0])...)
		haveNums := util.NewSet(strings.Fields(numbers[1])...)
		nMatches := len(winningNums.Intersection(haveNums))
		if nMatches >= 1 {
			total += 1 << (nMatches - 1) // equiv 2^(nMatches - 1)
		}
	}
	return total, nil
}

func Part2(input string) (int, error) {
	cards := parse.Lines(input)

	cardCopyCount := map[int]int{}

	total := 0
	for i, card := range cards {
		cardIndex := i + 1

		numbers := strings.Split(strings.Split(card, ": ")[1], " | ")
		winningNums := util.NewSet(strings.Fields(numbers[0])...)
		haveNums := util.NewSet(strings.Fields(numbers[1])...)
		nMatches := len(winningNums.Intersection(haveNums))

		for j := cardIndex + 1; nMatches > 0; j++ {
			cardCopyCount[j] += (1 + cardCopyCount[cardIndex])
			nMatches--
		}

		total += (1 + cardCopyCount[cardIndex])
	}

	return total, nil
}
