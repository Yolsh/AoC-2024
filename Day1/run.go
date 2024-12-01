package Day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./Day1/Input.txt")
	if err != nil {
		panic(err)
	}
	extrapDat := string(data)
	lines := strings.Split(strings.ReplaceAll(extrapDat, "\r", ""), "\n")
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		lPlusR := strings.Split(line, "   ")
		lnum, err := strconv.Atoi(lPlusR[0])
		rnum, err := strconv.Atoi(lPlusR[1])
		if err != nil {
			panic(err)
		}
		left = append(left, lnum)
		right = append(right, rnum)
	}
	sort.Ints(left)
	sort.Ints(right)
	tot := 0
	for i := range left {
		tot += left[i] * count(right, left[i])
	}
	fmt.Println(tot)
}

func count(right []int, num int) int {
	total := 0
	for _, val := range right {
		if val == num {
			total++
		}
	}
	return total
}
