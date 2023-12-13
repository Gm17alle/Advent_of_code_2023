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
}

func contains(s string, c rune) bool {
	for _, r := range s {
		if r == c {
			return true
		}
	}
	return false
}

type pair struct {
	nI int
	sI int
}

func main() {
	records := getStuff("day12/input.txt")
	//fmt.Printf("%+v", records)

	count := 0
	for i, r := range records {
		if i == 18 {
			fmt.Printf("Record %d\n: ", i)
		}
		cache := make(map[pair]int)
		nums, springs := r.nums, r.springs
		var gcR func(nI, sI int) int
		gcR = func(nI, sI int) int {
			p := pair{nI, sI}
			if v, e := cache[p]; e {
				return v
			}
			if sI >= len(springs) { //If we've checked past the last val (this is fucking weird but it works idk)
				if nI == len(nums) { //Check the end if we've gotten all the way through
					return 1
				}
				return 0
			}

			curCount := 0
			if springs[sI] != '#' { // If it's not a #, we should treat it like a dot first
				curCount += gcR(nI, sI+1)
			}
			if springs[sI] != '.' && nI < len(nums) { // If it's not a '.', we should try to treat it like a #

				//See if we can have a run of num[nI] #s here
				sEnd := sI + nums[nI]
				// in bounds               no dots (so it can be a run)            next one isn't a # or doesn't exist
				if sEnd <= len(springs) && !contains(springs[sI:sEnd], '.') && (sEnd == len(springs) || springs[sEnd] != '#') {
					curCount += gcR(nI+1, sEnd+1)
				}
			}
			cache[p] = curCount
			return curCount
		}
		c := gcR(0, 0)
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

	nums := make([]int, 0)
	bigNums := make([]int, 0)
	for _, v := range ints {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
		bigNums = append(bigNums, n)
	}
	bigSprings := springs
	for i := 0; i < 4; i++ {
		bigSprings = bigSprings + "?" + springs
		bigNums = append(bigNums, nums...)
	}
	return record{bigSprings, bigNums}
}
