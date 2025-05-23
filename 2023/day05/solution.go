package day05

import (
	"cmp"
	
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/advent-of-code/2023/helpers/util"
	"github.com/zyedidia/generic/stack"
)

type span struct {
	start int
	end   int // inclusive
}

type mapping struct {
	src span
	dst span
}

type converter []mapping

func (c converter) convertSingle(i int) int {
	for _, m := range c {
		if m.src.start <= i && i <= m.src.end {
			return i + (m.dst.start - m.src.start)
		}
	}
	return i
}

func (c converter) convertSpan(s span) []span {
	st := stack.New[span]()
	st.Push(s)
	a := []span{}

	for _, m := range c {
		nst := stack.New[span]()
		for st.Size() > 0 {
			z := st.Pop()

			before := span{z.start, min(z.end, m.src.start-1)}
			overlap := span{max(z.start, m.src.start), min(z.end, m.src.end)}
			after := span{max(z.start, m.src.end+1), z.end}

			if before.end > before.start {
				nst.Push(before)
			}
			if overlap.end > overlap.start {
				transform := m.dst.start - m.src.start
				a = append(a, span{overlap.start + transform, overlap.end + transform})
			}
			if after.end > after.start {
				nst.Push(after)
			}
		}
		st = nst
	}

	for st.Size() > 0 {
		a = append(a, st.Pop())
	}

	return a
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)

	seeds, err := parseSeeds(lines[0])
	if err != nil {
		return 0, err
	}

	convs, err := parseConverters(lines[1:])
	if err != nil {
		return 0, err
	}

	locations := []int{}
	for _, v := range seeds {
		for _, c := range convs {
			v = c.convertSingle(v)
		}
		locations = append(locations, v)
	}
	return slices.Min(locations), nil
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)

	seedSpans, err := parseSeedSpans(lines[0])
	if err != nil {
		return 0, err
	}

	convs, err := parseConverters(lines[1:])
	if err != nil {
		return 0, err
	}

	locations := []span{}

	for _, s := range seedSpans {
		ss := []span{s}
		for _, c := range convs {
			nss := []span{}
			for _, sss := range ss {
				nss = append(nss, c.convertSpan(sss)...)
			}
			ss = nss
		}
		locations = append(locations, ss...)
	}

	minSpan := slices.MinFunc(locations, func(a, b span) int {
		return cmp.Compare(a.start, b.start)
	})

	return minSpan.start, nil
}

// parsers

func parseSeeds(line string) ([]int, error) {
	seeds, err := util.AtoiSlice(strings.Fields(strings.Split(line, ": ")[1]))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't convert seeds into ints; %q", err))
	}

	return seeds, nil
}

func parseSeedSpans(line string) ([]span, error) {
	seedfs, err := util.AtoiSlice(strings.Fields(strings.Split(line, ": ")[1]))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("couldn't convert seed fields into ints; %q", err))
	}
	if len(seedfs)%2 != 0 {
		return nil, errors.New("seed fields not even")
	}

	seedSpans := []span{}
	for i := 0; i < len(seedfs); i += 2 {
		start := seedfs[i]
		r := seedfs[i+1]
		seedSpans = append(seedSpans, span{start, start + r - 1})

	}

	return seedSpans, nil
}

func parseConverters(lines []string) ([]converter, error) {
	convs := []converter{}

	for i := 0; i <= len(lines); i++ {
		if strings.HasSuffix(lines[i], ":") {
			i++
			conv := converter{}
			for i < len(lines) && lines[i] != "" {
				fs, err := util.AtoiSlice(strings.Fields(lines[i]))
				if err != nil {
					return nil, errors.New(fmt.Sprintf("problem parsing mapping: couldn't convert to ints: %s", lines[i]))
				}
				if len(fs) != 3 {
					return nil, errors.New("problem parsing mapping: not three fields")
				}
				dst := fs[0]
				src := fs[1]
				r := fs[2]
				conv = append(conv, mapping{span{src, src + r - 1}, span{dst, dst + r - 1}})
				i++
			}
			convs = append(convs, conv)
		}
	}

	return convs, nil
}
