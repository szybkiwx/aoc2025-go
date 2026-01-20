package main

import ( 
	"fmt"
	"aoc2025/common"
	"strings"
	"strconv"
	"sort"
)


type Junction struct {
	X int
	Y int
	Z int
}

type JunctionPair struct {
	J1 Junction 
	J2 Junction 
	Distance int
}

/**
	we don't care about exact values, so no need to sqrt 
*/
func dist(j1 Junction, j2 Junction) int {
	dx := j1.X - j2.X
	dy := j1.Y - j2.Y
	dz := j1.Z - j2.Z
	return dx*dx + dy*dy + dz * dz  

}

func readInput(fileName string) []Junction {
	lines, _ := common.ReadLines(fileName)

	var junctions []Junction

	for _, line := range lines {
		tokens := strings.Split(line, ",")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])
		junctions = append(junctions, 
			Junction{
				X: x,
				Y: y,
				Z: z,
			},
		)
	}

	
	return junctions
 

}


func areEqual(j1 Junction, j2 Junction) bool {
	return j1.X == j2.X && j1.Y == j2.Y && j1.Z == j2.Z
}

func calcDistAndSort(junctions []Junction) []JunctionPair {
	var distances []JunctionPair
	for i, x := range junctions {


		for j:= i + 1 ; j < len(junctions);j++ {
			y:= junctions[j]
			
			distances = append(distances,JunctionPair{ J1: x, J2: y, Distance: dist(x, y)} )
		}
	}

	sort.Slice(distances, func(i, j int) bool {
  		return distances[i].Distance < distances[j].Distance
	})


	return  distances

}



func main() {
	test := false
	var input []Junction
	if test {
		input = readInput("../test_input.txt")	
	} else {
		input= readInput("../input.txt")	
	}
	var circuits []map[Junction]int

	distances:=calcDistAndSort(input)

	var lastJunction int


	availableJunctions:= make(map[Junction]int)

	for _, j := range input {
		availableJunctions[j] = 1
	}


	for _, distance := range distances {

		//fmt.Printf("circuits: %v\n", circuits)
		//fmt.Printf("j1: %v, j2: %v\n", distance.J1, distance.J2)

		var currentCircuit map[Junction]int

		var newCircuits	[]map[Junction]int
	
		for _, circuit := range circuits {
			_, ok1 := circuit[distance.J1]
			_, ok2 := circuit[distance.J2] 
			if currentCircuit == nil && (ok1 || ok2) {
				currentCircuit = circuit
				
			} else if currentCircuit != nil && (ok1 || ok2) {
				// merge other circuits if they have overspanned junctions
				for k := range circuit {
					currentCircuit[k] = 1
					delete(circuit, k)
				}
			}

		}

		if currentCircuit != nil {
			currentCircuit[distance.J1] = 1
			currentCircuit[distance.J2] = 1
		} else {
			newCircuit := map[Junction]int {
				distance.J1 : 1,
				distance.J2 : 1,
			}
			
			circuits = append(circuits, newCircuit)
		}

		for _, circuit := range circuits {
			if len(circuit) > 0 {
				newCircuits = append(newCircuits, circuit)
			}
		}

		delete(availableJunctions, distance.J1)
		delete(availableJunctions, distance.J2)

		if len(availableJunctions) == 0 {
			lastJunction = distance.J1.X * distance.J2.X
			break
		}

		circuits = newCircuits

	}

/*	fmt.Println(circuits)

	var circuitSizes []int 

	for _, circuit := range circuits {
		size := len(circuit)
		circuitSizes = append(circuitSizes, size)
	}

	sort.Slice(circuitSizes, func (i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]	
	})	
*/


	fmt.Println(lastJunction)
}
