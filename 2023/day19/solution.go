package day19

import (
	"strconv"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
)

type rule struct {
	condition string
	result    string
}

type (
	workflow    []rule
	workflowMap map[string]workflow
	part        map[string]int
)

func parseWorkflows(workflowLines []string) workflowMap {
	workflows := map[string]workflow{} // name to rules
	for _, line := range workflowLines {
		curlyi := strings.IndexRune(line, '{')
		name := line[:curlyi]
		ruleStrings := strings.Split(strings.Trim(line[curlyi:], "{}"), ",")
		for _, r := range ruleStrings {
			fs := strings.Split(r, ":")
			var cond string
			var res string
			if len(fs) == 1 {
				cond = "fallthrough"
				res = fs[0]
			} else {
				cond = fs[0]
				res = fs[1]
			}
			workflows[name] = append(workflows[name], rule{cond, res})
		}
	}
	return workflows
}

func parseParts(partLines []string) []part {
	parts := []part{}
	for _, line := range partLines {
		ratingStrings := strings.Split(strings.Trim(line, "{}"), ",")
		p := part{} // category to score
		for _, r := range ratingStrings {
			fs := strings.Split(r, "=")
			category := fs[0]
			score, err := strconv.Atoi(fs[1])
			if err != nil {
				panic("couldnt convert part score")
			}
			p[category] = score
		}
		parts = append(parts, p)
	}
	return parts
}

func isComparator(r rune) bool {
	return r == '<' || r == '>'
}

func passesCondition(cond string, p part) bool {
	if cond == "fallthrough" {
		return true
	}
	opi := strings.IndexFunc(cond, isComparator)
	category := cond[:opi]
	op := rune(cond[opi])
	score, err := strconv.Atoi(cond[opi+1:])
	if err != nil {
		panic("couldn't convert condition num to int")
	}

	switch op {
	case '<':
		return p[category] < score
	case '>':
		return p[category] > score
	default:
		panic("invalid operator")
	}
}

func nextWorkflowName(workflows workflowMap, cur string, p part) string {
	for _, r := range workflows[cur] {
		if passesCondition(r.condition, p) {
			return r.result
		}
	}
	panic("didn't match any rule including fallthrough")
}

func Part1(input string) (int, error) {
	insplit := strings.Split(input, "\n\n")
	workflowLines := parse.Lines(insplit[0])
	partLines := parse.Lines(insplit[1])

	workflows := parseWorkflows(workflowLines)

	parts := parseParts(partLines)

	// fmt.Println(workflows)
	// fmt.Println(parts)

	a := []part{}
	for _, p := range parts {
		cur := "in"
		for cur != "A" && cur != "R" {
			cur = nextWorkflowName(workflows, cur, p)
		}
		if cur == "A" {
			a = append(a, p)
		}
		// ignore R parts
	}

	total := 0
	for _, p := range a {
		partTotal := 0
		for _, v := range p {
			partTotal += v
		}
		total += partTotal
	}

	return total, nil
}

type span struct {
	start int
	end   int
}

type partspan map[string]span

type state struct {
	wfn string
	ps  partspan
}

func splitOnCondition(cond string, ps partspan) (partspan, partspan) {
	if cond == "fallthrough" {
		return ps, nil
	}
	opi := strings.IndexFunc(cond, isComparator)
	category := cond[:opi]
	op := rune(cond[opi])
	score, err := strconv.Atoi(cond[opi+1:])
	if err != nil {
		panic("couldn't convert condition num to int")
	}

	// make two copies of partspan
	a := make(partspan) // ps that matches
	b := make(partspan) // ps that doesn't match
	for k, v := range ps {
		a[k] = v
		b[k] = v
	}

	switch op {
	case '<':
		a[category] = span{a[category].start, min(a[category].end, score-1)}
		b[category] = span{max(b[category].start, score), b[category].end}
	case '>':
		a[category] = span{max(a[category].start, score+1), a[category].end}
		b[category] = span{b[category].start, min(b[category].end, score)}
	default:
		panic("invalid operator")
	}

	return a, b
}

// put state through workflow, return slice of split states
func processState(wfs workflowMap, st state) []state {
	res := []state{}
	curState := st
	for _, r := range wfs[st.wfn] {
		match, nonmatch := splitOnCondition(r.condition, curState.ps)
		res = append(res, state{r.result, match})
		curState.ps = nonmatch
	}
	return res
}

func Part2(input string) (int, error) {
	insplit := strings.Split(input, "\n\n")
	workflowLines := parse.Lines(insplit[0])

	workflows := parseWorkflows(workflowLines)

	startps := partspan{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}

	initState := state{
		wfn: "in",
		ps:  startps,
	}

	a := []partspan{}
	stateQueue := []state{initState}
	for len(stateQueue) != 0 {
		// pop from queue
		st := stateQueue[0]
		stateQueue = stateQueue[1:]

		res := processState(workflows, st)
		for _, r := range res {
			if r.wfn == "A" {
				a = append(a, r.ps)
			} else if r.wfn == "R" {
				// do nothing
			} else {
				stateQueue = append(stateQueue, r)
			}
		}
	}

	total := 0
	for _, ps := range a {
		// fmt.Println(ps)
		spanTotal := 1
		for _, v := range ps {
			spanTotal *= v.end - v.start + 1
		}
		// fmt.Println(spanTotal)
		total += spanTotal
	}

	return total, nil
}
