package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	notes, _ := util.ReadLines("./input")

	earliest, err := strconv.Atoi(notes[0])
	if err != nil {
		panic(fmt.Sprintf("invalid earliest %q", notes[0]))
	}

	buses := strings.Split(notes[1], ",")

	fmt.Println(earliest, buses)

	minVal := earliest
	busWaits := make(map[int]int)
	var closestBusID int
	for _, bus := range buses {
		if bus == "x" {
			continue
		}

		busID, err := strconv.Atoi(bus)
		if err != nil {
			panic(fmt.Sprintf("invalid bus id %q", bus))
		}

		departedAgo := int(math.Mod(float64(earliest), float64(busID)))
		nextDeparture := busID - departedAgo

		fmt.Printf("bus id: %d departed %d mins ago, meaning will depart again in %d mins\n", busID, departedAgo, nextDeparture)
		busWaits[busID] = nextDeparture

		if nextDeparture < minVal {
			minVal = nextDeparture
			closestBusID = busID
		}
	}

	fmt.Printf("next bus id: %d\n", closestBusID)
	fmt.Printf("answer: %d\n", closestBusID*busWaits[closestBusID])
}
