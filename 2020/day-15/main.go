package main

import "fmt"

func main() {
	// start := []int{0, 3, 6} // sample
	start := []int{0, 8, 15, 2, 12, 1, 4} // my puzzle input
	lenStart := len(start)

	memIdx1Turn := make(map[int]int)
	memIdx2Turn := make(map[int]int)

	var lastSaid int
	for i, n := range start {
		memIdx1Turn[n] = i
		fmt.Printf("%04d: %d\n", i+1, n)
	}
	lastSaid = start[lenStart-1]

	for i := lenStart; i < 2020; i++ {
		var numToSay int
		_, lastSaidOnce := memIdx1Turn[lastSaid]
		if !lastSaidOnce {
			panic(fmt.Sprintf("last said %d is not in memIdx1", lastSaid))
		}

		_, lastSaidMore := memIdx2Turn[lastSaid]
		if lastSaidMore {
			numToSay = memIdx1Turn[lastSaid] - memIdx2Turn[lastSaid]
		} else {
			numToSay = 0
		}
		fmt.Printf("%04d: %d\n", i+1, numToSay)

		_, numToSayAlreadySaid := memIdx1Turn[numToSay]
		if numToSayAlreadySaid {
			memIdx2Turn[numToSay] = memIdx1Turn[numToSay]
		}

		memIdx1Turn[numToSay] = i
		lastSaid = numToSay
	}
}
