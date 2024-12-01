package Day2

import (
	"os"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./Day2/Input.txt")
	if err != nil {
		panic(err)
	}
	extrapDat := string(data)
	lines := strings.Split(strings.ReplaceAll(extrapDat, "\r", ""), "\n")
}
