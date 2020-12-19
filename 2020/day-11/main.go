package main

import (
	"fmt"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	lines, _ := util.ReadLines("./input-small")
	var seatMap [][]string

	for _, line := range lines {
		row := strings.Split(line, "")
		seatMap = append(seatMap, row)
	}

	var iterations [][][]string
	iterations = append(iterations, seatMap)

	for {
		nextIt, switched := nextIterationPart2(iterations[len(iterations)-1])

		fmt.Printf("\n--- iteration %d \n", len(iterations))
		for _, line := range nextIt {
			fmt.Println(line)
		}

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

func nextIterationPart2(m [][]string) ([][]string, int) {
	countVisible := func(x int, y int) (empty int, occupied int) {
		// visibleChairs is an array that represents the eight directions in which
		// the person seated can look. In these indexes:
		//
		// [0] [1] [2]
		// [3] [X] [4]
		// [5] [6] [7]
		visibleChairs := make([]string, 8)

		keepLooking := func() bool {
			for _, c := range visibleChairs {
				if len(c) == 0 {
					return true
				}
			}
			return false
		}

		clear := func(mode string) {
			clearIfEmpty := func(indexes ...int) {
				for _, i := range indexes {
					if len(visibleChairs[i]) == 0 {
						visibleChairs[i] = "*"
					}
				}
			}
			switch mode {
			case "left":
				clearIfEmpty(0, 3, 5)
			case "right":
				clearIfEmpty(2, 4, 7)
			case "top":
				clearIfEmpty(0, 1, 2)
			case "bottom":
				clearIfEmpty(5, 6, 7)
			default:
				panic(fmt.Sprintf("unexpected clear mode %q", mode))
			}
		}

		getIndex := func(i int, j int) int {
			if i < x && j < y {
				return 0
			} else if i == x && j < y {
				return 1
			} else if i > x && j < y {
				return 2
			} else if i < x && j == y {
				return 3
			} else if i > x && j == y {
				return 4
			} else if i < x && j > y {
				return 5
			} else if i == x && j > y {
				return 6
			} else if i > x && j > y {
				return 7
			}
			panic(fmt.Sprintf("unexpected values for getIndex %d, %d, %d, %d", i, j, x, y))
		}

		// layer is the amount of layers (think of distance from the person) in
		// which we're currently evaluating.
		var layer int
		for {
			layer++
			for i := x - layer; i <= x+layer; i += layer {
				if i < 0 {
					clear("top")
					continue
				}
				if i >= len(m) {
					clear("bottom")
					continue
				}
				for j := y - layer; j <= y+layer; j += layer {
					if j < 0 {
						clear("left")
						continue
					}
					if j >= len(m[i]) {
						clear("right")
						continue
					}
					if i == x && j == y {
						continue
					}

					idx := getIndex(i, j)
					switch m[i][j] {
					case "L":
						visibleChairs[idx] = "L"
					case "#":
						visibleChairs[idx] = "#"
					case ".":
					default:
						panic(fmt.Sprintf("unrecognized character %q in map[%d][%d]", m[i][j], i, j))
					}
				}
			}
			if !keepLooking() {
				break
			}
		}

		for _, c := range visibleChairs {
			switch c {
			case "L":
				empty++
			case "#":
				occupied++
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
				_, o := countVisible(i, j)
				if o == 0 {
					result[i][j] = "#"
					switched++
				}
			case "#": // occupied
				_, o := countVisible(i, j)
				if o >= 5 {
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

func nextIterationPart1(m [][]string) ([][]string, int) {
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
