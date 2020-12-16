package main

import (
	"fmt"
	"strconv"

	"github.com/unmultimedio/adventofcode/util"
)

const preamble = 25

var numbers []int

func main() {
	numbersRaw, _ := util.ReadLines("./input")

	for _, nRaw := range numbersRaw {
		n, err := strconv.Atoi(nRaw)
		if err != nil {
			panic(fmt.Sprintf("wrong number in input: %s, err: %v", nRaw, err))
		}
		numbers = append(numbers, n)
	}

	for i := preamble; i < len(numbers); i++ {
		if !isValid(i) {
			fmt.Printf("invalid number %d at index %d\n", numbers[i], i)
			return
		}
	}
}

func isValid(idx int) bool {
	expected := numbers[idx]
	for i := idx - preamble; i < idx; i++ {
		for j := i; j < idx; j++ {
			if numbers[i]+numbers[j] == expected {
				return true
			}
		}
	}

	return false
}
