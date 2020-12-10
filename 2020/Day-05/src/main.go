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

/*
seatFinder will find the largest seat ID in the input array
*/
func seatFinder(input []string) (int, [][]int, error) {
	plane := make([][]int, 8)
	for i := range plane {
		plane[i] = make([]int, 128)
	}

	maxID := 0
	for _, line := range input {
		seatID := 0
		rowEnd := 127
		rowStart := 0
		colEnd := 7
		colStart := 0
		colSelect := 0
		rowSelect := 0
		for index, char := range line {
			if index < 7 {
				if index == 6 {
					if char == 'B' {
						rowSelect = rowEnd
					}
					if char == 'F' {
						rowSelect = rowStart
					}

					continue
				}

				if char == 'B' {
					rowStart = rowEnd - ((rowEnd - rowStart) / 2)
				}
				if char == 'F' {
					rowEnd = rowEnd - ((rowEnd - rowStart) / 2) - 1
				}
			}

			if index == len(line)-1 {
				if char == 'R' {
					colSelect = colEnd
				}
				if char == 'L' {
					colSelect = colStart
				}

				break
			}

			if char == 'R' {
				colStart = colEnd - ((colEnd - colStart) / 2)
			}
			if char == 'L' {
				colEnd = colEnd - ((colEnd - colStart) / 2) - 1
			}
		}
		seatID = (rowSelect * 8) + colSelect
		if seatID > maxID {
			maxID = seatID
		}
		plane[colSelect][rowSelect] = 1

	}

	return maxID, plane, nil
}

func main() {
	boardingPasses, err := readInput("acido-input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	seatID, plane, err := seatFinder(boardingPasses)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(seatID)
	for colInx := range plane {
		for rowInx := range plane[colInx] {
			if plane[colInx][rowInx] == 0 {
				plane[colInx][rowInx] = (rowInx * 8) + colInx
			}
		}
		fmt.Println(plane[colInx])
	}
}
