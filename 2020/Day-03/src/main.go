package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

/*	replaceAtIndex
	Replaces a given rune in a string using an index.
*/
func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

/*	treeCounter
	Goes trough the map and counts the trees that are in the way of the toboggan.
*/
func treeCounter(treeMap []string, right int, down int) (int, error) {
	treeCount := 0
	for i := 0; i < down; i++ {
		fmt.Println(treeMap[i])
	}

	iterCount := 1
	for index := down; index < len(treeMap); {
		value := treeMap[index]

		lineIndex := (right * (iterCount)) % len(value)

		if value[lineIndex] == '#' {
			fmt.Println(replaceAtIndex(value, 'X', lineIndex))
			treeCount++
		} else {
			fmt.Println(replaceAtIndex(value, 'O', lineIndex))
		}

		index += down
		iterCount++
	}

	return treeCount, nil
}

func main() {
	treeMap, err := readInput("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	forest := 1
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		trees, err := treeCounter(treeMap, slope[0], slope[1])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(trees)
		forest *= trees
	}

	fmt.Println(forest)
}
