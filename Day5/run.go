package Day5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Order struct {
	before string
	after  string
}

func Run() {
	data, err := os.ReadFile("./Day5/Input.txt")
	if err != nil {
		panic(err)
	}
	extrapDat := string(data)
	lines := strings.Split(strings.ReplaceAll(extrapDat, "\r", ""), "\n")
	var rules []Order
	var updates []string
	for i := range lines {
		if lines[i] == "" {
			rules = parseRules(lines[:i])
			updates = lines[i+1:]
		}
	}
	tot := 0
	for _, update := range updates {
		ordered := IsOrdered(update, rules)
		if ordered {
			continue
		}
		for !ordered {
			for _, rule := range rules {
				if strings.Contains(update, rule.before) && strings.Contains(update, rule.after) {
					arr := strings.Split(update, ",")
					if !(slices.Index(arr, rule.before) < slices.Index(arr, rule.after)) {
						update = Swapper(arr, rule)
						ordered = IsOrdered(update, rules)
					}
				}
			}
		}
		fmt.Printf("%s: %t\n", update, !ordered)
		if ordered {
			arr := strings.Split(update, ",")
			middleStr := arr[len(arr)/2]
			num, err := strconv.Atoi(middleStr)
			if err != nil {
				panic(err)
			}
			tot += num
		}
	}
	fmt.Printf("total: %d\n", tot)
}

func parseRules(rules []string) []Order {
	var out []Order
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		out = append(out, Order{split[0], split[1]})
	}
	return out
}

func Swapper(update []string, rule Order) string {
	fmt.Printf("%s: (%s|%s)\n", strings.Join(update, ","), rule.before, rule.after)
	befIdx := slices.Index(update, rule.before)
	aftIdx := slices.Index(update, rule.after)
	temp := update[aftIdx]
	update[aftIdx] = update[befIdx]
	update[befIdx] = temp
	return strings.Join(update, ",")
}

func IsOrdered(update string, rules []Order) bool {
	ordered := true
	for _, rule := range rules {
		if strings.Contains(update, rule.before) && strings.Contains(update, rule.after) {
			arr := strings.Split(update, ",")
			if !(slices.Index(arr, rule.before) < slices.Index(arr, rule.after)) {
				ordered = false
			}
		}
	}
	return ordered
}
