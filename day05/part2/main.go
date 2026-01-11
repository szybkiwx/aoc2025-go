package main

import (
	"aoc2025/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"math"
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

func mergeRanges(ranges []Range) []Range {
	var mergedRanges []Range 
	
	sort.Slice(ranges, func(i, j int) bool {
	   return ranges[i].start < ranges[j].start
	})


	mergedRanges = append(mergedRanges, ranges[0])
	
	for _, curr := range ranges[1:] {
		
		prev := mergedRanges[len(mergedRanges) - 1]
		if curr.start <= prev.end {
			mergedRanges[len(mergedRanges) -1] = Range{start: prev.start, end:  int(math.Max(float64(curr.end), float64(prev.end))) }
		} else {
			mergedRanges = append(mergedRanges, curr)
		}  	
	}

	return mergedRanges

}

func main() {
	//ranges,_ := readInput("../test_input.txt")
	ranges, _ := readInput("../input.txt")

	mergedRanges := mergeRanges(ranges)

	all := 0

	for _, rng := range mergedRanges {
		fmt.Printf("s: %v, e: %v\n", rng.start, rng.end)
		all += (rng.end - rng.start + 1)
	}


	fmt.Println(all)

}
