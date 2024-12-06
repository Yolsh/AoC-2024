package Day6

import (
	"fmt"
	"os"
	"regexp"
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

var re = regexp.MustCompile(`(^(X|O)(X|O)+#$)|(^#(X|O)(X|O)+$)`)

func Run() {
	data, err := os.ReadFile("./Day6/Input.txt")
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
	newmap := guardPath(guard, sitmap)
	tot := 1
	p2 := 0
	for _, line := range newmap {
		fmt.Printf("%s\n", line)
	}
	for _, line := range newmap {
		tot += strings.Count(line, "X")
		tot += strings.Count(line, "O")
		p2 += strings.Count(line, "O")
	}
	fmt.Println(tot, p2)
}

func guardPath(guard Guard, sitmap [][]rune) []string {
	for (guard.X > 0 && guard.X < len(sitmap[0])-1) && (guard.Y > 0 && guard.Y < len(sitmap)-1) {
		if sitmap[guard.Y+guard.rotation.Y][guard.X+guard.rotation.X] == '#' {
			guard.rotation = rotateGuard(guard)
		}
		if intersected(guard, rotateGuard(guard), sitmap) {
			sitmap[guard.Y+guard.rotation.Y][guard.X+guard.rotation.X] = 'O'
		}
		if sitmap[guard.Y][guard.X] != 'O' {
			sitmap[guard.Y][guard.X] = 'X'
		}
		guard.X += guard.rotation.X
		guard.Y += guard.rotation.Y
	}
	var result []string
	for _, line := range sitmap {
		result = append(result, string(line))
	}
	fmt.Printf("%+v\n", guard)
	return result
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

func intersected(guard Guard, rog Rotation, sitmap [][]rune) bool {
	var rLine string
	if rog.Y != 0 {
		for Y := range sitmap {
			if rog.Y == 1 && Y > guard.Y {
				rLine += string(sitmap[Y][guard.X])
			} else if rog.Y == -1 && Y < guard.Y {
				rLine += string(sitmap[Y][guard.X])
			}
		}
	} else {
		for X := range sitmap[guard.Y] {
			if rog.X == 1 && X > guard.X {
				rLine += string(sitmap[guard.Y][X])
			} else if rog.X == -1 && X < guard.X {
				rLine += string(sitmap[guard.Y][X])
			}
		}
	}
	return re.MatchString(strings.Trim(rLine, "."))
}
