package main

import (
	"fmt"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	lines, _ := util.ReadLines("./input")
	var seatMap [][]string

	for _, line := range lines {
		row := strings.Split(line, "")
		seatMap = append(seatMap, row)
	}

	var iterations [][][]string
	iterations = append(iterations, seatMap)

	for {
		nextIt, switched := nextIteration(iterations[len(iterations)-1])
		iterations = append(iterations, nextIt)
		if switched == 0 {
			fmt.Printf("stabilized after %d iterations\n", len(iterations))
			break
		}
	}

	e, o := countSeats(iterations[len(iterations)-1])
	fmt.Printf("stable map, empty: %d, occupied: %d\n", e, o)
}

func countSeats(m [][]string) (empty int, occupied int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			switch m[i][j] {
			case "L":
				empty++
			case "#":
				occupied++
			}
		}
	}

	return empty, occupied
}

func nextIteration(m [][]string) ([][]string, int) {
	countAdjacents := func(x int, y int) (empty int, occupied int) {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(m) {
				continue
			}
			for j := y - 1; j <= y+1; j++ {
				if j < 0 || j >= len(m[i]) {
					continue
				}
				if i == x && j == y {
					continue
				}
				switch m[i][j] {
				case "L", ".":
					empty++
				case "#":
					occupied++
				default:
					panic(fmt.Sprintf("unrecognized character %q in map[%d][%d]", m[i][j], i, j))
				}
			}
		}

		return empty, occupied
	}

	var switched int
	result := make([][]string, len(m))
	for i := 0; i < len(m); i++ {
		result[i] = make([]string, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			result[i][j] = m[i][j]
			switch m[i][j] {
			case "L": // empty
				_, o := countAdjacents(i, j)
				if o == 0 {
					result[i][j] = "#"
					switched++
				}
			case "#": // occupied
				_, o := countAdjacents(i, j)
				if o >= 4 {
					result[i][j] = "L"
					switched++
				}
			case ".": // floor
			default:
				panic(fmt.Sprintf("unrecognized character %q in map[%d][%d]", m[i][j], i, j))
			}
		}
	}

	return result, switched
}
