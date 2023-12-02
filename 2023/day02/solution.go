package day02

import (
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

func Part1(input string) (string, error) {
	games, err := parse.GetLines(input)
	if err != nil {
		return "", err
	}

	var (
		maxR = 12
		maxG = 13
		maxB = 14
	)

	total := 0
gameloop:
	for i, game := range games {
		draws := getAllDraws(game)
		for _, draw := range draws {
			switch draw.color {
			case "red":
				if draw.num > maxR {
					continue gameloop
				}
			case "green":
				if draw.num > maxG {
					continue gameloop
				}
			case "blue":
				if draw.num > maxB {
					continue gameloop
				}
			default:
			}
		}
		total += i + 1
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	games, err := parse.GetLines(input)
	if err != nil {
		return "", err
	}

	total := 0
	for _, game := range games {
		var minR, minG, minB int
		draws := getAllDraws(game)
		for _, draw := range draws {
			switch draw.color {
			case "red":
				if draw.num > minR {
					minR = draw.num
				}
			case "green":
				if draw.num > minG {
					minG = draw.num
				}
			case "blue":
				if draw.num > minB {
					minB = draw.num
				}
			}
		}
		total += minR * minG * minB
	}
	return strconv.Itoa(total), nil
}

type draw struct {
	num   int
	color string
}

func getAllDraws(game string) []draw {
	draws := []draw{}
	sets := strings.Split(strings.Split(game, ": ")[1], "; ")
	for _, set := range sets {
		ds := strings.Split(set, ", ")
		for _, d := range ds {
			drawParts := strings.Split(d, " ")
			num, err := strconv.Atoi(drawParts[0])
			if err != nil {
				panic("invalid parsing of num in draw pair")
			}
			color := drawParts[1]
			draws = append(draws, draw{num: num, color: color})
		}
	}
	return draws
}
