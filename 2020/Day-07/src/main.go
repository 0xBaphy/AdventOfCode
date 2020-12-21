package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type bags struct {
	bag map[string]map[string]int
}

/*	readInput
	Reads the filename and returns a list.
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

/*	sorter
	Sorts the provided bag rules into a bag struct.
*/
func sorter(input []string) (bags, error) {
	var tempBag bags
	parentBag := make(map[string]map[string]int)

	for _, thisBag := range input {
		childBag := make(map[string]int)

		split := strings.Split(thisBag, " ")

		if split[4] == "no" {
			parentBag[split[0]+" "+split[1]] = childBag
			tempBag.bag = parentBag
			continue
		}

		for index := 4; index < len(split); index++ {
			if strings.Contains(split[index], "bag") {
				color := split[index-2] + " " + split[index-1]
				num, err := strconv.Atoi(split[index-3])
				if err != nil {
					log.Fatalln(err)
				}

				childBag[color] = num
			}

		}

		parentBag[split[0]+" "+split[1]] = childBag
	}

	tempBag.bag = parentBag
	return tempBag, nil
}

/*	counterOne
	Counts how many bags can eventually contain at least one target bag.
*/
func counterOne(input bags, target string) int {
	bagCount := make(map[string]bool)
	bagCount[target] = true

start:
	for parentColor := range input.bag {
		for childColor := range input.bag[parentColor] {
			if bagCount[childColor] && !bagCount[parentColor] {
				bagCount[parentColor] = true
				goto start
			}
		}
	}

	return len(bagCount) - 1
}

/*	counterTwo
	Counts how many individual bags can fit inside the provided target bag.
*/
func counterTwo(input bags, target string) int {
	bagCount := 0

	for color, num := range input.bag[target] {
		bagCount += num * counterTwo(input, color)

	}

	return bagCount + 1
}

func main() {
	input, err := readInput("baphy-input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sorted, err := sorter(input)
	if err != nil {
		log.Fatalln(err)
	}

	one := counterOne(sorted, "shiny gold")
	fmt.Println(one)

	two := counterTwo(sorted, "shiny gold")
	fmt.Println(two - 1)
}
