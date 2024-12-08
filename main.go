package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Yolsh/AoC-2024/Day1"
	"github.com/Yolsh/AoC-2024/Day2"
	"github.com/Yolsh/AoC-2024/Day3"
	"github.com/Yolsh/AoC-2024/Day4"
	"github.com/Yolsh/AoC-2024/Day5"
	"github.com/Yolsh/AoC-2024/Day6"
	"github.com/Yolsh/AoC-2024/Day7"
	"github.com/Yolsh/AoC-2024/Day8"
	"github.com/Yolsh/AoC-2024/Day9"
)

func main() {
	var Ans string
	var i int16
	packages := []string{}
	files, err := os.ReadDir(".")
	check(err)
	fmt.Println("What Package would you like to run?")
	for _, val := range files {
		if !strings.ContainsRune(val.Name(), '.') {
			i++
			fmt.Printf("%v: %v\n", i, val.Name())
			packages = append(packages, val.Name())
		}
	}
	fmt.Scan(&Ans)
	val, err := strconv.ParseInt(Ans, 10, 64)
	check(err)
	switch val {
	case int64(slices.Index(packages, "Day1")) + 1:
		Day1.Run()
	case int64(slices.Index(packages, "Day2")) + 1:
		Day2.Run()
	case int64(slices.Index(packages, "Day3")) + 1:
		Day3.Run()
	case int64(slices.Index(packages, "Day4")) + 1:
		Day4.Run()
	case int64(slices.Index(packages, "Day5")) + 1:
		Day5.Run()
	case int64(slices.Index(packages, "Day6")) + 1:
		Day6.Run()
	case int64(slices.Index(packages, "Day7")) + 1:
		Day7.Run()
	case int64(slices.Index(packages, "Day8")) + 1:
		Day8.Run()
	case int64(slices.Index(packages, "Day9")) + 1:
		Day9.Run()
	default:
		fmt.Println("That option isn't available")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
