package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

var (
	maskRx = regexp.MustCompile(`mask = ([X01]{36})`)
	memRx  = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
)

func main() {
	instructions, _ := util.ReadLines("./input")

	// part1(instructions)
	part2(instructions)
}

func part2(instructions []string) {
	var mask, maskOr string
	var xpos []int
	var variations int
	mem := make(map[int]int64)
	for _, ins := range instructions {
		fmt.Println("--- ", ins)
		if maskRx.MatchString(ins) {
			m := maskRx.FindAllStringSubmatch(ins, -1)
			mask = m[0][1]
			maskOr = strings.ReplaceAll(mask, "X", "0")
			xpos = make([]int, 0)
			for i, c := range mask {
				if c == 'X' {
					xpos = append(xpos, i)
				}
			}
			variations = int(math.Pow(float64(2), float64(len(xpos))))

			fmt.Printf("mask set:\n%s\n", mask)
			fmt.Printf("xpos: %v\n", xpos)
			fmt.Printf("address variations: %d\n", variations)
			fmt.Printf("mask or :\n%s\n", maskOr)
			continue
		}

		if memRx.MatchString(ins) {
			m := memRx.FindAllStringSubmatch(ins, -1)
			addS, valS := m[0][1], m[0][2]
			add, _ := strconv.Atoi(addS)
			val, _ := strconv.Atoi(valS)
			fmt.Printf("mem[%d]=%d\n", add, val)

			nOr, _ := strconv.ParseInt(maskOr, 2, 64)
			addToMask := int64(add) | nOr

			addBin := fmt.Sprintf("%036b", add)
			addToMaskBin := fmt.Sprintf("%036b", addToMask)
			fmt.Printf("add %d:\n%s\n", add, addBin)
			fmt.Printf("add 2 mask %d:\n%s\n", addToMask, addToMaskBin)

			var addresses []int

			for i := 0; i < variations; i++ {
				variation := fmt.Sprintf("%0*b", len(xpos), i)
				fmt.Printf("to replace: %s\n", variation)
				for vIdx, xIdx := range xpos {
					addToMaskBin = replace(addToMaskBin, rune(variation[vIdx]), xIdx)
				}

				newAdd, _ := strconv.ParseInt(addToMaskBin, 2, 32)
				fmt.Printf("replaced mask %d:\n%s\n", newAdd, addToMaskBin)
				addresses = append(addresses, int(newAdd))
			}

			for _, add := range addresses {
				mem[add] = int64(val)
			}
			continue
		}
		panic(fmt.Sprintf("unrecognized instruction %q", ins))
	}

	var total int64
	for _, val := range mem {
		total += val
	}
	fmt.Println("part 2: ", total)
	// 43360294212 wrong answer
}

func replace(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func part1(instructions []string) {
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
