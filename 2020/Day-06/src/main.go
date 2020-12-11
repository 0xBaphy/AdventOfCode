package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type group struct {
	answers map[rune]int
	people  int
}

/*	readInput
	Read the filename and returns a list.
*/
func readInput(filename string) ([]group, error) {
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

	var allGroups []group
	var thisGroup group
	thisLineMap := make(map[rune]int)
	for index, line := range lines {

		if line != "" {
			for _, char := range line {
				thisLineMap[char]++
			}

			thisGroup.answers = thisLineMap
			thisGroup.people++
		}

		if line == "" || index == len(lines)-1 {
			allGroups = append(allGroups, thisGroup)
			thisGroup = group{}
			thisLineMap = make(map[rune]int)

			continue
		}
	}

	return allGroups, nil
}

/*	answerCounter
	Counts the answered questions on each group.
*/
func answerCounter(input []group) (int, error) {
	counter := 0
	for _, thisGroup := range input {
		for range thisGroup.answers {
			counter++
		}
	}

	return counter, nil
}

/*	answerCounterTwo
	Counts the answered questions on each group using the new instructions.
*/
func answerCounterTwo(input []group) (int, error) {
	counter := 0
	for _, thisGroup := range input {
		for _, answer := range thisGroup.answers {
			if thisGroup.people == answer {
				counter++
			}
		}
	}

	return counter, nil
}

func main() {
	groups, err := readInput("baphy-input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	one, err := answerCounter(groups)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(one)

	two, err := answerCounterTwo(groups)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(two)
}
