package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	m, _ := readIntFile("day17/myinput.txt")
	wayTooBig := len(m) * len(m[0]) * 10
	dp := createGrid(len(m), len(m[0]), wayTooBig)
	dp[0][0] = 0

	ans := bfs(m, dp)
	fmt.Printf("%d", ans)
}

func createGrid(m, n, v int) [][]int {
	// Initialize the 2D slice with dimensions m by n
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	// Set all values in the grid to v
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = v
		}
	}

	return grid
}

type iter struct {
	i          int
	j          int
	total      int
	travelling string
	stepsInDir int
}

func bfs(grid, dp [][]int) int {
	q := make([]iter, 2)
	q[0] = iter{0, 1, 0, "E", 1}
	q[1] = iter{1, 0, 0, "S", 1}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.i < 0 || cur.i >= len(grid) || cur.j < 0 || cur.j >= len(grid[0]) || cur.stepsInDir > 3 {
			continue
		}
		newTot := grid[cur.i][cur.j] + cur.total
		if cur.i == 0 && cur.j == 7 {
			fmt.Print("")
		}
		if newTot > dp[cur.i][cur.j] {
			continue // don't rewrite anything if we've already found a cheaper route to get here
		}
		dp[cur.i][cur.j] = newTot
		curDir := cur.travelling
		var a, b, c iter
		if curDir == "E" {
			a = iter{cur.i - 1, cur.j, newTot, "N", 1}
			b = iter{cur.i + 1, cur.j, newTot, "S", 1}
			c = iter{cur.i, cur.j + 1, newTot, "E", cur.stepsInDir + 1}
		} else if curDir == "S" {
			a = iter{cur.i, cur.j + 1, newTot, "E", 1}
			b = iter{cur.i, cur.j - 1, newTot, "W", 1}
			c = iter{cur.i + 1, cur.j, newTot, "S", cur.stepsInDir + 1}
		} else if curDir == "W" {
			a = iter{cur.i - 1, cur.j, newTot, "N", 1}
			b = iter{cur.i + 1, cur.j, newTot, "S", 1}
			c = iter{cur.i, cur.j - 1, newTot, "W", cur.stepsInDir + 1}
		} else {
			a = iter{cur.i, cur.j + 1, newTot, "E", 1}
			b = iter{cur.i, cur.j - 1, newTot, "W", 1}
			c = iter{cur.i - 1, cur.j, newTot, "N", cur.stepsInDir + 1}
		}
		q = append(q, a, b, c)
	}
	printAligned(dp)
	return dp[len(dp)-1][len(dp[0])-1]
}

func printAligned(grid [][]int) {
	// Find the maximum width for each column
	maxWidth := make([]int, len(grid[0]))
	for _, row := range grid {
		for i, num := range row {
			width := int(math.Log10(float64(num))) + 1 // Calculate the number of digits
			if width > maxWidth[i] {
				maxWidth[i] = width
			}
		}
	}

	// Print the aligned grid
	for _, row := range grid {
		for i, num := range row {
			fmt.Printf("%*d", maxWidth[i]+1, num)
		}
		fmt.Println()
	}
}

func readIntFile(fileLoc string) ([][]int, error) {
	// Open the file
	file, err := os.Open(fileLoc)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize the 2D slice
	var result [][]int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into individual digits
		digits := strings.Split(line, "")

		// Initialize a slice to store the integers from the current line
		var intLine []int

		// Convert each digit to an integer and append it to intLine
		for _, digitStr := range digits {
			digit, err := strconv.Atoi(digitStr)
			if err != nil {
				return nil, err
			}
			intLine = append(intLine, digit)
		}

		// Append the line to the result slice
		result = append(result, intLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
