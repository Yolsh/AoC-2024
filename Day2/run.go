package Day2

import (
	"fmt"
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
	for _, line := range lines {
		report := strings.Split(line, " ")
		fmt.Println(report)
		if checkReport(report) {
			tot++
		} else {
			for i := range report {
				temp := slices.Concat(report[:i], report[i+1:])
				fmt.Printf("temp: %s\n", temp)
				if checkReport(temp) {
					tot++
					break
				}
			}
		}
	}
	fmt.Println(tot)
}

func checkReport(report []string) bool {
	var increasing bool
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
		if i == 1 && dif > 0 {
			increasing = true
		}
		if increasing {
			if dif < 1 || dif > 3 || dif < 0 {
				return false
			}
		} else {
			if dif > -1 || dif < -3 || dif > 0 {
				return false
			}
		}
	}
	return true
}
