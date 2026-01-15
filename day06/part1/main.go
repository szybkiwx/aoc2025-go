package main

import (
	"fmt"
	"aoc2025/common"
	"strings"
	"strconv"
)

type Task struct {
	operator string
	operands []int64
}


func parseInput(lines []string) []Task {


	s:= strings.Fields(lines[0])

	noOfTasks := len(s)

	tasks := make([]Task, noOfTasks)
	//for  i := range noOfTasks {
	//	tasks[i] = Task { operator: "", operands: make([]int64)
	//}


	for _, line := range lines[:len(lines)-1] {
		for i, field := range strings.Fields(line) {

			val, _ := strconv.ParseInt(field, 10, 64) 
			//tasks[i] = Task{operator: "", operands: append(tasks[i].operands, val)} 	
			tasks[i].operands = append(tasks[i].operands, val)
		}
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
