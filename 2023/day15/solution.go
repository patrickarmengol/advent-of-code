package day15

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var instructionRX = regexp.MustCompile(`(\w+)([=-])(\d+)?`)

func hash(s string) int {
	result := 0
	for _, char := range s {
		result += int(char)
		result *= 17
		result %= 256
	}
	return result
}

func Part1(input string) (int, error) {
	instructions := strings.Split(strings.TrimSpace(input), ",")

	total := 0
	for _, instruction := range instructions {
		total += hash(instruction)
	}

	return total, nil
}

type lens struct {
	label  string
	foclen int
}

func Part2(input string) (int, error) {
	boxes := map[int][]lens{}

	for _, ins := range instructionRX.FindAllStringSubmatch(input, -1) {
		if len(ins) != 4 {
			return 0, errors.New("problem parsing instructions")
		}
		label := ins[1]
		oper := ins[2]
		h := hash(label)
		if oper == "=" {
			foclen, err := strconv.Atoi(ins[3])
			if err != nil {
				return 0, errors.New("probelm converting instruction foclen to int")
			}
			replaced := false
			for i, l := range boxes[h] {
				if l.label == label {
					boxes[h][i] = lens{label, foclen}
					replaced = true
					break
				}
			}
			if !replaced {
				boxes[h] = append(boxes[h], lens{label, foclen})
			}
		} else if oper == "-" {
			for i, l := range boxes[h] {
				if l.label == label {
					boxes[h] = slices.Delete(boxes[h], i, i+1)
				}
			}
		}

	}

	total := 0
	for i, box := range boxes {
		for j, l := range box {
			total += (i + 1) * (j + 1) * l.foclen
		}
	}

	return total, nil
}
