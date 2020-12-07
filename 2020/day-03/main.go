package main

import (
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

	const slopeX, slopeY = 3, 1
	var x int
	var trees int
	for y := 0; y < length; y = y + slopeY {
		if isTree(x, y) {
			trees = trees + 1
		}
		x = x + slopeX
	}

	println(trees)
}

func isTree(x, y int) bool {
	return geoMap[y%length][x%width] == "#"
}
