package Day2

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./Day2/Input.txt")
	if err != nil {
		panic(err)
	}
	extrapDat := string(data)
	lines := strings.Split(strings.ReplaceAll(extrapDat, "\r", ""), "\n")
	tot := 0
	for ln, line := range lines {
		report := strings.Split(line, " ")
		if checkReport(report, ln, false) {
			tot++
		}
	}
	fmt.Println(tot)
}

func checkReport(report []string, ln int, removed bool) bool {
	var increasing int = 0
	for i := 1; i < len(report); i++ {
		thisNum, err := strconv.Atoi(report[i])
		if err != nil {
			panic(err)
		}
		prevNum, err := strconv.Atoi(report[i-1])
		if err != nil {
			panic(err)
		}
		dif := thisNum - prevNum
		fmt.Printf("(%d, %d): %d - %d => (%d, %d)\n", ln, i, prevNum, thisNum, dif, increasing)
		if (math.Abs(float64(dif)) < 1 || math.Abs(float64(dif)) > 3) || (increasing != 0 && increasing != i-1) {
			if !removed {
				report = slices.Concat(report[:i-1], report[i:])
				fmt.Println(report)
				return checkReport(report, ln, true)
			} else {
				fmt.Println("fail")
				return false
			}
		}
		if dif >= 1 {
			increasing++
		}
	}
	return true
}
