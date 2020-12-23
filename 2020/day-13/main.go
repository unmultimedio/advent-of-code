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

	part1(earliest, buses)
	part2(buses)
}

func part2(buses []string) {
	// We'll apply chinese remainder theorem
	// https://www.youtube.com/watch?v=zIFehsBHB8o&ab_channel=MathswithJay
	type crtFormula struct {
		b    int // remainders (offset)
		n    int // busID
		bigN int
		x    int
	}

	var formulas []crtFormula
	for i, bus := range buses {
		if bus == "x" {
			continue
		}

		busID, err := strconv.Atoi(bus)
		if err != nil {
			panic(fmt.Sprintf("invalid bus id %q", bus))
		}
		formulas = append(formulas, crtFormula{b: busID - i, n: busID})
	}
	fmt.Printf("formulas to start: %+v\n", formulas)

	prodN := 1
	for _, f := range formulas {
		prodN *= f.n
	}
	fmt.Printf("prodN: %d\n", prodN)

	for i, f := range formulas {
		f.bigN = prodN / f.n
		fmt.Printf("--- iteration %d, formula: %+v\n", i, f)

		m := mod(f.bigN, f.n)
		newM := m
		x := 1
		fmt.Printf("m := mod(%d, %d): %d\n", f.bigN, f.n, m)
		for {
			if newM == 1 {
				f.x = x
				break
			}
			x++
			newM = mod(m*x, f.n)
			fmt.Printf("newM = mod(%d * %d, %d): %d\n", m, x, f.n, newM)
		}
		formulas[i] = f
	}
	fmt.Printf("formulas calculated: %+v\n", formulas)

	var bnx int
	for _, f := range formulas {
		bnx += f.b * f.bigN * f.x
	}
	fmt.Printf("bnx: %d\n", bnx)
	bnx -= (bnx / prodN) * prodN
	fmt.Printf("bnx simplified: %d\n", bnx)
}

func mod(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func part1(earliest int, buses []string) {
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
