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

	/*
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var allGroups []group
	var thisGroup group
	for _, value := range lines {
		if value != "" {	
			thisLineMap := make(map[rune]int)

			for _, char := range value {
				thisLineMap[char]++
			}

			thisGroup.answers = thisLineMap
			thisGroup.people++
		}			
		
		if value == "" {
			allGroups = append(allGroups, thisGroup)
			thisGroup = group{}
		}
	}
	*/

	var allGroups []group
	var thisGroup group
	for scanner.Scan() {
		thisLine := scanner.Text()
		if thisLine == "" {
			allGroups = append(allGroups, thisGroup)
			thisGroup = group{}
			continue
		}

		thisLineMap := make(map[rune]int)
		for _, char := range thisLine {
			thisLineMap[char]++
		}

		thisGroup.answers = thisLineMap
		thisGroup.people++
	}

	return allGroups, nil
}

/*	answerCounter
	Counts the answered questions on each group and returns the sum value of all questions answered.
*/
func answerCounter(input []string) (int, error) {
	answers := make(map[rune]int)
	counter := 0
	for _, value := range input {
		for _, char := range value {
			answers[char]++
		}

		counter += len(answers)
		answers = make(map[rune]int)
	}

	return counter, nil
}

func answerCounterTwo(input []group) (int, error) {
	return 0, nil
}

func main() {
	groups, err := readInput("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	for _, thisGroup := range groups {
		fmt.Println(thisGroup.people)
	}
}
