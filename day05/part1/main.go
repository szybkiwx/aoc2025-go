package main

import (
	"aoc2025/common"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	start int 
	end int
}

func readInput(fileName string) ([]Range, []int) {
	raw_input, _ := common.ReadWholeFile(fileName)
	input := strings.TrimSpace(string(raw_input[:])) 
	parts := strings.Split(input, "\n\n")

	var ranges []Range

	for r := range strings.SplitSeq(parts[0], "\n") {
		tokens := strings.Split(r, "-")
		start, _ := strconv.Atoi(tokens[0])
		end, _ := strconv.Atoi(tokens[1])
		ranges = append(ranges, Range{start:start,   end: end })
	}

	var ingredients []int
	
	for i := range strings.SplitSeq(parts[1], "\n") {
		ing, _ := strconv.Atoi(i)
		ingredients = append(ingredients, ing)
	}
	
	return ranges, ingredients
	

}


func isInRange(ingredient int, ranges []Range) bool {
	for _, rng := range ranges {
		if ingredient >= rng.start && ingredient <= rng.end {
			return true
		}
	}
	return false 	

}

func main() {
	//ranges,_ := readInput("../test_input.txt")
	ranges, ingredients := readInput("../input.txt")

	fresh := 0
	for _, ingredient := range ingredients {
		if isInRange(ingredient, ranges) {
			fresh += 1
		}
	}


	fmt.Println(fresh)

}
