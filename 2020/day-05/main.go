package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	boardingPasses, _ := util.ReadLines("./input")

	var maxID int
	for _, bp := range boardingPasses {
		_, _, id := decodeBoardingPass(bp)
		if id > maxID {
			maxID = id
		}
	}

	fmt.Println(maxID)
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

	return row, column, (row * 8) + column
}
