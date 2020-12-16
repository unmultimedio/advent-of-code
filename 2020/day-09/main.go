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

	// for i := preamble; i < len(numbers); i++ {
	// 	if !isValid(i) {
	// 		fmt.Printf("invalid number %d at index %d\n", numbers[i], i)
	// 		return
	// 	}
	// }
	const firstInvalidNumber = 1124361034

	set := contiguosSetFor(firstInvalidNumber)
	min, max := set[0], set[0]
	for _, n := range set {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	fmt.Printf("smallest: %d, largest: %d, sum: %d\n", min, max, min+max)
}

func contiguosSetFor(target int) []int {
	for i := 0; i < len(numbers); i++ {
		var acc int
		for j := i; j < len(numbers); j++ {
			acc += numbers[j]
			if acc == target {
				fmt.Printf("set found for target %d at (%d,%d)\n", target, i, j)
				return numbers[i:j]
			}
			if acc > target {
				break
			}
		}
	}

	panic(fmt.Sprintf("couldn't find a contiguos set for number: %d", target))
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
