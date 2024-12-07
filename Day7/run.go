package Day7

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result   int
	operands []int
}

func Run() {
	fmt.Printf("part1: %d\n", p1(getInput()))
	fmt.Printf("part2: %d\n", p2(getInput()))
}

func p1(equations []Equation) int {
	tot := 0
	for _, eq := range equations {
		if slices.Contains(getResults(eq.operands), eq.result) {
			tot += eq.result
		}
	}
	return tot
}

func getResults(opands []int) []int {
	var out []int
	for n := range int(math.Pow(2, float64(len(opands)-1))) {
		result := opands[0]
		ops := strconv.FormatInt(int64(n), 2)
		for len(ops) != len(opands)-1 {
			ops = "0" + ops
		}
		for i := 1; i < len(opands); i++ {
			switch ops[i-1] {
			case '0':
				result += opands[i]
			case '1':
				result *= opands[i]
			}
		}
		out = append(out, result)
	}
	return out
}

func p2(equations []Equation) int {
	tot := 0
	for _, eq := range equations {
		if slices.Contains(getResultsP2(eq.operands), eq.result) {
			tot += eq.result
		}
	}
	return tot
}

func getResultsP2(opands []int) []int {
	var out []int
	for n := range int(math.Pow(3, float64(len(opands)-1))) {
		result := opands[0]
		ops := strconv.FormatInt(int64(n), 3)
		for len(ops) != len(opands)-1 {
			ops = "0" + ops
		}
		for i := 1; i < len(opands); i++ {
			switch ops[i-1] {
			case '0':
				result += opands[i]
			case '1':
				result *= opands[i]
			case '2':
				res, err := strconv.Atoi(strconv.FormatInt(int64(result), 10) + strconv.FormatInt(int64(opands[i]), 10))
				if err != nil {
					panic(err)
				}
				result = res
			}
		}
		out = append(out, result)
	}
	return out
}

func getInput() []Equation {
	data, err := os.ReadFile("./Day7/Input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")
	var out []Equation
	for _, line := range lines {
		l := strings.Split(line, " ")
		result, err := strconv.Atoi(strings.ReplaceAll(l[0], ":", ""))
		if err != nil {
			panic(err)
		}
		var opands []int
		for i := 1; i < len(l); i++ {
			opand, err := strconv.Atoi(l[i])
			if err != nil {
				panic(err)
			}
			opands = append(opands, opand)
		}
		out = append(out, Equation{result: result, operands: opands})
	}
	return out
}
