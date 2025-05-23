package day09

import (
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

func allZeros(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func diffs(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		result = append(result, diff)
	}
	return result
}

func history(nums []int) [][]int {
	history := [][]int{}
	for !allZeros(nums) {
		history = append(history, nums)
		nums = diffs(nums)
	}
	return history
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	total := 0
	for _, line := range lines {
		nums, err := util.AtoiSlice(strings.Fields(line))
		if err != nil {
			return 0, err
		}
		hist := history(nums)
		result := 0
		for i := len(hist) - 1; i >= 0; i-- {
			result += hist[i][len(hist[i])-1]
		}
		total += result
	}
	return total, nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	total := 0
	for _, line := range lines {
		nums, err := util.AtoiSlice(strings.Fields(line))
		if err != nil {
			return 0, err
		}
		hist := history(nums)
		result := 0
		for i := len(hist) - 1; i >= 0; i-- {
			result = hist[i][0] - result
		}
		total += result
	}
	return total, nil
}
