package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type record struct {
	springs string
	nums    []int
	qs      int
}

func (r record) isValid(attempt string) bool {
	var result []int
	count := 0

	for _, char := range attempt {
		if char == '#' {
			count++
		} else if count > 0 {
			result = append(result, count)
			count = 0
		}
	}

	// If the sequence ends with '#', add the last count to the result
	if count > 0 {
		result = append(result, count)
	}

	b := reflect.DeepEqual(result, r.nums)
	return b
}

func intToBinaryWithLeadingZeroes(num int64, size int) string {
	binary := fmt.Sprintf("%0"+fmt.Sprintf("%d", size)+"b", num)
	return binary
}

func main() {
	//fmt.Printf(intToBinaryWithLeadingZeroes(5, 10))
	records := getStuff("day12/input.txt")
	//fmt.Printf("%+v", records)
	replacer := strings.NewReplacer("0", ".", "1", "#")

	count := 0
	for _, r := range records {
		var i int64
		//fmt.Printf("record %+v tries: \n ", r)
		for i = 0; i < pow2(r.qs); i++ {
			toWriteInto := replacer.Replace(intToBinaryWithLeadingZeroes(i, r.qs))
			attempt := replaceQuestionMarks(toWriteInto, r.springs)
			if r.isValid(attempt) {
				//fmt.Printf("success: %s\n", attempt)
				count++
			} else {
				//fmt.Printf("failure: %s\n", attempt)
			}
		}
		//fmt.Print("\n\n\n")
	}
	fmt.Printf("The answer is %d", count)
}

func replaceQuestionMarks(s1, s2 string) string {
	// Convert s1 to []rune for easy iteration
	s1Runes := []rune(s1)

	// Use a counter to keep track of the index in s1
	counter := 0

	// Convert s2 to []rune so we can modify it
	s2Runes := []rune(s2)

	// Loop through s2 and replace '?' with characters from s1
	for i := range s2Runes {
		if s2Runes[i] == '?' && counter < len(s1Runes) {
			s2Runes[i] = s1Runes[counter]
			counter++
		}
	}

	return string(s2Runes)
}

func pow2(n int) int64 {
	var r int64 = 1
	for i := 0; i < n; i++ {
		r *= 2
	}
	return r
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
