package day06

import (
	
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

const (
	highCard     = iota + 1 // all unique
	onePair                 // {2, 1, 1, 1}
	twoPair                 // {2, 2, 1}
	threeOfAKind            // {3, 1, 1}
	fullHouse               // {3, 2}
	fourOfAKind             // {4, 1}
	fiveOfAKind             // {5}
)

type hand struct {
	cards string
	bid   int
	typ   int
}

func parseHands(s string, jr bool) ([]hand, error) {
	lines := parse.Lines(s)

	hands := []hand{}
	for _, line := range lines {
		fs := strings.Fields(line)
		cs := fs[0]
		b, err := strconv.Atoi(fs[1])
		if err != nil {
			return nil, err
		}
		t := findTyp(cs, jr)
		hands = append(hands, hand{cs, b, t})
	}
	return hands, nil
}

func findTyp(cards string, jokerRule bool) int {
	counts := map[rune]int{}

	for _, card := range cards {
		counts[card] += 1
	}
	var js int
	if jokerRule {
		// can i always just add to largest?
		js = counts['J']
		delete(counts, 'J')
	}

	signatureSlice := []int{}
	for _, count := range counts {
		signatureSlice = append(signatureSlice, count)
	}
	slices.Sort(signatureSlice)
	if jokerRule {
		if len(signatureSlice) == 0 {
			signatureSlice = append(signatureSlice, 0)
		}
		signatureSlice[len(signatureSlice)-1] += js
	}

	signature := fmt.Sprint(signatureSlice)

	switch signature {
	case "[1 1 1 1 1]":
		return highCard
	case "[1 1 1 2]":
		return onePair
	case "[1 2 2]":
		return twoPair
	case "[1 1 3]":
		return threeOfAKind
	case "[2 3]":
		return fullHouse
	case "[1 4]":
		return fourOfAKind
	case "[5]":
		return fiveOfAKind
	default:
		panic("impossible configuration of cards")
	}
}

func compareHands(jr bool) func(a, b hand) int {
	return func(a, b hand) int {
		var order string
		if jr {
			order = "J23456789TQKA"
		} else {
			order = "23456789TJQKA"
		}
		if a.typ < b.typ {
			return -1
		} else if a.typ > b.typ {
			return 1
		} else {
			for i := 0; i < 5; i++ {
				as := strings.IndexByte(order, a.cards[i])
				bs := strings.IndexByte(order, b.cards[i])
				if as < bs {
					return -1
				} else if as > bs {
					return 1
				}
			}
			return 0
		}
	}
}

func Part1(input string) (int, error) {
	hands, err := parseHands(input, false)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, compareHands(false))

	total := 0
	for i, h := range hands {
		rank := i + 1
		total += h.bid * rank
	}

	return total, nil
}

func Part2(input string) (int, error) {
	hands, err := parseHands(input, true)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(hands, compareHands(true))

	total := 0
	for i, h := range hands {
		rank := i + 1
		total += h.bid * rank
	}

	return total, nil
}
