package Day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	re := regexp.MustCompile(`mul\(([0-9]{1}|[0-9]{2}|[0-9]{3}),([0-9]{1}|[0-9]{2}|[0-9]{3})\)|do\(\)|don't\(\)`)
	data, err := os.ReadFile("./Day3/Input.txt")
	if err != nil {
		panic(err)
	}
	extrapDat := string(data)
	lines := strings.Split(strings.ReplaceAll(extrapDat, "\r", ""), "\n")
	tot := 0
	enabled := true
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else if enabled {
				nums := strings.Split(strings.Split(strings.Replace(match, ")", "", -1), "(")[1], ",")
				num1, err := strconv.Atoi(nums[0])
				num2, err := strconv.Atoi(nums[1])
				if err != nil {
					panic(err)
				}
				tot += num1 * num2
			}
		}
	}
	fmt.Println(tot)
}
