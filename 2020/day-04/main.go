package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	passportsBytes, _ := ioutil.ReadFile("./input")
	// fmt.Println(string(passportsBytes))

	passportsRaw := strings.Split(string(passportsBytes), "\n\n")
	// fmt.Println(len(passportsRaw), passportsRaw[0])

	var validPassports int
	for _, passportRaw := range passportsRaw {
		pass := make(map[string]string)
		fields := strings.Fields(passportRaw)
		for _, f := range fields {
			parts := strings.Split(f, ":")
			if len(parts) != 2 {
				fmt.Printf("warning: invalid field %q for passport %q\n", fields, passportRaw)
				continue
			}
			pass[parts[0]] = parts[1]
		}
		if isValidPassport(pass) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func isValidPassport(pass map[string]string) bool {
	var ok bool
	_, ok = pass["byr"] // (Birth Year)
	if !ok {
		return false
	}
	_, ok = pass["iyr"] // (Issue Year)
	if !ok {
		return false
	}
	_, ok = pass["eyr"] // (Expiration Year)
	if !ok {
		return false
	}
	_, ok = pass["hgt"] // (Height)
	if !ok {
		return false
	}
	_, ok = pass["hcl"] // (Hair Color)
	if !ok {
		return false
	}
	_, ok = pass["ecl"] // (Eye Color)
	if !ok {
		return false
	}
	_, ok = pass["pid"] // (Passport ID)
	return ok

	// optional
	// _, ok = pass["cid"] // (Country ID)
	// if !ok {
	// 	return false
	// }
}
