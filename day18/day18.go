package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
#####
#   #
#   #
#  ##
#  ##
#####
*/

func main() {
	edges, _ := readEdgesFromFile("day18/input.txt")
	//fmt.Println(edges)
	lowI, highI, lowJ, highJ := 0, 0, 0, 0
	cur := point{0, 0}
	perimeter := make(map[point]bool)
	perimeter[cur] = true
	points1 := make([]point, 1)
	points1[0] = point{0, 0}
	// Draw the border
	for _, e := range edges {
		if e.dir == "U" {
			for i := cur.i; i <= cur.i+e.dist; i++ {
				perimeter[point{i, cur.j}] = true
			}
			cur = point{cur.i + e.dist, cur.j}
			if cur.i > highI {
				highI = cur.i
			}
		} else if e.dir == "L" {
			for j := cur.j; j >= cur.j-e.dist; j-- {
				perimeter[point{cur.i, j}] = true
			}
			cur = point{cur.i, cur.j - e.dist}
			if cur.j < lowJ {
				lowJ = cur.j
			}
		} else if e.dir == "D" {
			for i := cur.i; i >= cur.i-e.dist; i-- {
				perimeter[point{i, cur.j}] = true
			}
			cur = point{cur.i - e.dist, cur.j}
			if cur.i < lowI {
				lowI = cur.i
			}
		} else if e.dir == "R" {
			for j := cur.j; j <= cur.j+e.dist; j++ {
				perimeter[point{cur.i, j}] = true
			}
			cur = point{cur.i, cur.j + e.dist}
			if cur.j > highJ {
				highJ = cur.j
			}
		} else {
			panic("unexpected direction")
		}
		points1 = append(points1, cur)
	}

	//floodFill(perimeter, point{1, -1})
	A1 := shoelace(points1)
	b1 := len(perimeter)
	fmt.Printf("The p1 answer is: %d\n", A1-b1/2+1+b1)

	p := point{0, 0}
	allThePoints := make([]point, 1)
	allThePoints[0] = p

	b2 := 0
	for _, e := range edges {
		color := e.color[2:7]
		dir := e.color[7:8]
		cDist, _ := strconv.ParseInt(color, 16, 0)
		dist := int(cDist)

		//fmt.Printf("%v%v", color, dir)
		if dir == "0" {
			p = point{p.i, p.j + dist}
		} else if dir == "1" {
			p = point{p.i - dist, p.j}
		} else if dir == "2" {
			p = point{p.i, p.j - dist}
		} else if dir == "3" {
			p = point{p.i + dist, p.j}
		} else {
			panic("panik")
		}
		allThePoints = append(allThePoints, p)
		b2 += dist
	}
	A2 := shoelace(allThePoints)

	fmt.Printf("The p2 answer is: %d", A2-b2/2+1+b2)
}

func floodFill(points map[point]bool, p point) {
	points[p] = true
	surrounding := []point{{p.i, p.j + 1}, {p.i, p.j - 1}, {p.i + 1, p.j}, {p.i - 1, p.j}}
	for _, np := range surrounding {
		if !points[np] {
			floodFill(points, np)
		}
	}
}

func shoelace(points []point) int {
	area := 0
	for i := 0; i < len(points)-1; i++ {
		area += points[i].i*points[i+1].j - points[i+1].i*points[i].j
	}
	return area / 2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type edge struct {
	dir   string
	dist  int
	color string
}

var left point = point{0, -1}
var right point = point{0, 1}
var up point = point{1, 0}
var down point = point{-1, 0}

type point struct {
	i, j int
}

func (p point) add(p2 point) point {
	return point{p.i + p2.i, p.j + p2.j}
}

// Function to read file and populate edge struct
func readEdgesFromFile(filename string) ([]edge, error) {
	var edges []edge

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 3 {
			return nil, fmt.Errorf("invalid format in line: %s", line)
		}

		intValue, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("invalid integer in line: %s", line)
		}

		edge := edge{
			dir:   fields[0],
			dist:  intValue,
			color: fields[2],
		}

		edges = append(edges, edge)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return edges, nil
}
