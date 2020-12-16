package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

var accumulator int
var orders []string
var executed map[int]struct{}

func main() {
	orders, _ = util.ReadLines("./input")
	executed = make(map[int]struct{})
	executeOrder(0)
}

func executeOrder(idx int) {
	if idx >= len(orders) {
		fmt.Printf("index %d is out of bounds\n", idx)
		return
	}

	_, alreadyExecuted := executed[idx]
	if alreadyExecuted {
		fmt.Printf("order %d was already executed. current accumulator: %d\n", idx, accumulator)
		return
	}

	executed[idx] = struct{}{}
	order := orders[idx]
	parts := strings.Split(order, " ")
	if len(parts) != 2 {
		fmt.Printf("wrong order, not two parts: %q\n", order)
		return
	}

	argument, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("wrong order, bad argument: %q\n", parts[1])
		return
	}
	switch parts[0] {
	case "acc":
		accumulator += argument
		executeOrder(idx + 1)
	case "jmp":
		executeOrder(idx + argument)
	case "nop":
		executeOrder(idx + 1)
	default:
		fmt.Printf("wrong order, unknown operation: %q\n", parts[0])
	}
}
