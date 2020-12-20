package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/unmultimedio/adventofcode/util"
)

type position struct {
	lat int
	lon int
}
type ship struct {
	pos position
	dir string
	wp  position
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

func (p position) rotate(direction string, value int) position {
	var newPos position
	switch direction {
	case "R":
		switch value {
		case 90:
			newPos.lat = -p.lon
			newPos.lon = p.lat
		case 180:
			newPos.lat = -p.lat
			newPos.lon = -p.lon
		case 270:
			newPos.lat = p.lon
			newPos.lon = -p.lat
		}
	case "L":
		switch value {
		case 90:
			newPos.lat = p.lon
			newPos.lon = -p.lat
		case 180:
			newPos.lat = -p.lat
			newPos.lon = -p.lon
		case 270:
			newPos.lat = -p.lon
			newPos.lon = p.lat
		}
	}
	return newPos
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
		s.wp.lat += value
	case "S":
		s.wp.lat -= value
	case "E":
		s.wp.lon += value
	case "W":
		s.wp.lon -= value
	case "L", "R":
		s.wp = s.wp.rotate(action, value)
	case "F":
		s.pos.lat += value * s.wp.lat
		s.pos.lon += value * s.wp.lon
	default:
		panic(fmt.Sprintf("invalid action %q", action))
	}
}

func main() {
	instructions, _ := util.ReadLines("./input")

	myShip := ship{dir: "E", wp: position{lat: 1, lon: 10}}
	for _, instruction := range instructions {
		myShip.move(parseInstruction(instruction))
	}

	fmt.Printf("current position: %d, %d, %s\n", myShip.pos.lat, myShip.pos.lon, myShip.dir)
	fmt.Printf("manhattan distance: %d\n", int(math.Abs(float64(myShip.pos.lat))+math.Abs(float64(myShip.pos.lon))))
}
