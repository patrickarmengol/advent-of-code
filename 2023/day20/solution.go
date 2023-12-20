package day20

import (
	"slices"
	"strings"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/parse"
	"github.com/patrickarmengol/megs/deque/lldeque"
)

// modmap (name to mod)

type pulse struct {
	source      string
	destination string
	kind        string // low or high
}

type mod interface {
	process(p pulse) []pulse
	outs() []string
}

type ffmod struct {
	name         string   // TODO: can be inferred from p.destination
	state        bool     // off or on (default off)
	connectedOut []string // names of connected out modules
}

func (ffm *ffmod) process(p pulse) []pulse {
	res := []pulse{}
	if p.kind == "low" {
		ffm.state = !ffm.state

		var kindToSend string
		if ffm.state {
			// case was off and turns on
			kindToSend = "high"
		} else {
			// case was on and turns off
			kindToSend = "low"
		}

		for _, c := range ffm.connectedOut {
			res = append(res, pulse{ffm.name, c, kindToSend})
		}
	}

	return res
}

func (ffm *ffmod) outs() []string {
	return ffm.connectedOut
}

type cmod struct {
	name         string
	rememberedIn map[string]string // kind of recent pulse for each connected in mod name (default to low)
	connectedOut []string
}

func (cm *cmod) process(p pulse) []pulse {
	res := []pulse{}
	cm.rememberedIn[p.source] = p.kind
	allHigh := true
	for _, v := range cm.rememberedIn {
		if v == "low" {
			allHigh = false
			break
		}
	}

	var kindToSend string
	if allHigh {
		kindToSend = "low"
	} else {
		kindToSend = "high"
	}

	for _, c := range cm.connectedOut {
		res = append(res, pulse{cm.name, c, kindToSend})
	}
	return res
}

func (cm *cmod) outs() []string {
	return cm.connectedOut
}

func (cm *cmod) initInput(in string) {
	cm.rememberedIn[in] = "low"
}

type bmod struct {
	connectedOut []string
}

func (bm *bmod) process(p pulse) []pulse {
	res := []pulse{}
	for _, c := range bm.connectedOut {
		res = append(res, pulse{"broadcaster", c, p.kind})
	}
	return res
}

func (bm *bmod) outs() []string {
	return bm.connectedOut
}

func pushButton(mods map[string]mod) (int, int) {
	pulsequeue := lldeque.New[pulse]()
	pulsequeue.PushBack(pulse{"button", "broadcaster", "low"})

	countLow := 0
	countHigh := 0
	for !pulsequeue.IsEmpty() {
		p := pulsequeue.PopFront()
		// fmt.Printf("%s -%s-> %s\n", p.source, p.kind, p.destination)
		if p.kind == "low" {
			countLow += 1
		} else if p.kind == "high" {
			countHigh += 1
		}
		if m, ok := mods[p.destination]; ok {
			pulsequeue.PushBack(m.process(p)...)
		}
	}

	return countLow, countHigh
}

func Part1(input string) (int, error) {
	lines := parse.Lines(input)
	// fmt.Println(lines)

	mods := map[string]mod{}

	for _, line := range lines {
		sd := strings.Split(line, " -> ")
		src := sd[0]
		dsts := strings.Split(sd[1], ", ")
		if src[0] == '%' {
			name := src[1:]
			mods[name] = &ffmod{name, false, dsts}
		} else if src[0] == '&' {
			name := src[1:]
			mods[name] = &cmod{name, make(map[string]string), dsts}
		} else if src == "broadcaster" {
			mods["broadcaster"] = &bmod{dsts}
		} else {
			panic("couldn't parse mod type")
		}
	}
	// iterate again to initialize inputs of conjunction modules
	for n, m := range mods {
		for _, o := range m.outs() {
			if oo, ok := mods[o].(*cmod); ok {
				oo.initInput(n)
			}
		}
	}

	totalLow := 0
	totalHigh := 0
	for i := 0; i < 1000; i++ {
		cl, ch := pushButton(mods)
		totalLow += cl
		totalHigh += ch
	}
	// fmt.Println(totalLow, totalHigh)
	// fmt.Println(totalLow * totalHigh)

	return totalLow * totalHigh, nil
}

func pushButtonRx(mods map[string]mod, brx string) []string {
	pulsequeue := lldeque.New[pulse]()
	pulsequeue.PushBack(pulse{"button", "broadcaster", "low"})

	leading := []string{}

	for !pulsequeue.IsEmpty() {
		p := pulsequeue.PopFront()
		// fmt.Printf("%s -%s-> %s\n", p.source, p.kind, p.destination)
		if p.destination == brx && p.kind == "high" {
			// fmt.Printf("%s -%s-> %s\n", p.source, p.kind, p.destination)
			leading = append(leading, p.source)
		}
		// if p.source == "gq" {
		// 	fmt.Printf("%s -%s-> %s\n", p.source, p.kind, p.destination)
		// }
		// if p.destination == "rx" && p.kind == "low" {
		// 	return true
		// }
		if m, ok := mods[p.destination]; ok {
			pulsequeue.PushBack(m.process(p)...)
		}
	}

	return leading
}

func Part2(input string) (int, error) {
	lines := parse.Lines(input)
	// fmt.Println(lines)

	mods := map[string]mod{}

	for _, line := range lines {
		sd := strings.Split(line, " -> ")
		src := sd[0]
		dsts := strings.Split(sd[1], ", ")
		if src[0] == '%' {
			name := src[1:]
			mods[name] = &ffmod{name, false, dsts}
		} else if src[0] == '&' {
			name := src[1:]
			mods[name] = &cmod{name, make(map[string]string), dsts}
		} else if src == "broadcaster" {
			mods["broadcaster"] = &bmod{dsts}
		} else {
			panic("couldn't parse mod type")
		}
	}
	// iterate again to initialize inputs of conjunction modules
	for n, m := range mods {
		for _, o := range m.outs() {
			if oo, ok := mods[o].(*cmod); ok {
				oo.initInput(n)
			}
		}
	}
	var beforeRx string
	leadingCycleLen := map[string]int{}
	for n, m := range mods {
		if b, ok := m.(*cmod); ok && slices.Contains(m.outs(), "rx") {
			beforeRx = n
			for k := range b.rememberedIn {
				leadingCycleLen[k] = 0
			}
			break
		}
	}

	// fmt.Println(beforeRx)

	total := 0
	for anyZero(leadingCycleLen) {
		total += 1
		// fmt.Println(total)
		for _, r := range pushButtonRx(mods, beforeRx) {
			if leadingCycleLen[r] == 0 {
				leadingCycleLen[r] = total
			}
		}
		// fmt.Println(leadingCycleLen)
	}

	// fmt.Println(leadingCycleLen)
	cycleLens := []int{}
	for _, l := range leadingCycleLen {
		cycleLens = append(cycleLens, l)
	}

	return lcmMulti(cycleLens...), nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmMulti(ns ...int) int {
	result := ns[0]
	for _, n := range ns[1:] {
		result = lcm(result, n)
	}
	return result
}

func anyZero(m map[string]int) bool {
	for _, v := range m {
		if v == 0 {
			return true
		}
	}
	return false
}
