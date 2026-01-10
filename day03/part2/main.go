package main

import (
	"fmt"
	"aoc2025/common"
	"math"
)
func findHighestLeftMost(slice string) (int64, int){
	var maxVal int64 = 0
	maxIndex := -1

	for i := 0; i < len(slice); i++ {
		val := int64(slice[i] - '0')
		if val >  maxVal {
			maxVal = val
			maxIndex = i
		}
	}

	return maxVal, maxIndex

}


func calcJoltage(bank string) int64 {
	totalIdx:=0
	var result int64 = 0	

	for i := range 12 {
		val, idx := findHighestLeftMost(bank[totalIdx:len(bank)-11+i])
		totalIdx += idx + 1
		result += (val * int64(math.Pow(10, float64(11-i)))) 
	}

	fmt.Printf("joltage for bank %v: %v\n", bank, result)
	return result
}

func main() {
	//lines, _ := common.ReadLines("../test_input.txt")
	lines, _ := common.ReadLines("../input.txt")
	var joltage int64 = 0
	for _, line := range lines {
		joltage += calcJoltage(line)
	}
	
	fmt.Println(joltage)	

}
