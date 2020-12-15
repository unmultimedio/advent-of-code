package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passportsBytes, _ := ioutil.ReadFile("./input")
	passportsRaw := strings.Split(string(passportsBytes), "\n\n")

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
	byr, ok := pass["byr"] // (Birth Year)
	if !ok || !validInt(byr, 1920, 2002) {
		return false
	}
	iyr, ok := pass["iyr"] // (Issue Year)
	if !ok || !validInt(iyr, 2010, 2020) {
		return false
	}
	eyr, ok := pass["eyr"] // (Expiration Year)
	if !ok || !validInt(eyr, 2020, 2030) {
		return false
	}
	hgt, ok := pass["hgt"] // (Height)
	if !ok || !validHeight(hgt) {
		return false
	}
	hcl, ok := pass["hcl"] // (Hair Color)
	if !ok || !validHairColor(hcl) {
		return false
	}
	ecl, ok := pass["ecl"] // (Eye Color)
	if !ok || !validEyeColor(ecl) {
		return false
	}
	pid, ok := pass["pid"] // (Passport ID)
	return ok && validPassportID(pid)

	// optional
	// _, ok = pass["cid"] // (Country ID)
	// if !ok {
	// 	return false
	// }
}

func validInt(v string, min int, max int) bool {
	number, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	return (number >= min && number <= max)
}

func validHeight(v string) bool {
	if strings.HasSuffix(v, "cm") {
		return validInt(v[:len(v)-2], 150, 193)
	} else if strings.HasSuffix(v, "in") {
		return validInt(v[:len(v)-2], 59, 76)
	}
	return false
}

var hairColorRx = regexp.MustCompile(`\A#[0-9a-f]{6}\z`)

func validHairColor(v string) bool {
	return hairColorRx.MatchString(v)
}

func validEyeColor(v string) bool {
	switch v {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

var passportIDRx = regexp.MustCompile(`\A[0-9]{9}\z`)

func validPassportID(v string) bool {
	return passportIDRx.MatchString(v)
}
