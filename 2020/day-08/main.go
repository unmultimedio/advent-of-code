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
var changeAttempts map[int]struct{}
var alreadyFlipped bool

func main() {
	orders, _ = util.ReadLines("./input")
	changeAttempts = make(map[int]struct{})

	for {
		fmt.Println("--- starting over")
		accumulator = 0
		alreadyFlipped = false
		executed = make(map[int]struct{})
		keepTrying := executeOrder(0)
		if !keepTrying {
			break
		}
	}
}

func executeOrder(idx int) (keepTrying bool) {
	if idx == len(orders) {
		fmt.Printf("arrived to the final instruction + 1! index: %d. current accumulator: %d\n", idx, accumulator)
		return false
	}

	if idx < 0 || idx > len(orders) {
		fmt.Printf("index %d is out of bounds. current accumulator: %d\n", idx, accumulator)
		return true
	}

	_, alreadyExecuted := executed[idx]
	if alreadyExecuted {
		fmt.Printf("order %d was already executed. current accumulator: %d\n", idx, accumulator)
		return true
	}

	executed[idx] = struct{}{}
	order := orders[idx]
	parts := strings.Split(order, " ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("wrong order, not two parts: %q\n", order))
	}

	argument, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(fmt.Sprintf("wrong order, bad argument: %q\n", parts[1]))
	}

	switch parts[0] {
	case "acc":
		accumulator += argument
		return executeOrder(idx + 1)
	case "jmp":
		_, alreadyTried := changeAttempts[idx]
		if alreadyTried || alreadyFlipped {
			return executeOrder(idx + argument)
		}
		// execute nop
		alreadyFlipped = true
		changeAttempts[idx] = struct{}{}
		fmt.Printf("flipping %d: %q\n", idx, order)
		return executeOrder(idx + 1)
	case "nop":
		_, alreadyTried := changeAttempts[idx]
		if alreadyTried || alreadyFlipped {
			return executeOrder(idx + 1)
		}
		// execute jmp
		alreadyFlipped = true
		changeAttempts[idx] = struct{}{}
		fmt.Printf("flipping %d: %q\n", idx, order)
		return executeOrder(idx + argument)
	}

	panic(fmt.Sprintf("wrong order %q, unknown operation: %q\n", order, parts[0]))
}
