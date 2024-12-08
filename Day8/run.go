package Day8

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type pos struct {
	X int
	Y int
}

func Run() {
	fmt.Printf("part1: %d\npart2: %d\n", p1(getInput()), p2(getInput()))
}

func p1(input []string, positions map[rune][]pos) int {
	allAntinodes := make([]pos, 0)
	for _, arrPos := range positions {
		for i, Pos := range arrPos {
			for j, relpos := range arrPos {
				if i == j {
					continue
				}
				nPos := pos{Pos.X + 2*(relpos.X-Pos.X), Pos.Y + 2*(relpos.Y-Pos.Y)}
				if !slices.Contains(allAntinodes, nPos) && nPos.X >= 0 && nPos.X < len(input[0]) && nPos.Y >= 0 && nPos.Y < len(input) {
					allAntinodes = append(allAntinodes, nPos)
				}
			}
		}
	}
	return len(allAntinodes)
}

func p2(input []string, positions map[rune][]pos) int {
	allAntinodes := make([]pos, 0)
	for _, arrPos := range positions {
		for i, Pos := range arrPos {
			for j, relpos := range arrPos {
				if i == j {
					continue
				}
				var nPos pos
				count := 0
				for nPos.X >= 0 && nPos.X < len(input[0]) && nPos.Y >= 0 && nPos.Y < len(input) {
					nPos = pos{Pos.X + count*(relpos.X-Pos.X), Pos.Y + count*(relpos.Y-Pos.Y)}
					if !slices.Contains(allAntinodes, nPos) && nPos.X >= 0 && nPos.X < len(input[0]) && nPos.Y >= 0 && nPos.Y < len(input) {
						allAntinodes = append(allAntinodes, nPos)
					}
					count++
				}
			}
		}
	}
	for _, Pos := range allAntinodes {
		input[Pos.Y] = input[Pos.Y][:Pos.X] + "#" + input[Pos.Y][Pos.X+1:]
	}
	for _, line := range input {
		fmt.Println(line)
	}
	return len(allAntinodes)
}

func getInput() ([]string, map[rune][]pos) {
	data, err := os.ReadFile("./Day8/Input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")
	out := make(map[rune][]pos)
	for Y, line := range lines {
		for X, ch := range line {
			if ch != '.' {
				if out[ch] == nil {
					out[ch] = make([]pos, 0)
				}
				out[ch] = append(out[ch], pos{X, Y})
			}
		}
	}
	return lines, out
}
