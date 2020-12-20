package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/unmultimedio/adventofcode/util"
)

type ship struct {
	lat int
	lon int
	dir string
}

var turns = map[string]map[string]map[int]string{
	"N": {
		"R": {
			90:  "E",
			180: "S",
			270: "W",
		},
		"L": {
			90:  "W",
			180: "S",
			270: "E",
		},
	},
	"S": {
		"R": {
			90:  "W",
			180: "N",
			270: "E",
		},
		"L": {
			90:  "E",
			180: "N",
			270: "W",
		},
	},
	"E": {
		"R": {
			90:  "S",
			180: "W",
			270: "N",
		},
		"L": {
			90:  "N",
			180: "W",
			270: "S",
		},
	},
	"W": {
		"R": {
			90:  "N",
			180: "E",
			270: "S",
		},
		"L": {
			90:  "S",
			180: "E",
			270: "N",
		},
	},
}

func parseInstruction(instruction string) (string, int) {
	a := string(instruction[0])
	switch a {
	case "N", "S", "E", "W", "L", "R", "F":
	default:
		panic(fmt.Sprintf("invalid instruction action %q", instruction))
	}

	v, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(fmt.Sprintf("invalid instruction value %q", instruction))
	}

	return a, v
}

func (s *ship) move(action string, value int) {
	switch action {
	case "N":
		s.lat += value
	case "S":
		s.lat -= value
	case "E":
		s.lon += value
	case "W":
		s.lon -= value
	case "L":
		s.dir = turns[s.dir][action][value]
	case "R":
		s.dir = turns[s.dir][action][value]
	case "F":
		s.move(s.dir, value)
	default:
		panic(fmt.Sprintf("invalid action %q", action))
	}
}

func main() {
	instructions, _ := util.ReadLines("./input")

	myShip := ship{dir: "E"}
	for _, instruction := range instructions {
		myShip.move(parseInstruction(instruction))
	}

	fmt.Printf("current position: %d, %d, %s\n", myShip.lat, myShip.lon, myShip.dir)
	fmt.Printf("manhattan distance: %d\n", int(math.Abs(float64(myShip.lat))+math.Abs(float64(myShip.lon))))
}
