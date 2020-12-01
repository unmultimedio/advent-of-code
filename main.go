package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	puzzle1()
}

func puzzle1() {
	expenses, _ := readLines("./puzzle1")
	values := make([]int, 0)

	for _, exp := range expenses {
		val, _ := strconv.Atoi(exp)
		values = append(values, val)
	}

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values)-1; j++ {
			if values[i]+values[j] == 2020 {
				fmt.Printf(
					"[%d]: %d, [%d]: %d, (sum: 2020), (prod: %d)",
					i, values[i],
					j, values[j],
					values[i]*values[j],
				)
			}
		}
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
