package Day4

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./Day4/Input.txt")
	if err != nil {
		panic(err)
	}
	extrapDat := string(data)
	lines := strings.Split(strings.ReplaceAll(extrapDat, "\r", ""), "\n")
	tot := 0
	for i, line := range lines {
		for j, ch := range line {
			if ch == 'A' {
				s3 := lookDiagF(lines, i, j)
				s4 := lookDiagB(lines, i, j)
				tot += checkFinds(s3, s4, i, j)
			}
		}
	}
	fmt.Println(tot)
}

func lookDiagF(lines []string, i, j int) string {
	if 1+i < len(lines) && 1+j < len(lines[0]) && j-1 >= 0 && i-1 >= 0 {
		return string(lines[i+1][j-1]) + string(lines[i][j]) + string(lines[i-1][j+1])
	} else {
		return ""
	}
}

func lookDiagB(lines []string, i, j int) string {
	if 1+i < len(lines) && 1+j < len(lines[0]) && j-1 >= 0 && i-1 >= 0 {
		return string(lines[i+1][j+1]) + string(lines[i][j]) + string(lines[i-1][j-1])
	} else {
		return ""
	}
}

func checkFinds(s3, s4 string, i, j int) int {
	tot := 0
	if (s3 == "MAS" || s3 == "SAM") && (s4 == "MAS" || s4 == "SAM") {
		tot++
		fmt.Printf("(%d, %d)\n", i, j)
	}
	return tot
}
