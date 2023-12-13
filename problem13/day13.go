package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arrays := read2DCharArrayFromFile("day13/input.txt")
	fancyCount := 0
	//numSmudges := 0 // part 1
	numSmudges := 1 // part 2
	for i, arr := range arrays {
		rowR := findRowReflection(arr, numSmudges)
		colR := findColReflection(arr, numSmudges)
		if colR != -1 {
			fancyCount += colR
		} else if rowR != -1 {
			fancyCount += rowR
		} else {
			fmt.Printf("Problem with arr %d %v", i, arr)
		}
	}
	fmt.Printf("The answer is: %d", fancyCount)
}

func findColReflection(arr [][]rune, numSmudges int) int {
	for i := 0; i < len(arr)-1; i++ {
		smudge := 0
		top := i
		bottom := i + 1
		for top >= 0 && bottom < len(arr) && smudge < numSmudges+1 {
			for j := 0; j < len(arr[0]); j++ {
				if arr[top][j] != arr[bottom][j] {
					smudge++
				}
			}
			top--
			bottom++
		}
		if smudge == numSmudges {
			return (i + 1) * 100
		}
	}
	return -1
}

func findRowReflection(arr [][]rune, numSmudges int) int {
	for j := 0; j < len(arr[0])-1; j++ {
		smudge := 0
		left := j
		right := j + 1
		for left >= 0 && right < len(arr[0]) && smudge < numSmudges+1 {
			for i := 0; i < len(arr); i++ {
				if arr[i][left] != arr[i][right] {
					smudge++
				}
			}
			left--
			right++
		}
		if smudge == numSmudges {
			return j + 1
		}
	}
	return -1
}

func read2DCharArrayFromFile(filename string) [][][]rune {
	content, _ := os.ReadFile(filename)

	lines := strings.Split(string(content), "\n")
	var arrays [][][]rune
	var currentArray [][]rune

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			arrays = append(arrays, currentArray)
			currentArray = nil
		} else {
			currentArray = append(currentArray, []rune(line))
		}
	}

	if len(currentArray) > 0 {
		arrays = append(arrays, currentArray)
	}

	return arrays
}
