package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type (
	person struct {
		answer map[rune]bool
	}

	group struct {
		answers []person
	}
)

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

	var allGroups []group
	var thisGroup group
	for scanner.Scan() {
		thisLine := scanner.Text()
		if thisLine == "" {
			allGroups = append(allGroups, thisGroup)
			thisGroup = group{}
			continue
		}
		thisLineMap := make(map[rune]bool)
		for _, char := range thisLine {
			thisLineMap[char] = true
		}
		thisPerson := person{
			answer: thisLineMap,
		}
		thisGroup.answers = append(thisGroup.answers, thisPerson)

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
	var sum = 0
	for _, thisGroup := range input {
		answers := make(map[rune]bool)
		for _, thisPerson := range thisGroup.answers {
			for thisAnswer := range thisPerson.answer {
				answers[thisAnswer] = true
			}
		}
		fmt.Println(len(answers))
		sum += len(answers)
	}
	return sum, nil
}

func main() {
	questions, err := readInput("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	uwu, err := answerCounterTwo(questions)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(uwu)
}
