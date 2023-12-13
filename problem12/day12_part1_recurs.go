package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	springs string
	nums    []int
	qs      int
}

func getCountRecurs(nums []int, springs string, nI, sI int) int {
	if sI >= len(springs) { //If we've checked past the last val (this is fucking weird but it works idk)
		if nI == len(nums) { //Check the end if we've gotten all the way through
			return 1
		}
		return 0
	}

	count := 0
	if springs[sI] != '#' { // If it's not a #, we should treat it like a dot first
		count += getCountRecurs(nums, springs, nI, sI+1)
	}
	if springs[sI] != '.' && nI < len(nums) { // If it's not a '.', we should try to treat it like a #

		//See if we can have a run of num[nI] #s here
		sEnd := sI + nums[nI]
		// in bounds               no dots (so it can be a run)            next one isn't a # or doesn't exist
		if sEnd <= len(springs) && !contains(springs[sI:sEnd], '.') && (sEnd == len(springs) || springs[sEnd] != '#') {
			count += getCountRecurs(nums, springs, nI+1, sEnd+1)
		}
	}
	return count
}

func contains(s string, c rune) bool {
	for _, r := range s {
		if r == c {
			return true
		}
	}
	return false
}

func main() {
	records := getStuff("day12/input.txt")
	//fmt.Printf("%+v", records)

	count := 0
	for _, r := range records {
		fmt.Printf("record %+v ", r)
		c := getCountRecurs(r.nums, r.springs, 0, 0)
		fmt.Printf("count: %d\n", c)
		count += c
	}
	fmt.Printf("The answer is %d", count)
}

func getStuff(path string) []record {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error %+v", err)
	}

	records := make([]record, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		records = append(records, toRecord(scanner.Text()))
	}

	return records
}

func toRecord(s string) record {
	springs := strings.Split(s, " ")[0]
	ints := strings.Split(strings.Split(s, " ")[1], ",")
	numQs := 0
	for _, c := range springs {
		if c == '?' {
			numQs++
		}
	}

	nums := make([]int, 0)
	for _, v := range ints {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	return record{springs, nums, numQs}
}
