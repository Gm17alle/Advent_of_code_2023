package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bricks, _ := readBricksFromFile("day22/input.txt")

	mX, mY, mZ := -1, -1, -1
	for _, b := range bricks {
		if max(b.start.x, b.end.x) > mX {
			mX = max(b.start.x, b.end.x)
		}
		if max(b.start.y, b.end.y) > mY {
			mY = max(b.start.y, b.end.y)
		}
		if max(b.start.z, b.end.z) > mZ {
			mZ = max(b.start.z, b.end.z)
		}
	}
	grid := make([][][]int, mX+1)
	for h := 0; h <= mX; h++ {
		xy := make([][]int, mY+1)
		for i := 0; i <= mY; i++ {
			y := make([]int, mZ+1)
			for j := 0; j <= mZ; j++ {
				y[j] = -1
			}
			xy[i] = y
		}
		grid[h] = xy
	}

	for i, b := range bricks {
		bricks[i].gridVal = i
		if b.start.x != b.end.x {
			for x := min(b.start.x, b.end.x); x <= max(b.start.x, b.end.x); x++ {
				grid[x][b.start.y][b.start.z] = i
			}
		} else if b.start.y != b.end.y {
			for y := min(b.start.y, b.end.y); y <= max(b.start.y, b.end.y); y++ {
				grid[b.start.x][y][b.start.z] = i
			}
		} else {
			for z := min(b.start.z, b.end.z); z <= max(b.start.z, b.end.z); z++ {
				grid[b.start.x][b.start.y][z] = i
			}
		}
	}

	// Sort bricks by Z
	sortByLowerZ(bricks)

	// Drop bricks by sorted Z order
	newBricks := make([]brick, 0)
	for _, b := range bricks {
		v := grid[b.minX()][b.minY()][b.minZ()]
		if b.start.z != b.end.z {
			x := b.minX()
			y := b.minY()
			zBottom := b.minZ()
			zTop := b.maxZ()
			for zBottom-1 > 0 && grid[b.minX()][b.minY()][zBottom-1] == -1 {
				grid[x][y][zBottom-1] = v
				grid[x][y][zTop] = -1
				zBottom--
				zTop--
				b.start.z = b.start.z - 1
				b.end.z = b.end.z - 1
			}
		} else {
			z := b.minZ()
			for z > 1 && isEmptyBelow(b, grid) {
				for i := b.minX(); i <= b.maxX(); i++ {
					for j := b.minY(); j <= b.maxY(); j++ {
						grid[i][j][z-1] = v
						grid[i][j][z] = -1
					}
				}
				z--
				b.start.z = b.start.z - 1
				b.end.z = b.end.z - 1
			}
		}
		newBricks = append(newBricks, b)
	}

	// See what bricks can be removed
	// Maybe make a map from brick to list of bricks supporting?

	supporting := make(map[int]map[int]bool)
	supportedBy := make(map[int]map[int]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < len(grid[0][0])-1; k++ {
				v := grid[i][j][k]
				above := grid[i][j][k+1]
				if grid[i][j][k] != -1 && grid[i][j][k+1] != -1 && v != above {
					if _, e := supporting[v]; !e {
						supporting[v] = make(map[int]bool)
					}
					supporting[v][above] = true

					if _, e := supportedBy[above]; !e {
						supportedBy[above] = make(map[int]bool)
					}
					supportedBy[above][v] = true
				}
			}
		}
	}

	count := 0
	necessaryBricks := make([]int, 0)
	for i := 0; i < len(bricks); i++ {
		if above, e := supporting[i]; !e {
			fmt.Printf("Brick %d not supporting anything\n", i)
			count++
		} else {
			ok := true
			for c, _ := range above {
				if len(supportedBy[c]) < 2 {
					ok = false
					break
				}
			}
			if ok {
				fmt.Printf("Brick %d is redundant \n", i)
				count++
			} else {
				fmt.Printf("Brick %d is necessary \n", i)
				necessaryBricks = append(necessaryBricks, i)
			}
		}
	}

	fmt.Printf("The answer is %d", count)

	//howManyWouldFall := 0
	//for _, i := range necessaryBricks {
	//	visited := make(map[int]bool)
	//	q := make([]int, 0)
	//	for v, _ := range supporting[i] {
	//		q = append(q, v)
	//	}
	//	for len(q) > 0 {
	//		curBrick := q[0]
	//		q = q[1:]
	//		if visited[curBrick] {
	//			continue
	//		}
	//
	//		howManyWouldFall++
	//		visited[curBrick] = true
	//		for k, _ := range supporting[curBrick] {
	//			if !visited[k] {
	//				q = append(q, k)
	//			}
	//		}
	//	}
	//}

	p2 := 0
	for _, curBrick := range newBricks {
		newNewBricks := make([]brick, 0)
		for _, b := range newBricks {
			if curBrick.gridVal != b.gridVal {
				newNewBricks = append(newNewBricks, b)
			}
		}
		newGrid := deepCopyWithReplacedValue(grid, curBrick.gridVal)
		//fmt.Printf("%+v%+v", newGrid, newNewBricks)
		p2 += countFallDowns(newNewBricks, newGrid)

	}
	fmt.Printf("\n Part 2: %d", p2)
}

func countFallDowns(bricks []brick, grid [][][]int) int {
	count := 0
	for _, b := range bricks {
		wentDown := false
		v := grid[b.minX()][b.minY()][b.minZ()]
		if b.start.z != b.end.z {
			x := b.minX()
			y := b.minY()
			zBottom := b.minZ()
			zTop := b.maxZ()
			for zBottom-1 > 0 && grid[b.minX()][b.minY()][zBottom-1] == -1 {
				wentDown = true
				grid[x][y][zBottom-1] = v
				grid[x][y][zTop] = -1
				zBottom--
				zTop--
				b.start.z = b.start.z - 1
				b.end.z = b.end.z - 1
			}
		} else {
			z := b.minZ()
			for z > 1 && isEmptyBelow(b, grid) {
				wentDown = true
				for i := b.minX(); i <= b.maxX(); i++ {
					for j := b.minY(); j <= b.maxY(); j++ {
						grid[i][j][z-1] = v
						grid[i][j][z] = -1
					}
				}
				z--
				b.start.z = b.start.z - 1
				b.end.z = b.end.z - 1
			}
		}
		if wentDown {
			count++
		}
	}
	return count
}

func deepCopyWithReplacedValue(arr [][][]int, toSetToNegativeOne int) [][][]int {
	copyArr := make([][][]int, len(arr))

	for i := range arr {
		copyArr[i] = make([][]int, len(arr[i]))
		for j := range arr[i] {
			copyArr[i][j] = make([]int, len(arr[i][j]))
			for k := range arr[i][j] {
				if arr[i][j][k] == toSetToNegativeOne {
					copyArr[i][j][k] = -1
				} else {
					copyArr[i][j][k] = arr[i][j][k]
				}
			}
		}
	}

	return copyArr
}

func isEmptyBelow(b brick, grid [][][]int) bool {
	z := b.minZ()
	for i := b.minX(); i <= b.maxX(); i++ {
		for j := b.minY(); j <= b.maxY(); j++ {
			if grid[i][j][z-1] != -1 {
				return false
			}
		}
	}
	return true
}

func sortByLowerZ(bricks []brick) {
	sort.Slice(bricks, func(i, j int) bool {
		if bricks[i].minZ() == bricks[j].minZ() {
			if bricks[i].minY() == bricks[j].minY() {
				return bricks[i].minX() < bricks[j].minX()
			}
			return bricks[i].minY() < bricks[j].minY()
		}
		return bricks[i].minZ() < bricks[j].minZ()
	})
}

type point struct {
	x int
	y int
	z int
}

type brick struct {
	gridVal int
	start   point
	end     point
}

func (b brick) minX() int {
	return min(b.start.x, b.end.x)
}

func (b brick) maxX() int {
	return max(b.start.x, b.end.x)
}

func (b brick) minY() int {
	return min(b.start.y, b.end.y)
}

func (b brick) maxY() int {
	return max(b.start.y, b.end.y)
}

func (b brick) minZ() int {
	return min(b.start.z, b.end.z)
}

func (b brick) maxZ() int {
	return max(b.start.z, b.end.z)
}

func parseLine(line string) (brick, error) {
	parts := strings.Split(line, "~")

	startPoint, err := parsePoint(parts[0])
	if err != nil {
		return brick{}, err
	}

	endPoint, err := parsePoint(parts[1])
	if err != nil {
		return brick{}, err
	}

	return brick{start: startPoint, end: endPoint}, nil
}

func parsePoint(s string) (point, error) {
	coords := strings.Split(s, ",")
	if len(coords) != 3 {
		return point{}, fmt.Errorf("invalid point format: %s", s)
	}

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return point{}, fmt.Errorf("failed to parse x coordinate: %v", err)
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return point{}, fmt.Errorf("failed to parse y coordinate: %v", err)
	}

	z, err := strconv.Atoi(coords[2])
	if err != nil {
		return point{}, fmt.Errorf("failed to parse z coordinate: %v", err)
	}

	return point{x: x, y: y, z: z}, nil
}

func readBricksFromFile(filePath string) ([]brick, error) {
	var bricks []brick

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		b, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		bricks = append(bricks, b)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bricks, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
