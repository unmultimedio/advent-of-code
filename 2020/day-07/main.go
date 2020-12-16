package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

type rule struct {
	container  string
	canContain map[string]int
}

var shinyGoldContainers map[string]struct{}

func main() {
	shinyGoldContainers = make(map[string]struct{})

	rulesRaw, _ := util.ReadLines("./input")

	var rules []rule
	for _, ruleRaw := range rulesRaw {
		rule, err := decodeRule(ruleRaw)
		if err != nil {
			fmt.Printf("skipped rule %q, err: %v\n", ruleRaw, err)
		}
		rules = append(rules, rule)
		// fmt.Println(rule)
	}

	completeShinyGoldInspection(rules)
	fmt.Printf("amount of bags that can contain shiny gold bags: %d\n", len(shinyGoldContainers))
}

func completeShinyGoldInspection(rules []rule) {
	var newContainersAdded int
	for _, rule := range rules {
		containersCopy := shinyGoldContainers
		for shinyGoldContainer := range containersCopy {
			_, canContain := rule.canContain[shinyGoldContainer]
			_, isContainerAlready := shinyGoldContainers[rule.container]
			if canContain && !isContainerAlready {
				shinyGoldContainers[rule.container] = struct{}{}
				newContainersAdded++
			}
		}
	}
	if newContainersAdded > 0 {
		completeShinyGoldInspection(rules)
	}
}

func decodeRule(r string) (rule, error) {
	// (some bag) contain (some children bags)
	parts := strings.Split(r, " contain ")
	if len(parts) != 2 ||
		!strings.HasSuffix(parts[0], " bags") ||
		!strings.HasSuffix(parts[1], ".") {
		return rule{}, errors.New(`rule has no " contain " divisor, no bags suffix in parts[0] or no final dot in part[1]`)
	}
	containerBag := parts[0][:len(parts[0])-5]

	childrenBags := make(map[string]int)
	// ( contain ) (bag 1), (bag 2), ... (bag n).
	children := strings.Split(parts[1], ", ")
	for _, child := range children {
		// (0:amount) (1:modifier) (2:color) (3:bag(s))
		parts := strings.Fields(child)
		if len(parts) != 4 {
			return rule{}, fmt.Errorf("child bag has no four fields: %q", child)
		}

		amount, err := strconv.Atoi(parts[0])
		if err != nil {
			return rule{}, fmt.Errorf("child bag has no valid amount: %q", parts[0])
		}

		childBag := parts[1] + " " + parts[2]
		if childBag == "shiny gold" {
			shinyGoldContainers[containerBag] = struct{}{}
		}
		childrenBags[childBag] = amount
	}

	return rule{
		container:  containerBag,
		canContain: childrenBags,
	}, nil
}
