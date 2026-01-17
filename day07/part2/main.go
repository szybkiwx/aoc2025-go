package main

import (
	"fmt"
	"aoc2025/common"
)



func main() {
	lines, _ := common.ReadLines("../test_input.txt")
	//lines, _ := common.ReadLines("../input.txt")

	beams := make(map[int]int)

	for i, c := range lines[0] {
		if c == 'S' {
			beams[i] = 1
		}
	}

	partTraversed := 0

	for _, line := range lines[2:] {
		
		newBeams := make(map[int]int)
		for i, c := range line {
			if c == '^' && beams[i] > 0 {
				partTraversed += beams[i]
				newBeams[i-1] = beams[i]
				newBeams[i+1] = beams[i]
			} else if beams[i] > 0 {
				newBeams[i] = beams[i]
			}
		}

		beams = newBeams	
		

	}


	fmt.Println(partTraversed)
}
