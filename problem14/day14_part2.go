package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	grid := getSlices("day14/testinput.txt")

	seen := make(map[int][]int)
	count := 0
	foundCycle := false
	cycleHypothesis := make(map[int][]int)
	startCycle := 0

	for numTimes := 0; numTimes < 1000000000; numTimes++ {
		if numTimes%10000000 == 0 {
			fmt.Printf("On try %d", numTimes)
		}

		//NORTH
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

		//WEST
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				c := grid[i][j]
				if c == 'O' {
					curJ := j
					for curJ-1 > -1 && grid[i][curJ-1] == '.' {
						grid[i][curJ] = '.'
						grid[i][curJ-1] = 'O'
						curJ--
					}
				}
			}
		}

		//SOUTH
		for i := len(grid) - 1; i > -1; i-- {
			for j := 0; j < len(grid[0]); j++ {
				c := grid[i][j]
				if c == 'O' {
					curI := i
					for curI+1 < len(grid) && grid[curI+1][j] == '.' {
						grid[curI][j] = '.'
						grid[curI+1][j] = 'O'
						curI++
					}
				}
			}
		}

		// EAST
		for i := 0; i < len(grid); i++ {
			for j := len(grid[0]) - 1; j > -1; j-- {
				c := grid[i][j]
				if c == 'O' {
					curJ := j
					for curJ+1 < len(grid[0]) && grid[i][curJ+1] == '.' {
						grid[i][curJ] = '.'
						grid[i][curJ+1] = 'O'
						curJ++
					}
				}
			}
		}
		count = 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 'O' {
					count += len(grid) - i
				}
			}
		}

		if _, e := seen[count]; e && len(seen[count]) > 1000 && !foundCycle {
			foundCycle = true
			cycleHypothesis[count] = append(cycleHypothesis[count], numTimes)
			startCycle = count
		} else if foundCycle {
			if count == startCycle && len(cycleHypothesis[count]) == 10 {
				for k, v := range cycleHypothesis {
					fmt.Printf("Load: %d     happens at : %v\n", k, v)
				}
				break
			}
			cycleHypothesis[count] = append(cycleHypothesis[count], numTimes)
		}

		seen[count] = append(seen[count], numTimes)

	}

	var keyValuePairs []struct {
		key   int
		value int
	}

	// Populate the slice with elements from the map
	for k, v := range cycleHypothesis {
		keyValuePairs = append(keyValuePairs, struct {
			key   int
			value int
		}{key: k, value: v[0]})
	}

	// Define a custom sorting function based on the first element of the []int
	sort.SliceStable(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].value < keyValuePairs[j].value
	})

	// Print the sorted key-value pairs
	fmt.Println("Sorted key-value pairs")
	for _, pair := range keyValuePairs {
		fmt.Printf("Key: %d, Value: %v\n", pair.key, pair.value)
	}

	fmt.Printf("The answer is probably %d", keyValuePairs[((1000000000-(keyValuePairs[0].value+1))%len(keyValuePairs))].key)
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
