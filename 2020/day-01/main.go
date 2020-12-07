package main

import (
	"fmt"
	"strconv"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	expenses, _ := util.ReadLines("./input")
	values := make([]int, 0)

	for _, exp := range expenses {
		val, _ := strconv.Atoi(exp)
		values = append(values, val)
	}

	twoNumbers(values)
	threeNumbers(values)
}

func twoNumbers(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == 2020 {
				fmt.Printf(
					"[%d]: %d, [%d]: %d, (sum: 2020), (prod: %d)\n",
					i, values[i],
					j, values[j],
					values[i]*values[j],
				)
			}
		}
	}
}

func threeNumbers(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				if values[i]+values[j]+values[k] == 2020 {
					fmt.Printf(
						"[%d]: %d, [%d]: %d, [%d]: %d, (sum: 2020), (prod: %d)\n",
						i, values[i],
						j, values[j],
						k, values[k],
						values[i]*values[j]*values[k],
					)
				}
			}
		}
	}
}
