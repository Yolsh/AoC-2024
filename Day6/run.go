package Day6

import (
	"fmt"
	"os"
	"strings"
)

type Rotation struct {
	X int
	Y int
}

type Guard struct {
	X        int
	Y        int
	rotation Rotation
}

func Run() {
	newmap := p1(getMap())
	tot := 1
	for _, line := range newmap {
		for _, ch := range line {
			if ch == 'X' {
				tot++
			}
			fmt.Print(string(ch))
		}
		fmt.Println("")
	}
	obstacles := p2(getMap())
	fmt.Println(tot, obstacles)
}

func p1(guard Guard, sitmap [][]rune) [][]rune {
	for (guard.X > 0 && guard.X < len(sitmap[0])-1) && (guard.Y > 0 && guard.Y < len(sitmap)-1) {
		if sitmap[guard.Y+guard.rotation.Y][guard.X+guard.rotation.X] == '#' {
			guard.rotation = rotateGuard(guard)
		}
		sitmap[guard.Y][guard.X] = 'X'
		guard.X += guard.rotation.X
		guard.Y += guard.rotation.Y
	}
	return sitmap
}

func rotateGuard(guard Guard) Rotation {
	if guard.rotation.X == 0 && guard.rotation.Y == -1 {
		return Rotation{1, 0}
	} else if guard.rotation.X == 1 && guard.rotation.Y == 0 {
		return Rotation{0, 1}
	} else if guard.rotation.X == 0 && guard.rotation.Y == 1 {
		return Rotation{-1, 0}
	} else {
		return Rotation{0, -1}
	}
}

func tryLoop(sX, sY int, sitmap [][]rune, hx, hy int) bool {
	var pArr [][]int
	for range sitmap {
		var pLine []int
		for range sitmap[0] {
			pLine = append(pLine, 0)
		}
		pArr = append(pArr, pLine)
	}
	sitmap[hy][hx] = '#'
	guard := Guard{sX, sY, Rotation{0, -1}}
	for (guard.X > 0 && guard.X < len(sitmap[0])-1) && (guard.Y > 0 && guard.Y < len(sitmap)-1) {
		if sitmap[guard.Y+guard.rotation.Y][guard.X+guard.rotation.X] == '#' {
			guard.rotation = rotateGuard(guard)
		}
		if pArr[guard.Y][guard.X] == 5 {
			sitmap[hy][hx] = '.'
			return true
		}
		pArr[guard.Y][guard.X]++
		guard.X += guard.rotation.X
		guard.Y += guard.rotation.Y
	}
	sitmap[hy][hx] = '.'
	return false
}

func p2(guard Guard, sitmap [][]rune) int {
	tot := 0
	sX, sY := guard.X, guard.Y
	for (guard.X > 0 && guard.X < len(sitmap[0])-1) && (guard.Y > 0 && guard.Y < len(sitmap)-1) {
		if sitmap[guard.Y+guard.rotation.Y][guard.X+guard.rotation.X] == '#' {
			guard.rotation = rotateGuard(guard)
		}
		if (guard.X != sX && guard.Y != sY) && tryLoop(sX, sY, sitmap, guard.X+guard.rotation.X, guard.Y+guard.rotation.Y) {
			tot++
			fmt.Printf("%d\n", tot)
		}
		guard.X += guard.rotation.X
		guard.Y += guard.rotation.Y
	}
	return tot
}

func getMap() (Guard, [][]rune) {
	data, err := os.ReadFile("./Day6/Test.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")
	var guard Guard
	var sitmap [][]rune
	for Y, line := range lines {
		var sitline []rune
		for X, ch := range line {
			if ch == '^' {
				guard = Guard{X: X, Y: Y, rotation: Rotation{X: 0, Y: -1}}
			}
			sitline = append(sitline, ch)
		}
		sitmap = append(sitmap, sitline)
	}
	return guard, sitmap
}
