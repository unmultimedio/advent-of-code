package main

import (
	"fmt"
	"strconv"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	numbersRaw, _ := util.ReadLines("./input-sorted")

	var numbers []int
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
}
