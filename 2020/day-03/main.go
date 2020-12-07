package main

import (
	"fmt"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

// from input
const width, length = 31, 323

var geoMap [][]string

func main() {
	geoMapRaw, _ := util.ReadLines("./input")

	geoMap = make([][]string, length)
	for i, line := range geoMapRaw {
		ltrs := strings.Split(line, "")
		geoMap[i] = ltrs
	}

	t0 := calculateTrees(1, 1)
	t1 := calculateTrees(3, 1)
	t2 := calculateTrees(5, 1)
	t3 := calculateTrees(7, 1)
	t4 := calculateTrees(1, 2)

	fmt.Println(t0 * t1 * t2 * t3 * t4)
}

func calculateTrees(slopeX, slopeY int) int {
	var x int
	var trees int

	for y := 0; y < length; y = y + slopeY {
		if isTree(x, y) {
			trees = trees + 1
		}
		x = x + slopeX
	}

	return trees
}

func isTree(x, y int) bool {
	return geoMap[y%length][x%width] == "#"
}
