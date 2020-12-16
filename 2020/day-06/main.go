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

// you need to identify the questions to which everyone answered "yes"!, meaning
// I can iterate over the first one only.
func countPositiveAnswers(groupAswers string) int {
	peopleAnswers := strings.Split(groupAswers, "\n")

	yeses := make(map[int]struct{})
	// only first person
	for _, s := range peopleAnswers[0] {
		ascii := int(s)
		if ascii < 97 || ascii > 122 { // a-z
			continue
		}
		yeses[ascii] = struct{}{}

		// check if others have it as well, or we should remove.
		for idx, others := range peopleAnswers {
			if idx == 0 {
				continue
			}

			if !strings.Contains(others, string(s)) {
				delete(yeses, ascii)
			}
		}
	}

	return len(yeses)
}
