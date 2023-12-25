package day24

import (
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
)

type vector struct {
	x float64
	y float64
	z float64
}

type hailstone struct {
	position vector
	velocity vector
}

func parseHailstones(input string) []hailstone {
	hailstones := []hailstone{}
	for _, line := range parse.Lines(input) {
		fs := strings.Split(line, " @ ")
		ps, err := util.AtoiSlice(strings.Split(fs[0], ", "))
		if err != nil {
			panic("problem parsing positions")
		}
		vs, err := util.AtoiSlice(strings.Split(fs[1], ", "))
		if err != nil {
			panic("problem parsing velocities")
		}
		hailstones = append(hailstones, hailstone{
			vector{float64(ps[0]), float64(ps[1]), float64(ps[2])},
			vector{float64(vs[0]), float64(vs[1]), float64(vs[2])},
		})
	}
	return hailstones
}

func checkIntersect2d(a, b hailstone) (vector, bool) {
	// r1(t) = p1 + t*v1
	// r2(s) = p2 + s*v2
	// p1 + t*v1 = p2 + s*v2
	// p1x + t*v1x = p2x + s*v2x
	// p1y + t*v1y = p2y + s*v2y
	// solve for t and s:
	// t = (delta_p * v2) / (v1 X v2)
	// t = (delta_p * v1) / (v1 X v2)

	cross := (a.velocity.x * b.velocity.y) - (a.velocity.y * b.velocity.x)

	// check parallel or coincident
	if cross == 0 {
		return vector{}, false
	}

	dp := vector{b.position.x - a.position.x, b.position.y - a.position.y, 0}

	t := ((dp.x * b.velocity.y) - (dp.y * b.velocity.x)) / cross
	s := ((dp.x * a.velocity.y) - (dp.y * a.velocity.x)) / cross

	if t < 0 || s < 0 {
		return vector{}, false
	}

	inter := vector{a.position.x + (t * a.velocity.x), a.position.y + (t * a.velocity.y), 0}
	return inter, true
}

func Part1(input string) (int, error) {
	hailstones := parseHailstones(input)

	total := 0
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if inter, ok := checkIntersect2d(hailstones[i], hailstones[j]); ok {
				// fmt.Println(hailstones[i], hailstones[j], inter)
				// might have problems with float comparisons here
				if inter.x >= 200000000000000 && inter.x <= 400000000000000 &&
					inter.y >= 200000000000000 && inter.y <= 400000000000000 {
					total += 1
				}
			}
		}
	}

	return total, nil
}

func Part2(input string) (int, error) {
	return 24, nil
}
