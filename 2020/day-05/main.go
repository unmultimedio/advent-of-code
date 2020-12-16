package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

var available map[int]struct{}
var taken map[int]struct{}

func main() {
	available = make(map[int]struct{})
	taken = make(map[int]struct{})
	preparePlane()

	boardingPasses, _ := util.ReadLines("./input")

	var maxID int
	for _, bp := range boardingPasses {
		_, _, id := decodeBoardingPass(bp)
		if id > maxID {
			maxID = id
		}

		taken[id] = struct{}{}
		delete(available, id)
	}

	fmt.Printf("the max id is: %d\n", maxID)

	for id := range available {
		_, bef := taken[id-1]
		_, aft := taken[id+1]
		if bef && aft {
			fmt.Printf("my seat has id: %d\n", id)
			break
		}
	}

}

func preparePlane() {
	for r := 0; r < 127; r++ {
		for c := 0; c < 7; c++ {
			available[seatID(r, c)] = struct{}{}
		}
	}
}

var boardingPassRx = regexp.MustCompile(`\A(F|B){7}(R|L){3}\z`)

func decodeBoardingPass(pass string) (row int, column int, id int) {
	if !boardingPassRx.MatchString(pass) {
		fmt.Printf("invalid boarding pass: %s\n", pass)
		return -1, -1, -1
	}

	rowStr := strings.ReplaceAll(
		strings.ReplaceAll(
			pass[0:7],
			"F", "0",
		),
		"B", "1",
	)
	r, err := strconv.ParseInt(rowStr, 2, 32)
	if err != nil {
		fmt.Printf("error decoding row: %v\n", err)
		return -1, -1, -1
	}
	row = int(r)

	colStr := strings.ReplaceAll(
		strings.ReplaceAll(
			pass[7:],
			"L", "0",
		),
		"R", "1",
	)
	c, err := strconv.ParseInt(colStr, 2, 32)
	if err != nil {
		fmt.Printf("error decoding column: %v\n", err)
		return -1, -1, -1
	}
	column = int(c)

	return row, column, seatID(row, column)
}

func seatID(row, column int) int {
	return (row * 8) + column
}
