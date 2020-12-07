package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/unmultimedio/adventofcode/util"
)

type pwdPolicy struct {
	min, max int
	letter   string
}

type password struct {
	policy pwdPolicy
	pass   string
}

var pwdPattern = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func main() {
	passwordsLines, _ := util.ReadLines("./input")

	var valid int
	for _, pwdRaw := range passwordsLines {
		matches := pwdPattern.FindAllStringSubmatch(pwdRaw, -1)
		m := matches[0]
		min, _ := strconv.Atoi(m[1])
		max, _ := strconv.Atoi(m[2])
		pwd := password{
			policy: pwdPolicy{
				min:    min,
				max:    max,
				letter: m[3],
			},
			pass: m[4],
		}
		if isValid(pwd) {
			valid = valid + 1
		}
	}

	fmt.Println(valid)
}

func isValid(pwd password) bool {
	count := strings.Count(pwd.pass, pwd.policy.letter)
	return count >= pwd.policy.min && count <= pwd.policy.max
}
