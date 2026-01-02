package main

import (
	"os"
	"fmt"
	"aoc2025/common"
	"strconv"
)

func rotateLeft(pos int, value int) int {
	mValue := value % 100
	
	if mValue < pos {
		return pos - mValue
	} else if mValue == pos {
		return 0	
	} else {
		return 100 - (mValue - pos)
	}

}


func rotateRight(pos int, value int) int {
	mValue := value % 100

	return (pos + mValue ) % 100 

}

func main() {
	//lines, err := common.ReadLines("test_input.txt")
	lines, err := common.ReadLines("../input.txt")
	if err != nil {
		fmt.Printf("Error reading file")
		os.Exit(1)
	}

	position := 50
	passwrod := 0

	for _, element := range lines {
		rot := element[0]
		value, _ := strconv.Atoi(element[1:])
		
	
		if rot == 'L' {
			position = rotateLeft(position, value)
		} else {
			position = rotateRight(position, value)
		}

		fmt.Println(position)
	
		if position == 0 {
			passwrod++
		}


	}	

	fmt.Println(passwrod)


}
