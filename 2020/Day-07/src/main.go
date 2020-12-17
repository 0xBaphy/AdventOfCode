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

type bag struct {
	color    string
	contains map[string]int
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

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func bagSorter(input []string) ([]bag, error) {
	var bags []bag
	regex := regexp.MustCompile(`([a-z].+?)(?:bags )|([0-9])(.+?)(?:bags|bag)`)
	reNum := regexp.MustCompile(`\w[0-9]|[0-9]`)
	reBag := regexp.MustCompile(`(bags)|(bag)`)

	for _, thisBag := range input {
		matches := regex.FindAllString(thisBag, -1)
		var b bag
		nestedBag := make(map[string]int)
		for index, match := range matches {
			bagNum := reNum.FindString(match)
			match = reBag.ReplaceAllString(match, "")
			match = reNum.ReplaceAllString(match, "")
			match = strings.TrimSpace(match)

			if index == 0 {
				b.color = match

			} else {
				num, err := strconv.Atoi(bagNum)
				if err != nil {
					log.Fatalln(err)
				}

				nestedBag[match] = num
			}
		}
		b.contains = nestedBag
		bags = append(bags, b)
	}

	return bags, nil
}

func goldBagCounter(input []bag) (int, error) {
	bags := make(map[string]bool)
	prev := 0
	for index := 0; index < len(input); index++ {
		for color := range input[index].contains {
			if color == "shiny gold" {
				bags[input[index].color] = true
			}
			for mapColor := range bags {
				if color == mapColor {
					bags[input[index].color] = true
				}
			}
		}

		if index == len(input)-1 {
			index = 0
			if prev == len(bags) {
				break
			}
			prev = len(bags)
		}
	}

	return len(bags), nil
}

func main() {
	bags, err := readInput("baphy-input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sortedBags, err := bagSorter(bags)
	if err != nil {
		log.Fatalln(err)
	}

	answer, err := goldBagCounter(sortedBags)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(answer)
}
