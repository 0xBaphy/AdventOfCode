package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr *string
	iyr *string
	eyr *string
	hgt *string
	hcl *string
	ecl *string
	pid *string
	cid *string
}

/*
	checkByr (Birth Year) - four digits; at least 1920 and at most 2002.
*/
func (p Passport) checkByr() bool {
	if len(*p.byr) != 4 {
		return false
	}
	i, err := strconv.Atoi(*p.byr)
	if err != nil {
		return false
	}

	if i < 1920 || i > 2002 {
		return false
	}

	return true
}

/*
	checkIyr (Issue Year) - four digits; at least 2010 and at most 2020.
*/
func (p Passport) checkIyr() bool {
	if len(*p.iyr) != 4 {
		return false
	}
	i, err := strconv.Atoi(*p.iyr)
	if err != nil {
		return false
	}

	if i < 2010 || i > 2020 {
		return false
	}

	return true
}

/*
	checkEyr (Expiration Year) - four digits; at least 2020 and at most 2030.
*/
func (p Passport) checkEyr() bool {
	if len(*p.eyr) != 4 {
		return false
	}
	i, err := strconv.Atoi(*p.eyr)
	if err != nil {
		return false
	}

	if i < 2020 || i > 2030 {
		return false
	}

	return true
}

/*
	checkHgt (Height) - a number followed by either cm or in:
		- If cm, the number must be at least 150 and at most 193.
		- If in, the number must be at least 59 and at most 76.
*/
func (p Passport) checkHgt() bool {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(*p.hgt, -1)
	if len(numbers) == 0 {
		return false
	}
	d, err := strconv.Atoi(numbers[0])
	if err != nil {
		return false
	}
	if strings.Contains(*p.hgt, "cm") {
		if d < 150 || d > 193 {
			return false
		}
	} else if strings.Contains(*p.hgt, "in") {
		if d < 59 || d > 76 {
			return false
		}
	} else {
		return false
	}

	return true
}

/*
	checkHcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
*/
func (p Passport) checkHcl() bool {
	color := *p.hcl
	match, _ := regexp.MatchString("#[0-9a-f]{6}", color)

	if !match {
		return false
	}

	return true
}

/*
	checkEcl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
*/
func (p Passport) checkEcl() bool {
	color := *p.ecl

	if color == "amb" {
		return true
	}
	if color == "blu" {
		return true
	}
	if color == "brn" {
		return true
	}
	if color == "gry" {
		return true
	}
	if color == "grn" {
		return true
	}
	if color == "hzl" {
		return true
	}
	if color == "oth" {
		return true
	}

	return false
}

/*
	checkPid (Passport ID) - a nine-digit number, including leading zeroes.
*/
func (p Passport) checkPid() bool {
	pid := *p.pid
	if len(pid) > 9 {
		return false
	}
	match, _ := regexp.MatchString("[0-9]{9}", pid)

	if !match {
		return false
	}

	return true
}

/*	readInput
	Read the filename and returns a list.
*/
func readInput(filename string) ([]string, error) {
	filePath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passports []string

	var tempString string

	for scanner.Scan() {
		thisLine := scanner.Text()

		tempString = tempString + " " + thisLine

		if thisLine == ""{
			passports = append(passports, tempString)
			tempString = ""
		}
	}

	passports = append(passports, tempString)
	return passports, nil
}

/*	passportProcessing
	Tidies up the Passport data into a Passport struct.
*/
func passportProcessing(pass string) (Passport, error) {
	p := Passport{}

	words := strings.Split(pass, " ")
	for _, w := range words {
		if w == "" {
			continue
		}
		key := &strings.Split(w, ":")[1]

		if strings.Contains(w, "byr") {
			p.byr = key
			continue
		}

		if strings.Contains(w, "iyr") {
			p.iyr = key
			continue
		}

		if strings.Contains(w, "eyr") {
			p.eyr = key
			continue
		}

		if strings.Contains(w, "hgt") {
			p.hgt = key
			continue
		}

		if strings.Contains(w, "hcl") {
			p.hcl = key
			continue
		}

		if strings.Contains(w, "ecl") {
			p.ecl = key
			continue
		}

		if strings.Contains(w, "pid") {
			p.pid = key
			continue
		}

		if strings.Contains(w, "cid") {
			p.cid = key
			continue
		}
	}

	return p, nil
}

/*	passportChecker
	Checks if Passport is valid.
*/
func passportChecker(input []string) (int, error) {
	passportCount := 0

	for _, value := range input {
		passportStruct, err := passportProcessing(value)
		if err != nil {
			return 0, err
		}

		if passportStruct.byr == nil {
			continue
		} else if !passportStruct.checkByr() {
			continue
		}
		if passportStruct.iyr == nil {
			continue
		} else if !passportStruct.checkIyr() {
			continue
		}
		if passportStruct.eyr == nil {
			continue
		} else if !passportStruct.checkEyr() {
			continue
		}
		if passportStruct.hgt == nil {
			continue
		} else if !passportStruct.checkHgt() {
			continue
		}
		if passportStruct.hcl == nil {
			continue
		} else if !passportStruct.checkHcl() {
			continue
		}
		if passportStruct.ecl == nil {
			continue
		} else if !passportStruct.checkEcl() {
			continue
		}
		if passportStruct.pid == nil {
			continue
		} else if !passportStruct.checkPid() {
			continue
		}

		passportCount++
	}

	return passportCount, nil
}

func main() {
	passports, err := readInput("input-baphy.txt")
	if err != nil {
		log.Fatalln(err)
	}

	validPassports, err := passportChecker(passports)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(validPassports)
}
