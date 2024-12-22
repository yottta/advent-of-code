package main

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/yottta/advent-of-code/00_aoc"
)

var verbose bool

func main() {
	aoc.BasicRun(part1, part2)
}

func part1(content []string) {
	monkeys := createMonkeys(content, false)
	solve(monkeys, 20)
}

func part2(content []string) {
	monkeys := createMonkeys(content, true)

	// since the part 2 is stating that
	// "Worry levels are no longer divided by three after each item is inspected; you'll need to find another way to keep your worry levels manageable",
	// the worry levels will always increase now
	// BUT we can't just mod by any monkey's testBy number, because they're all throwing the items around,
	// so find a shared common denominator that can be used to keep the numbers under overflow
	bigMod := 1
	for _, m := range monkeys {
		bigMod = bigMod * m.targetRes.divBy
	}
	for i := range monkeys {
		monkeys[i].mod.mod = &bigMod
	}

	solve(monkeys, 10000)
}

func solve(monkeys []monkey, rounds int) {
	for r := 0; r < rounds; r++ {
		for i := range monkeys {
			analyzed := monkeys[i].analyzeItems()
			for target, items := range analyzed {
				monkeys[target].addItem(items...)
			}
		}
		logMonkeyItems(r+1, monkeys)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsAnalyzed > monkeys[j].itemsAnalyzed
	})
	fmt.Println(monkeys[0].itemsAnalyzed * monkeys[1].itemsAnalyzed)
}

func createMonkeys(content []string, ignoreRelief bool) []monkey {
	var (
		res     []monkey
		current monkey
	)
	for _, line := range content {
		if strings.Contains(line, "Monkey") {
			continue
		}
		if strings.TrimSpace(line) == "" {
			current.ignoreRelief = ignoreRelief
			res = append(res, current)
			current = monkey{id: current.id + 1, ignoreRelief: ignoreRelief}
			continue
		}
		if strings.Contains(line, "Starting") {
			split := strings.Split(line, ":")
			current.items = []item{}
			for _, it := range strings.Split(split[1], ",") {
				worryVal, err := strconv.Atoi(strings.TrimSpace(it))
				aoc.Must(err)
				current.addItem(item{worryLvl: worryVal})
			}
			continue
		}
		if strings.Contains(line, "Operation") {
			modifier, err := parseItemModifier(strings.TrimSpace(line))
			aoc.Must(err)
			current.mod = *modifier
			continue
		}
		if strings.Contains(line, "Test") {
			resolver, err := newTargetResolver(strings.TrimSpace(line))
			aoc.Must(err)
			current.targetRes = *resolver
			continue
		}
		if strings.Contains(line, "If") {
			aoc.Must(current.targetRes.processTargetInfo(strings.TrimSpace(line)))
			continue
		}
		aoc.Must(fmt.Errorf("invalid line '%s'", line))
	}
	res = append(res, current)
	return res
}

type monkey struct {
	id            int
	items         []item
	mod           itemModifier
	targetRes     targetResolver
	itemsAnalyzed int
	ignoreRelief  bool
}

// returns a map[target_monkey][]items_to_be_given_to_the_target
func (m *monkey) analyzeItems() map[int][]item {
	res := make(map[int][]item)
	for _, i := range m.items {
		m.itemsAnalyzed++
		newVal := m.mod.apply(i.worryLvl)
		if !m.ignoreRelief {
			newVal = newVal / 3
		}
		target := m.targetRes.resolveTarget(newVal)
		items := res[target]
		items = append(items, item{worryLvl: newVal})
		res[target] = items
	}
	m.items = []item{}
	return res
}

func (m *monkey) addItem(i ...item) {
	m.items = append(m.items, i...)
}

type item struct {
	worryLvl int
}

type variable struct {
	resolved *int
}

func (v variable) resolvedOrDefault(def int) int {
	if v.resolved != nil {
		return *v.resolved
	}
	return def
}

func parseVariable(val string) variable {
	intVal, err := strconv.Atoi(strings.TrimSpace(val))
	if err == nil {
		return variable{resolved: &intVal}
	}
	return variable{}
}

type operation byte

const (
	add operation = iota + 1
	subtract
	multiply
	divide
)

func (o operation) apply(val1, val2 int) int {
	switch o {
	case add:
		return val1 + val2
	case subtract:
		return val1 - val2
	case multiply:
		return val1 * val2
	case divide:
		return val1 / val2
	}
	return val1 // considered as "0" operation so just return the given value
}

func parseOperation(sign string) (operation, error) {
	switch sign {
	case "+":
		return add, nil
	case "-":
		return subtract, nil
	case "*":
		return multiply, nil
	case "/":
		return divide, nil
	}
	return 0, fmt.Errorf("invalid operation sign %s", sign)
}

// this can process multiple variables applying the same operation:
// new = old * <on>
// new = old * old
// new = 3 * 3
// new = old * old * 4
// new = 5 + old
type itemModifier struct {
	v   []variable
	op  operation
	mod *int
}

func (im itemModifier) apply(on int) int {
	res := im.v[0].resolvedOrDefault(on)
	for i := 1; i < len(im.v); i++ {
		res = im.op.apply(res, im.v[i].resolvedOrDefault(on))
	}
	if im.mod != nil {
		res = res % *im.mod
	}
	return res
}

func parseItemModifier(info string) (*itemModifier, error) {
	equParts := strings.Split(info, " ") // Operation: new = old * VAL
	o, err := parseOperation(equParts[4])
	if err != nil {
		return nil, err
	}
	var1 := parseVariable(equParts[3])
	var2 := parseVariable(equParts[5])
	return &itemModifier{
		v:  []variable{var1, var2},
		op: o,
	}, nil
}

type targetResolver struct {
	divBy   int
	targets []int
}

func (tr *targetResolver) resolveTarget(val int) int {
	if val%tr.divBy == 0 {
		return tr.targets[0]
	}
	return tr.targets[1]
}

func (tr *targetResolver) processTargetInfo(info string) error {
	split := strings.Split(info, " ") // If true: throw to monkey VAL
	target, err := strconv.Atoi(split[5])
	if err != nil {
		return err
	}
	if tr.targets == nil {
		tr.targets = []int{-1, -1}
	}
	if strings.Contains(info, "true") {
		tr.targets[0] = target
	} else {
		tr.targets[1] = target
	}
	return nil
}

func newTargetResolver(info string) (*targetResolver, error) {
	split := strings.Split(info, " ") // Test: divisible by VAL
	divBy, err := strconv.Atoi(split[3])
	if err != nil {
		return nil, err
	}
	return &targetResolver{
		divBy: divBy,
	}, nil
}

func logMonkeyItems(round int, monkeys []monkey) {
	slog.Debug("Round %d ended\n", round)
	for _, m := range monkeys {
		itemNames := make([]string, len(m.items))
		for i, it := range m.items {
			itemNames[i] = strconv.Itoa(it.worryLvl)
		}
		slog.Debug("Monkey %d items: %s", m.id, strings.Join(itemNames, ","))
	}
}
