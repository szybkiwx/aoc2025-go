package main

import (
	"os"
	"fmt"
	"aoc2025/common"
	"strconv"
)

func rotateLeft(pos int, value int) (int,int) {
	mValue := value % 100
	
	zero_pass := value / 100

	if pos == 0 {
		return 100 - mValue, zero_pass
	} else if mValue < pos {
		return pos - mValue, zero_pass
	} else if mValue == pos {
		return 0, zero_pass	
	} else {
		return 100 - (mValue - pos), zero_pass + 1 
	}

}


func rotateRight(pos int, value int) (int, int) {
	mValue := value % 100

	if (pos + mValue) % 100 == 0 {
		return (pos + mValue ) % 100, value/100
	} else {
		return (pos + mValue ) % 100, value/100 + (pos + mValue) / 100
	}
}

func main() {
	//lines, err := common.ReadLines("../test_input.txt")
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
		var pass int	
	
		if rot == 'L' {
			position, pass = rotateLeft(position, value)
		} else {
			position, pass = rotateRight(position, value)
		}

		//fmt.Printf("pos: %v, pass: %v\n", position, pass)
	
		if position == 0 {
			passwrod++
		}
		passwrod += pass 


	}	

	fmt.Println(passwrod)


}
