package day02

import (
	"regexp"
	"slices"
	"strconv"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

// leaving old solution commented out here
//
// func Part1(input string) (string, error) {
// 	games := parse.Lines(input)
//
// 	colorMaxes := map[string]int{
// 		"red":   12,
// 		"green": 13,
// 		"blue":  14,
// 	}
//
// 	total := 0
// gameloop:
// 	for i, game := range games {
// 		draws, err := getAllDraws(game)
// 		if err != nil {
// 			return "", err
// 		}
// 		for _, draw := range draws {
// 			cMax, ok := colorMaxes[draw.color]
// 			if !ok {
// 				return "", errors.New(fmt.Sprintf("draw color %v is not found in color max map", draw.color))
// 			}
// 			if draw.num > cMax {
// 				continue gameloop
// 			}
// 		}
// 		total += i + 1
// 	}
// 	return strconv.Itoa(total), nil
// }
//
// func Part2(input string) (string, error) {
// 	games := parse.Lines(input)
//
// 	total := 0
// 	for _, game := range games {
// 		colorMins := map[string]int{}
// 		draws, _ := getAllDraws(game)
// 		for _, draw := range draws {
// 			if draw.num > colorMins[draw.color] {
// 				colorMins[draw.color] = draw.num
// 			}
// 		}
// 		total += colorMins["red"] * colorMins["green"] * colorMins["blue"]
// 	}
// 	return strconv.Itoa(total), nil
// }
//
// type draw struct {
// 	num   int
// 	color string
// }
//
// func getAllDraws(game string) ([]draw, error) {
// 	draws := []draw{}
// 	sets := strings.Split(strings.Split(game, ": ")[1], "; ")
// 	for _, set := range sets {
// 		ds := strings.Split(set, ", ")
// 		for _, d := range ds {
// 			drawParts := strings.Split(d, " ")
// 			num, err := strconv.Atoi(drawParts[0])
// 			if err != nil {
// 				return nil, errors.New(fmt.Sprintln("problem parsing num for draw:", d))
// 			}
// 			color := drawParts[1]
// 			if !slices.Contains([]string{"red", "green", "blue"}, color) {
// 				return nil, errors.New(fmt.Sprintln("problem parsing color foor draw:", d))
// 			}
// 			draws = append(draws, draw{num: num, color: color})
// 		}
// 	}
// 	return draws, nil
// }

var (
	redRX   = regexp.MustCompile(`(\d+) red`)
	greenRX = regexp.MustCompile(`(\d+) green`)
	blueRX  = regexp.MustCompile(`(\d+) blue`)
)

func Part1(input string) (int, error) {
	total := 0
	for i, game := range parse.Lines(input) {
		maxR := slices.Max(numsFromMatches(redRX.FindAllStringSubmatch(game, -1)))
		maxG := slices.Max(numsFromMatches(greenRX.FindAllStringSubmatch(game, -1)))
		maxB := slices.Max(numsFromMatches(blueRX.FindAllStringSubmatch(game, -1)))
		if maxR > 12 || maxG > 13 || maxB > 14 {
			continue
		}
		total += i + 1
	}
	return total, nil
}

func Part2(input string) (int, error) {
	total := 0
	for _, game := range parse.Lines(input) {
		maxR := slices.Max(numsFromMatches(redRX.FindAllStringSubmatch(game, -1)))
		maxG := slices.Max(numsFromMatches(greenRX.FindAllStringSubmatch(game, -1)))
		maxB := slices.Max(numsFromMatches(blueRX.FindAllStringSubmatch(game, -1)))
		total += maxR * maxG * maxB
	}
	return total, nil
}

func numsFromMatches(matches [][]string) []int {
	nums := []int{}
	for _, m := range matches {
		num, err := strconv.Atoi(m[1]) // capture group 1
		if err != nil {
			panic("problem with a regex definition")
		}
		nums = append(nums, num)
	}
	return nums
}
