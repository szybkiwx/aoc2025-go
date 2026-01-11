package main

import (
	"fmt"
	"aoc2025/common"
	"strings"
)

func isAccessible(x int, y int, lines []string) bool {
	if lines[x][y] != '@' {
		return false 
	}

	var minX int
	var minY int
	var maxX int
	var maxY int

	if x > 0 {
		minX = x - 1
	} else {
		minX = 0
	}	

	if y > 0 {
		minY = y - 1
	} else {
		minY = 0
	}

	if x < len(lines) - 1 {
		maxX = x + 1
	} else {
		maxX = len(lines) - 1
	}

	if y < len(lines[0]) - 1 {
		maxY = y + 1
	} else {
		maxY = len(lines[0]) - 1
	}

	adj:=0
	fmt.Printf("x, y: %v, %v\n", x, y)
	fmt.Printf("minX: %v, minY:%v, maxX: %v, maxY: %v\n", minX, minY, maxX, maxY)

	for dx := minX; dx <= maxX; dx++ {

		for dy := minY; dy <= maxY; dy++ {

			//fmt.Printf("dx, dy: %v, %v\n", dx, dy)
			if dx == x && dy == y {
				continue
			}

			if lines[dx][dy] == '@' {
				adj += 1
			}
		}
	}

	fmt.Println(adj)

	return adj < 4

}


func main() {
	//lines, _ := common.ReadLines("../test_input.txt")
	lines, _ := common.ReadLines("../input.txt")

	newLines := make([]string, len(lines))
	removedRolls := 0
	for true {
		accessibleRolls := 0

		for x, line := range lines {
			var sb strings.Builder		
			for y, chr := range line {
				if isAccessible(x, y, lines) {
					accessibleRolls += 1
					sb.WriteString("x")
				} else {
					sb.WriteRune(chr)
				}
			}	
			newLines[x] = sb.String() 
		}
		removedRolls += accessibleRolls
		if accessibleRolls == 0 {
			break
		}
		lines = newLines
	}

	
	fmt.Println(removedRolls)
}
