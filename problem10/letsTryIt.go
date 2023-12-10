package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x    int
	y    int
	from string
}

type pair struct {
	x int
	y int
}

func (c coord) ctp() pair {
	return pair{c.x, c.y}
}

// might have some dupes to part 2 - copied over
// to preserve
func part1() {
	visited := make(map[pair]bool)
	grid := getSlices("input.txt")
	sLoc := getSLoc(grid)
	_, startLocs := getSConnected(grid, sLoc)
	visited[sLoc.ctp()] = true
	count := 1
	b1, b2 := startLocs[0], startLocs[1]
	for !visited[b1.ctp()] && !visited[b2.ctp()] {
		count++
		visited[b1.ctp()] = true
		visited[b2.ctp()] = true
		b1 = getNext(grid, b1)
		b2 = getNext(grid, b2)
	}
	if b1.x != b2.x || b2.y != b2.y {
		count--
	}
	fmt.Printf("The count is : %d", count)

}

func part2() {

}

func main() {
	visited := make(map[pair]bool)
	grid := getSlices("input.txt")
	sLoc := getSLoc(grid)
	r, startLocs := getSConnected(grid, sLoc)
	grid[sLoc.x][sLoc.y] = r
	visited[sLoc.ctp()] = true
	b1, b2 := startLocs[0], startLocs[1]
	for !visited[b1.ctp()] && !visited[b2.ctp()] {
		visited[b1.ctp()] = true
		visited[b2.ctp()] = true
		b1 = getNext(grid, b1)
		b2 = getNext(grid, b2)
	}

	area := 0
	for i, l := range grid {
		//isColinear := false
		for j, _ := range l {
			intersects := 0
			if visited[pair{i, j}] {
				continue
			}

			newX, newY := i-1, j-1
			for newX >= 0 && newY >= 0 {
				p := pair{newX, newY}
				curChar := grid[newX][newY]
				if visited[p] {
					if curChar == '7' || curChar == 'L' {
						intersects += 2 // pass through it twice (or not at all)
					} else {
						intersects++
					}
				}
				newX--
				newY--
			}
			if intersects%2 == 1 {
				grid[i][j] = '@'
				area++
			}
		}
	}
	printGrid(grid)
	fmt.Printf("The area is %d\n", area)
}

// idea 2:
//if visited[p] && isColinear && c == '-' {
//continue
//// do nothing
//}
//isColinear = false
//if visited[p] {
//intersects++
//if c == '-' { // colinear bits aren't treated as intersects
//intersects--
//isColinear = true
//}
//} else if c == '.' && intersects%2 == 1 { // a spot that can be filled in with an odd number of intersects
//grid[i][j] = '@'
//area++
//}

// old idea:
//if c == '.' {
//count := 0 //count of intersections down
//for v := i + 1; v < len(grid); v++ {
//if visited[pair{v, j}] && grid[v][j] != '|' {
//count++
//}
//}
//if count%2 == 1 {
//grid[i][j] = '@'
//area++
//}
//}

func printGrid(grid [][]rune) {
	for _, l := range grid {
		for _, c := range l {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func getNext(grid [][]rune, b coord) coord {
	var n coord
	c := grid[b.x][b.y]
	f := b.from
	if c == 'F' {
		if f == "S" {
			n = coord{b.x, b.y + 1, "W"}
		}
		if f == "E" {
			n = coord{b.x + 1, b.y, "N"}
		}
	} else if c == '|' {
		if f == "N" {
			n = coord{b.x + 1, b.y, "N"}
		}
		if f == "S" {
			n = coord{b.x - 1, b.y, "S"}
		}
	} else if c == '7' {
		if f == "W" {
			n = coord{b.x + 1, b.y, "N"}
		}
		if f == "S" {
			n = coord{b.x, b.y - 1, "E"}
		}
	} else if c == 'L' {
		if f == "N" {
			n = coord{b.x, b.y + 1, "W"}
		}
		if f == "E" {
			n = coord{b.x - 1, b.y, "S"}
		}
	} else if c == '-' {
		if f == "W" {
			n = coord{b.x, b.y + 1, "W"}
		}
		if f == "E" {
			n = coord{b.x, b.y - 1, "E"}
		}
	} else if c == 'J' {
		if f == "N" {
			n = coord{b.x, b.y - 1, "E"}
		}
		if f == "W" {
			n = coord{b.x - 1, b.y, "S"}
		}
	}
	return n
}

// F|7L-J
func getSConnected(grid [][]rune, c coord) (rune, []coord) {
	var toReturn []coord
	right, left, up, down := false, false, false, false
	if grid[c.x][c.y+1] == '-' || grid[c.x][c.y+1] == 'J' || grid[c.x][c.y+1] == '7' {
		toReturn = append(toReturn, coord{x: c.x, y: c.y + 1, from: "W"})
		right = true
	}
	if grid[c.x][c.y-1] == '-' || grid[c.x][c.y-1] == 'L' || grid[c.x][c.y-1] == 'F' {
		toReturn = append(toReturn, coord{x: c.x, y: c.y - 1, from: "E"})
		left = true
	}
	if grid[c.x-1][c.y] == '|' || grid[c.x-1][c.y] == '7' || grid[c.x-1][c.y] == 'F' {
		toReturn = append(toReturn, coord{x: c.x - 1, y: c.y, from: "S"})
		up = true
	}
	if grid[c.x+1][c.y] == '|' || grid[c.x+1][c.y] == 'J' || grid[c.x+1][c.y] == 'L' {
		toReturn = append(toReturn, coord{x: c.x + 1, y: c.y, from: "N"})
		down = true
	}
	var returnRune rune
	if right && left {
		returnRune = '-'
	} else if right && up {
		returnRune = 'L'
	} else if right && down {
		returnRune = 'F'
	} else if left && up {
		returnRune = 'J'
	} else if left && down {
		returnRune = '7'
	} else if up && down {
		returnRune = '|'
	}
	return returnRune, toReturn
}

func getSLoc(grid [][]rune) coord {
	for i, xs := range grid {
		for j, y := range xs {
			if y == 'S' {
				return coord{x: i, y: j}
			}
		}
	}
	return coord{}
}

func getSlices(path string) [][]rune {
	file, _ := os.Open(path)

	runeSlice := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		runeSlice = append(runeSlice, line)
	}

	return runeSlice
}
