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

func seatFinder(input []string)(int, error) {
// BFFFBBFRRR
// 	0-127 rows
// F means to take the lower half, keeping rows 0 through 63.
// B means to take the upper half, keeping rows 32 through 63.
// [][]bool{
//  {0,0,0,0,0,0,0,0,0}
//  {0,0,0,0,0,0,0,0,0}
//  {0,0,0,0,0,0,0,0,0}
//  {0,0,0,0,0,0,0,0,0}
//}

    plane := make([][]bool, 8)
    for i, _ := range plane {
        plane[i] = make([]bool, 128)
    }

    for _, line := range input {
        row := 128
        upperOrLower := 'U'
        for _, char := range line {
            if char == 'B' {
                row
			}
        }
    }






	return 0, nil
}

func main() {
	boardingPasses, err := readInput("2020/Day-05/src/test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(boardingPasses)
}
