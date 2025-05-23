package day01

import (
	
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	total := 0
	for _, line := range lines {
		left := strings.IndexAny(line, "0123456789")
		right := strings.LastIndexAny(line, "0123456789")
		calib, err := strconv.Atoi(fmt.Sprintf("%c%c", line[left], line[right]))
		if err != nil {
			return 0, err
		}
		total += calib
	}

	return total, nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	rx := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)

	total := 0
	for _, line := range lines {
		nums := []string{}
		for j := 0; j <= len(line); j++ {
			n := rx.FindString(line[j:])
			if n == "" {
				break
			}
			nums = append(nums, n)
		}

		for i, n := range nums {
			switch n {
			case "one":
				nums[i] = "1"
			case "two":
				nums[i] = "2"
			case "three":
				nums[i] = "3"
			case "four":
				nums[i] = "4"
			case "five":
				nums[i] = "5"
			case "six":
				nums[i] = "6"
			case "seven":
				nums[i] = "7"
			case "eight":
				nums[i] = "8"
			case "nine":
				nums[i] = "9"
			default:
			}
		}

		left := nums[0]
		right := nums[len(nums)-1]
		calib, err := strconv.Atoi(fmt.Sprintf("%s%s", left, right))
		if err != nil {
			return 0, err
		}

		total += calib
	}

	return total, nil
}
