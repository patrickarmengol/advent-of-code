package day18

import (
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

type vertex struct {
	row int
	col int
}

type instruction struct {
	dir   rune
	steps int
}

func intabs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func digBorder(instrs []instruction) ([]vertex, int) {
	cur := vertex{0, 0}
	vs := []vertex{}

	adjacencies := map[rune]vertex{
		'U': {-1, 0},
		'D': {1, 0},
		'L': {0, -1},
		'R': {0, 1},
	}

	totalTiles := 0
	for _, instr := range instrs {
		totalTiles += instr.steps
		cur = vertex{
			row: cur.row + adjacencies[instr.dir].row*instr.steps,
			col: cur.col + adjacencies[instr.dir].col*instr.steps,
		}
		vs = append(vs, cur)
	}
	return vs, totalTiles
}

func areaPolygon(vs []vertex) int {
	sum := 0
	for i := 0; i < len(vs)-1; i++ {
		sum += (vs[i].row * vs[i+1].col) - (vs[i].col * vs[i+1].row)
	}
	res := intabs(sum / 2)
	return res
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	instructions := []instruction{}
	for _, line := range lines {
		fs := strings.Fields(line)
		dir := rune(fs[0][0])
		steps, err := strconv.Atoi(fs[1])
		if err != nil {
			return 0, err
		}
		instructions = append(instructions, instruction{dir, steps})
	}

	vs, btiles := digBorder(instructions)

	area := areaPolygon(vs)

	inter := area + 1 - (btiles / 2)

	return inter + btiles, nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	instructions := []instruction{}
	for _, line := range lines {
		fs := strings.Fields(line)
		color := strings.Trim(fs[2], "()#")

		steps64, err := strconv.ParseInt(color[:len(color)-1], 16, 0)
		if err != nil {
			return 0, err
		}
		steps := int(steps64)

		var dir rune
		switch color[len(color)-1] {
		case '0':
			dir = 'R'
		case '1':
			dir = 'D'
		case '2':
			dir = 'L'
		case '3':
			dir = 'U'
		default:
			panic("invalid last char in color")
		}

		instructions = append(instructions, instruction{dir, steps})
	}

	vs, btiles := digBorder(instructions)

	area := areaPolygon(vs)

	inter := area + 1 - (btiles / 2)

	return inter + btiles, nil
}
