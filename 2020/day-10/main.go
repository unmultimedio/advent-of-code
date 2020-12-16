package main

import (
	"fmt"
	"strconv"

	"github.com/unmultimedio/adventofcode/util"
)

var numbers []int
var pathsFromIndex map[int]int

func main() {
	numbersRaw, _ := util.ReadLines("./input-sorted")

	numbers = append(numbers, 0) // port
	for _, nRaw := range numbersRaw {
		n, err := strconv.Atoi(nRaw)
		if err != nil {
			panic(fmt.Sprintf("wrong number in input: %s, err: %v", nRaw, err))
		}
		numbers = append(numbers, n)
	}
	numbers = append(numbers, numbers[len(numbers)-1]+3) // device

	differences := make(map[int]int)
	for i := 1; i < len(numbers); i++ {
		differences[numbers[i]-numbers[i-1]]++
	}
	fmt.Printf("differences in joltage: %v\n", differences)
	fmt.Printf("1-diff * 3-diff: %d\n", differences[1]*differences[3])

	pathsFromIndex = make(map[int]int)
	combs := pathsFrom(0)
	fmt.Printf("possible combinations: %d\n", combs)
}

func pathsFrom(idx int) int {
	if idx == len(numbers)-1 {
		return 1
	}

	paths, alreadyCalculated := pathsFromIndex[idx]
	if alreadyCalculated {
		return paths
	}

	for nxt := idx + 1; nxt < len(numbers); nxt++ {
		if numbers[nxt]-numbers[idx] <= 3 {
			paths += pathsFrom(nxt)
		} else {
			break
		}
	}
	pathsFromIndex[idx] = paths

	return paths
}
