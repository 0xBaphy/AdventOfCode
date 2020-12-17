package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

/*	readInput
	Read the filename and returns a list.
*/
func readInput(filename string) ([]int, error) {
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

	var numbers []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, i)
	}

	return numbers, nil
}

/*	checkEntries
	Finds the two entries that sum 2020 and then multiplies the two numbers together and returns the product.
*/
func checkEntries(entries []int) (int, error) {
	for x, y := 0, 1; x < len(entries)-2; {
		if entries[x]+entries[y] == 2020 {
			fmt.Printf("The values are %d and %d\n", entries[x], entries[y])

			return entries[x] * entries[y], nil
		}

		if y == len(entries)-1 {
			x++
			y = x + 1

			continue
		}

		y++
	}

	return 0, errors.New("couldn't find the required values")
}

/*	checkEntriesPartTwo
	Finds the three entries that sum 2020 and then multiplies the three numbers together and returns the product.
*/
func checkEntriesPartTwo(entries []int) (int, error) {
	for x, y, z := 0, 1, 2; x < len(entries)-3; {
		if entries[x]+entries[y]+entries[z] == 2020 {
			fmt.Printf("The values are %d, %d and %d\n", entries[x], entries[y], entries[z])

			return entries[x] * entries[y] * entries[z], nil
		}

		if y == len(entries)-2 && z == len(entries)-1 {
			x++
			y = x + 1
			z = y + 1

			continue
		}

		if z == len(entries)-1 {
			y++
			z = y + 1
			continue
		}

		z++
	}

	return 0, errors.New("couldn't find the required values")
}

func main() {
	numbers, err := readInput("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	answer, err := checkEntries(numbers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part one:", answer, "\n-------")

	answerTwo, err := checkEntriesPartTwo(numbers)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Part two:", answerTwo)
}
