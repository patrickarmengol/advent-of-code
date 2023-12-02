package day02

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

func Part1(input string) (string, error) {
	games := parse.Lines(input)

	colorMaxes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	total := 0
gameloop:
	for i, game := range games {
		draws, err := getAllDraws(game)
		if err != nil {
			return "", err
		}
		for _, draw := range draws {
			cMax, ok := colorMaxes[draw.color]
			if !ok {
				return "", errors.New(fmt.Sprintf("draw color %v is not found in color max map", draw.color))
			}
			if draw.num > cMax {
				continue gameloop
			}
		}
		total += i + 1
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	games := parse.Lines(input)

	total := 0
	for _, game := range games {
		colorMins := map[string]int{}
		draws, _ := getAllDraws(game)
		for _, draw := range draws {
			if draw.num > colorMins[draw.color] {
				colorMins[draw.color] = draw.num
			}
		}
		total += colorMins["red"] * colorMins["green"] * colorMins["blue"]
	}
	return strconv.Itoa(total), nil
}

type draw struct {
	num   int
	color string
}

func getAllDraws(game string) ([]draw, error) {
	draws := []draw{}
	sets := strings.Split(strings.Split(game, ": ")[1], "; ")
	for _, set := range sets {
		ds := strings.Split(set, ", ")
		for _, d := range ds {
			drawParts := strings.Split(d, " ")
			num, err := strconv.Atoi(drawParts[0])
			if err != nil {
				return nil, errors.New(fmt.Sprintln("problem parsing num for draw:", d))
			}
			color := drawParts[1]
			if !slices.Contains([]string{"red", "green", "blue"}, color) {
				return nil, errors.New(fmt.Sprintln("problem parsing color foor draw:", d))
			}
			draws = append(draws, draw{num: num, color: color})
		}
	}
	return draws, nil
}
