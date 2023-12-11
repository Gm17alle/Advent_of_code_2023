package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func (p point) distance(o point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	grid := getSlices("day11/input.txt")
	xSet, ySet := getToExpand(grid)
	//points := getOffsetPoints(grid, xSet, ySet, 1) // part 1
	points := getOffsetPoints(grid, xSet, ySet, 999999) // Part 2
	distance := getDistances(points)

	fmt.Printf("Distance %n: ", distance)
}

func getDistances(points []point) int {
	dist := 0
	for i, p := range points {
		for j := i; j < len(points); j++ {
			dist += p.distance(points[j])
		}
	}
	return dist
}

func getOffsetPoints(grid [][]rune, xS, yS map[int]bool, offset int) []point {
	points := make([]point, 0)
	realX := 0
	for i, l := range grid {
		realY := 0
		for j, v := range l {
			if v == '#' {
				points = append(points, point{realX, realY})
			}
			if yS[j] {
				realY += offset
			}
			realY++
		}
		if xS[i] {
			realX += offset
		}
		realX++
	}
	return points
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

func getToExpand(grid [][]rune) (map[int]bool, map[int]bool) {
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)

	for i, _ := range grid {
		isEmpty := true
		for j, _ := range grid[i] {
			if grid[i][j] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			xSet[i] = true
		}
	}
	for j, _ := range grid[0] {
		isEmpty := true
		for i, _ := range grid {
			if grid[i][j] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			ySet[j] = true
		}
	}

	return xSet, ySet
}
