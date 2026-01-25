package main

import (
	"fmt"
	"aoc2025/common"
	"os"
	"strings"
	"strconv"
)


type Point struct {
	X int
	Y int
}

func readInput(fileName string) []Point {
	lines, _ := common.ReadLines(fileName)
	
	var points []Point
	for _, line := range lines {
		tokens := strings.Split(line, ",")	
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		points = append(points, Point{X: x, Y: y})
	}

	return points
}

func calcRect(p1 Point, p2 Point) int {
	dx := p1.X - p2.X
	if dx < 0 {
		dx = -dx 
	}

	dy := p1.Y - p2.Y

	if dy < 0 {
		dy = -dy
	}
	
	return (dx + 1) * (dy + 1)
	

}


func main() {
	args := os.Args[1:]
	
	var test bool
	if len(args) == 0 {
		test = false 
	} else if args[0] == "test" {
		test = true
	}

	var fileName string 
	if test {
		fileName = "../test_input.txt"
	} else {
		fileName = "../input.txt"
	}
	
	input := readInput(fileName)


	largest := 0

	for i := range len(input) {
		for j := i + 1 ; j < len(input); j++ {
			size := calcRect(input[i], input[j])
			if size > largest {
				largest = size
			}
		}
	}

	

	fmt.Println(largest)
}
