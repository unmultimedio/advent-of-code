package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

type pwdPolicy struct {
	n1, n2 int
	letter string
}

type password struct {
	policy pwdPolicy
	pass   string
}

var pwdPattern = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func main() {
	passwordsLines, _ := util.ReadLines("./input")

	var valid1, valid2 int
	for _, pwdRaw := range passwordsLines {
		matches := pwdPattern.FindAllStringSubmatch(pwdRaw, -1)
		m := matches[0]
		min, _ := strconv.Atoi(m[1])
		max, _ := strconv.Atoi(m[2])
		pwd := password{
			policy: pwdPolicy{
				n1:     min,
				n2:     max,
				letter: m[3],
			},
			pass: m[4],
		}
		if isValidPolicy1(pwd) {
			valid1 = valid1 + 1
		}
		if isValidPolicy2(pwd) {
			valid2 = valid2 + 1
		}
	}

	fmt.Printf("valid1: %d, valid2: %d", valid1, valid2)
}

// isValidPolicy1 checks that the "1-3 a" matches the letter "a" between 1-3 times.
func isValidPolicy1(pwd password) bool {
	count := strings.Count(pwd.pass, pwd.policy.letter)
	return count >= pwd.policy.n1 && count <= pwd.policy.n2
}

// isValidPolicy2 checks that the "1-3 a" has the letter "a" just once at indexes-1.
func isValidPolicy2(pwd password) bool {
	if len(pwd.pass) < pwd.policy.n2 {
		return false
	}

	l1 := string(pwd.pass[pwd.policy.n1-1])
	l2 := string(pwd.pass[pwd.policy.n2-1])
	return (l1 == pwd.policy.letter && l2 != pwd.policy.letter) ||
		(l1 != pwd.policy.letter && l2 == pwd.policy.letter)
}
