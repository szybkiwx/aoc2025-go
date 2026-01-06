package main

import (
	"fmt"
	"aoc2025/common"
)

func calcJoltage(bank string) int {
	maxVal := 0
	for i:=0 ; i < len(bank) - 1; i++ {
		
		for j:= i+1; j < len(bank); j++ {
			val := int(bank[i] - '0') * 10 + int(bank[j] - '0')
			if val > maxVal {
				maxVal = val
			}	
		}
	}
	fmt.Println(maxVal)
	return maxVal 
}

func main() {
	//lines, _ := common.ReadLines("../test_input.txt")
	lines, _ := common.ReadLines("../input.txt")
	joltage := 0
	for _, line := range lines {
		joltage += calcJoltage(line)
	}
	
	fmt.Println(joltage)	

}
