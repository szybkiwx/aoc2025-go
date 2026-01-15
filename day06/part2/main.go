package main

import (
	"fmt"
	"aoc2025/common"
	"strings"
	"strconv"
	"bytes"
)


type Task struct {
	operator string
	operands []int64
}


func findNumberOfDigits(line string) []int{
	var digits []int
	
	currNUmber := 0 
	for _, chr := range line[1:] {
		
		if chr == '*' || chr == '+' {
			digits = append(digits, currNUmber)
			currNUmber = 0
		} else {
			currNUmber += 1
		}
	}
	
	
	return append(digits, currNUmber + 1)	
}


func parseInput(lines []string) []Task {
	digits := findNumberOfDigits(lines[len(lines) -1])

	
	noOfTasks := len(digits)

	tasks := make([]Task, noOfTasks)


	linePos:=0
	for i := range noOfTasks {
		d := digits[i]
		
		for pos := range d {
			var buffer bytes.Buffer
			for _, line := range lines[:len(lines)-1] {
				chr := line[linePos + pos]
				if chr != ' ' {
					buffer.WriteByte(chr)
				}
			}	
			val, _ := strconv.ParseInt(buffer.String(), 10, 64)	
			tasks[i].operands = append(tasks[i].operands, val)
		}
		linePos += (d + 1) 

	}


	for i, field := range strings.Fields(lines[len(lines) - 1]) { 
		tasks[i].operator = field
	}

	return tasks



}


func main() {
	//lines, _ := common.ReadLines("../test_input.txt")
	lines, _ := common.ReadLines("../input.txt")

	input := parseInput(lines)
	var total int64 = 0
	for _, task :=range input {
		var result int64 = 0
		if task.operator == "*" {
			var curr int64 = 1
			for _, op := range task.operands {
				curr = curr * op
			}
			result = curr
		} else {
			var curr int64 = 0
			for _, op:= range task.operands {
				curr += op
			}
			result = curr
		}
		total += result
	}



	fmt.Println(total)
}
