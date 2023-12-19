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
	grid, _ := readIntFile("day17/testinput.txt")
	wayTooBig := len(grid) * len(grid[0]) * 10

	dist := createGrid(len(grid), len(grid[0]), wayTooBig)
	dist[0][0] = pointWithDir{
		p:     point{0, 0},
		dir:   "",
		steps: 0,
		dist:  0,
	}

	q := createQ(len(grid), len(grid[0]))

	for len(q) > 0 {
		cur := getCur(q, dist)
		delete(q, cur)
		neighbors := getNeighbors(dist[cur.i][cur.j], len(grid), len(grid[0]))
		for _, n := range neighbors {
			if q[n.p] {
				newDist := dist[cur.i][cur.j].dist + grid[n.p.i][n.p.j]
				if newDist < dist[n.p.i][n.p.j].dist {
					dist[n.p.i][n.p.j] = pointWithDir{
						p:     point{n.p.i, n.p.j},
						dir:   n.dir,
						steps: n.steps,
						dist:  newDist,
					}
				}
			}
		}

	}

	fmt.Printf("The answer is: %d", dist[len(dist)-1][len(dist[0])-1].dist)
}

func getNeighbors(pwd pointWithDir, maxI, maxJ int) []pointWithDir {
	p := pwd.p

	r := make([]pointWithDir, 0)
	if p.i > 0 && (pwd.dir != "N" || pwd.steps < 3) {
		steps := 0
		if pwd.dir == "N" {
			steps = pwd.steps + 1
		}
		r = append(r, pointWithDir{
			p:     point{p.i - 1, p.j},
			dir:   "N",
			steps: steps,
		})
	}
	if p.i < maxI-1 && (pwd.dir != "S" || pwd.steps < 3) {
		steps := 0
		if pwd.dir == "S" {
			steps = pwd.steps + 1
		}
		r = append(r, pointWithDir{
			p:     point{p.i + 1, p.j},
			dir:   "S",
			steps: steps,
		})
	}
	if p.j > 0 && (pwd.dir != "W" || pwd.steps < 3) {
		steps := 0
		if pwd.dir == "W" {
			steps = pwd.steps + 1
		}
		r = append(r, pointWithDir{
			p:     point{p.i, p.i - 1},
			dir:   "S",
			steps: steps,
		})
	}
	if p.j < maxJ-1 && (pwd.dir != "E" || pwd.steps < 3) {
		steps := 0
		if pwd.dir == "E" {
			steps = pwd.steps + 1
		}
		r = append(r, pointWithDir{
			p:     point{p.i, p.j + 1},
			dir:   "E",
			steps: steps,
		})
	}
	return r
}

func getCur(q map[point]bool, dist [][]pointWithDir) point {
	curDist := 10000000
	var cur point
	for k, v := range q { // bug here - something weird happening with same distances potentially?
		if v && dist[k.i][k.j].dist < curDist {
			curDist = dist[k.i][k.j].dist
			cur = k
		}
	}
	return cur
}

func createQ(x, y int) map[point]bool {
	q := make(map[point]bool)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			q[point{i, j}] = true
		}
	}
	return q
}

func createEmptyPrev(i, j int) [][]*point {
	// Create the outer slice (rows)
	result := make([][]*point, i)

	// Create each row and initialize each element to nil
	for x := 0; x < i; x++ {
		result[x] = make([]*point, j)
	}

	return result
}

type pointWithDir struct {
	p     point
	dir   string
	steps int
	dist  int
}

type point struct {
	i int
	j int
}

func createGrid(m, n, v int) [][]pointWithDir {
	// Initialize the 2D slice with dimensions m by n
	grid := make([][]pointWithDir, m)
	for i := range grid {
		grid[i] = make([]pointWithDir, n)
	}

	// Set all values in the grid to v
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = pointWithDir{
				p:     point{i, j},
				dir:   "",
				steps: 0,
				dist:  v,
			}
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
