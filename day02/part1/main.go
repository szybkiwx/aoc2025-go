package main

import ( 
	"fmt" 
	"aoc2025/common"
	"os"
	"strings"
	"strconv"
)

func isInvalid(token string) bool {
	l := len(token)
	if l % 2 == 1 {
		return false
	}

	mid := l / 2


	x := token[0:mid]
	y := token[mid:]

	if len(x) != len(y) {
		fmt.Println("oops")
		os.Exit(1)
	}

	//fmt.Printf("x=%v, y=%v\n", x, y)
	return x == y
}

func main() {
	//input, err :=common.ReadWholeFile("../test_input.txt")	
	input, err :=common.ReadWholeFile("../input.txt")	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input_str:= strings.TrimSpace(string(input[:]))
	
	var password int64 = 0

	for token := range strings.SplitSeq(input_str, ",") {
		//fmt.Println(token)
		parts := strings.Split(token, "-")

		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end,   _ := strconv.ParseInt(parts[1], 10, 64)
		if start >= end {
			fmt.Printf("start: %v, end: %v, token: '%v'\n", start, end, token)
			os.Exit(1)
		}
		for i:= start; i <= end; i++ {
			val := strconv.FormatInt(i, 10)
			//fmt.Println(val)
			if isInvalid(val) {

				password += i
			} 
		}	

	}

	fmt.Println(password)

}
