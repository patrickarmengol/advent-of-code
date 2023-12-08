package day06

import (
	
	"math"
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

// seconds_held * (seconds_allowed - seconds_held) = distance
// -seconds_held^2 + seconds_allowed*seconds_held - distance = 0
// a = -1, b = seconds_allowed, c = -distance
// use quadratic formula to find the range in which seconds_held produces a better result than record
// a is known, b is known, c is set to record

func Part1(input string) (int, error) {
	lines := parse.Lines(input)
	times, err := util.AtoiSlice(strings.Fields(strings.Split(lines[0], ": ")[1]))
	if err != nil {
		return 0, err
	}
	distances, err := util.AtoiSlice(strings.Fields(strings.Split(lines[1], ": ")[1]))
	if err != nil {
		return 0, err
	}

	result := 1
	for i := 0; i < len(times); i++ {
		timeAllowed := times[i]
		recordDistance := distances[i]
		a := float64(-1)
		b := float64(timeAllowed)
		c := float64(-recordDistance) - 0.000001 // trick to beat record, not equal
		x1 := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
		x2 := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
		big := int(math.Floor(max(x1, x2)))
		small := int(math.Ceil(min(x1, x2)))
		// fmt.Printf("x1: %f, x2: %f, big: %d, small: %d\n", x1, x2, big, small)
		numWays := big - small + 1
		// fmt.Println("num ways:", numWays)
		result *= numWays
	}
	return result, nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)
	time, err := strconv.Atoi(strings.Replace(strings.Split(lines[0], ": ")[1], " ", "", -1))
	if err != nil {
		return 0, err
	}
	distance, err := strconv.Atoi(strings.Replace(strings.Split(lines[1], ": ")[1], " ", "", -1))
	if err != nil {
		return 0, err
	}

	timeAllowed := time
	recordDistance := distance
	a := float64(-1)
	b := float64(timeAllowed)
	c := float64(-recordDistance) - 0.000001 // trick to beat record, not equal
	x1 := (-b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	x2 := (-b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	big := int(math.Floor(max(x1, x2)))
	small := int(math.Ceil(min(x1, x2)))
	// fmt.Printf("x1: %f, x2: %f, big: %d, small: %d\n", x1, x2, big, small)
	numWays := big - small + 1
	// fmt.Println("num ways:", numWays)

	return numWays, nil
}
