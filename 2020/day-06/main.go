package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	answersBytes, _ := ioutil.ReadFile("./input")
	groupsAnswers := strings.Split(string(answersBytes), "\n\n")

	var totalYeses int
	for _, group := range groupsAnswers {
		totalYeses += countPositiveAnswers(group)
	}

	fmt.Printf("total yes: %d\n", totalYeses)
}

func countPositiveAnswers(groupAswers string) int {
	yeses := make(map[int]struct{})
	for _, s := range groupAswers {
		ascii := int(s)
		if ascii < 97 || ascii > 122 { // a-z
			continue
		}
		yeses[ascii] = struct{}{}
	}

	return len(yeses)
}
