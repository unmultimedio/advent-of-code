package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

func main() {
	instructions, _ := util.ReadLines("./input")
	maskRx := regexp.MustCompile(`mask = ([X01]{36})`)
	memRx := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

	var mask, maskAnd, maskOr string
	mem := make(map[int]int64)
	for _, instruction := range instructions {
		if maskRx.MatchString(instruction) {
			m := maskRx.FindAllStringSubmatch(instruction, -1)
			mask = m[0][1]
			maskAnd = strings.ReplaceAll(mask, "X", "1")
			maskOr = strings.ReplaceAll(mask, "X", "0")
			fmt.Printf("mask set: %s\n", mask)
			fmt.Printf("mask and: %s\n", maskAnd)
			fmt.Printf("mask or : %s\n", maskOr)
		} else if memRx.MatchString(instruction) {
			m := memRx.FindAllStringSubmatch(instruction, -1)
			addS, valS := m[0][1], m[0][2]
			add, _ := strconv.Atoi(addS)
			val, _ := strconv.Atoi(valS)
			fmt.Printf("mem[%d]=%d\n", add, val)

			nAnd, _ := strconv.ParseInt(maskAnd, 2, 64)
			nOr, _ := strconv.ParseInt(maskOr, 2, 64)
			mem[add] = (int64(val) & nAnd) | nOr
		} else {
			panic(fmt.Sprintf("unrecognized instruction %q", instruction))
		}
	}

	var total int64
	for _, val := range mem {
		total += val
	}
	fmt.Println(total)
}
