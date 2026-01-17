package main

import (
	"fmt"
	"aoc2025/common"
)



func main() {
	//lines, _ := common.ReadLines("../test_input.txt")
	lines, _ := common.ReadLines("../input.txt")

	var beams map[int]bool
	beams = make(map[int]bool)

	for i, c := range lines[0] {
		if c == 'S' {
			beams[i] = true
		}
	}

	timesSplit := 0

	for row, line := range lines[2:] {
		if row % 2 == 1 {
			continue
		}

		newBeams := make(map[int]bool)
		for i, c := range line {
			if c == '^' && beams[i] {
				newBeams[i-1] = true
				newBeams[i+1] = true
				timesSplit ++
			} else if beams[i] {
				newBeams[i] = true
			}
		}

		beams = newBeams	

	}


	fmt.Println(timesSplit)
}
