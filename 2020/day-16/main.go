package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

type valueRange struct {
	min int
	max int
}

type ticketRange struct {
	valueRange
	name string
}

type byMin []ticketRange

func (a byMin) Len() int           { return len(a) }
func (a byMin) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMin) Less(i, j int) bool { return a[i].min < a[j].min }

func main() {
	input, _ := util.ReadLines("./input")

	inputGroups := make([][]string, 3)
	var idx int
	for _, line := range input {
		if line == "" {
			idx++
			continue
		}

		inputGroups[idx] = append(inputGroups[idx], line)
	}
	// fmt.Println(groups)

	rangesRaw := inputGroups[0]
	var ticketLimitations []ticketRange
	for _, rr := range rangesRaw {
		rangeParts := strings.Split(rr, ": ")
		if len(rangeParts) != 2 {
			panic("unexpected field range: " + rr)
		}
		ticketField := rangeParts[0]

		ranges := strings.Split(rangeParts[1], " or ")
		for _, r := range ranges {
			rangeValues := strings.Split(r, "-")
			if len(rangeValues) != 2 {
				panic("unexpected field range: " + rr)
			}

			min, err := strconv.Atoi(rangeValues[0])
			if err != nil {
				panic("unexpected field range: " + rr)
			}
			max, err := strconv.Atoi(rangeValues[1])
			if err != nil {
				panic("unexpected field range: " + rr)
			}

			ticketLimitations = append(ticketLimitations, ticketRange{
				name: ticketField,
				valueRange: valueRange{
					min: min,
					max: max,
				},
			})
		}
	}

	// fmt.Println("parsed", ticketLimitations)

	sort.Sort(byMin(ticketLimitations))
	// fmt.Println("sorted", ticketLimitations)

	allTicketLimitations := mergeRanges(ticketLimitations)
	// fmt.Println("merged", allTicketLimitations)

	nearbyTickets := inputGroups[2]
	var ticketScanningErrorRate int
	for i, t := range nearbyTickets {
		if i == 0 {
			continue // ignore first line
		}

		values := strings.Split(t, ",")
		for _, vRaw := range values {
			v, err := strconv.Atoi(vRaw)
			if err != nil {
				panic("unexpected nearby ticket value: " + vRaw)
			}

			for _, tl := range allTicketLimitations {
				if v < tl.min || v > tl.max {
					ticketScanningErrorRate += v
				}
			}
		}
	}

	fmt.Println("ticket scanning error rate: ", ticketScanningErrorRate)
}

func mergeRanges(ranges []ticketRange) []valueRange {
	var stack []valueRange

	for _, r := range ranges {
		stackLen := len(stack)

		// empty stack, or current min bigger than last max
		if stackLen == 0 || r.min > stack[stackLen-1].max {
			stack = append(stack, r.valueRange)
			continue
		}

		// if current max is bigger than last max
		if r.max > stack[stackLen-1].max {
			stack[stackLen-1].max = r.max
		}

		// current is within last, just ignore
	}

	return stack
}
