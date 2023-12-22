package day22

import (
	"slices"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
	"github.com/patrickarmengol/megs/set/hashset"
)

type vertex struct {
	x int
	y int
	z int
}

type brick struct {
	name   int
	blocks []vertex
}

func NewBrick(n int, a, b vertex) brick {
	brick := brick{n, []vertex{}}
	if a.y == b.y && a.z == b.z {
		for i := min(a.x, b.x); i <= max(a.x, b.x); i++ {
			brick.blocks = append(brick.blocks, vertex{i, a.y, a.z})
		}
	} else if a.x == b.x && a.z == b.z {
		for i := min(a.y, b.y); i <= max(a.y, b.y); i++ {
			brick.blocks = append(brick.blocks, vertex{a.x, i, a.z})
		}
	} else if a.x == b.x && a.y == b.y {
		for i := min(a.z, b.z); i <= max(a.z, b.z); i++ {
			brick.blocks = append(brick.blocks, vertex{a.x, a.y, i})
		}
	} else {
		panic("uknown brick case")
	}
	return brick
}

func parseBricks(input string) []brick {
	bricks := []brick{}
	for i, line := range parse.Lines(input) {
		// could i parse lower z as brick a
		fs := strings.Split(line, "~")
		as, err := util.AtoiSlice(strings.Split(fs[0], ","))
		if err != nil {
			panic("problem parsing vertex numbers")
		}
		av := vertex{as[0], as[1], as[2]}
		bs, err := util.AtoiSlice(strings.Split(fs[1], ","))
		if err != nil {
			panic("problem parsing vertex numbers")
		}
		bv := vertex{bs[0], bs[1], bs[2]}
		bricks = append(bricks, NewBrick(i, av, bv))
	}

	return bricks
}

func drop(brs []brick) ([]brick, int) {
	// caterpillar falling (probably not the best implementation)
	curBlocks := hashset.New[vertex]()

	// deep copy bricks
	nbrs := []brick{}
	for _, br := range brs {
		nbr := brick{br.name, []vertex{}}
		for _, bl := range br.blocks {
			v := vertex{bl.x, bl.y, bl.z}
			curBlocks.Add(v)
			nbr.blocks = append(nbr.blocks, v)
		}
		nbrs = append(nbrs, nbr)
	}

	brickFell := hashset.New[int]()
	for {
		anyChanged := false
		// iter through bricks
		for i := range nbrs {
			// iter through blocks in brick
			canChange := true
			for _, bl := range nbrs[i].blocks {
				// check if can change
				below := vertex{bl.x, bl.y, bl.z - 1}
				if below.z == 0 || (curBlocks.Has(below) && !slices.Contains(nbrs[i].blocks, below)) {
					canChange = false
					break
				}
			}
			if canChange {
				anyChanged = true
				brickFell.Add(nbrs[i].name)
				// remove old from current state
				curBlocks.Remove(nbrs[i].blocks...)
				// drop down all blocks in brick
				for j := range nbrs[i].blocks {
					nbrs[i].blocks[j].z = nbrs[i].blocks[j].z - 1
				}
				// add new to current state
				curBlocks.Add(nbrs[i].blocks...)
			}
		}
		if !anyChanged {
			break
		}
	}

	return nbrs, brickFell.Len()
}

func Part1(input string) (int, error) {
	bricks := parseBricks(input)

	nbricks, _ := drop(bricks)

	// for _, br := range nbricks {
	// 	fmt.Printf("%d - %v\n", br.name, br.blocks)
	// }

	disintegratable := []int{}
	for i := range nbricks {
		a := []brick{}
		a = append(a, nbricks[:i]...)
		a = append(a, nbricks[i+1:]...)
		_, dh := drop(a)
		if dh == 0 {
			disintegratable = append(disintegratable, nbricks[i].name)
		}

	}

	// fmt.Println(disintegratable)

	return len(disintegratable), nil
}

func Part2(input string) (int, error) {
	bricks := parseBricks(input)

	nbricks, _ := drop(bricks)

	// for _, br := range nbricks {
	// 	fmt.Printf("%d - %v\n", br.name, br.blocks)
	// }

	totaldh := 0
	for i := range nbricks {
		a := []brick{}
		a = append(a, nbricks[:i]...)
		a = append(a, nbricks[i+1:]...)
		_, dh := drop(a)
		totaldh += dh

	}

	return totaldh, nil
}
