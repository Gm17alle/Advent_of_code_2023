package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := getSlices("day14/input.txt")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			c := grid[i][j]
			if c == 'O' {
				curI := i
				for curI-1 > -1 && grid[curI-1][j] == '.' {
					grid[curI][j] = '.'
					grid[curI-1][j] = 'O'
					curI--
				}
			}
		}
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'O' {
				count += len(grid) - i
			}
		}
	}
	fmt.Printf("The total load is: %d ", count)
}

func getSlices(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error %+v", err)
	}

	runeSlice := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		runeSlice = append(runeSlice, line)
	}

	return runeSlice
}
