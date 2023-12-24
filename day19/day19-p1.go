package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type xmas []int

func (w xmas) sumChristmasThisis() int {
	return w[0] + w[1] + w[2] + w[3]
}

func main() {
	s, _ := readFileToString("day19/input.txt")
	inputs := strings.Split(s, "\r\n\r\n")
	m := readMaps(strings.Split(inputs[0], "\r\n"))

	christmases := newXmasArrayFromStringArray(strings.Split(inputs[1], "\r\n"))
	sum := 0
	for _, c := range christmases {
		curWorkflow := "in"
		curIndex := 0

		for true {
			curFunc := m[curWorkflow][curIndex]
			v := getConditionResult(curFunc, c)
			if v == nil {
				curIndex++
			} else if *v == "A" {
				sum += c.sumChristmasThisis()
				break
			} else if *v == "R" {
				break
			} else {
				curWorkflow = *v
				curIndex = 0

			}
		}
	}

	fmt.Printf("%d", sum)
}

func newXmasArrayFromStringArray(input []string) []xmas {
	var xmasArray []xmas

	for _, str := range input {
		xmasStruct := newXmasFromString(str)
		xmasArray = append(xmasArray, xmasStruct)
	}

	return xmasArray
}

func newXmasFromString(input string) xmas {
	values := strings.Split(input, ",")
	result := make([]int, 4)

	for _, val := range values {
		parts := strings.Split(strings.Trim(val, "{}"), "=")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}

		switch key {
		case "x":
			result[0] = value
		case "m":
			result[1] = value
		case "a":
			result[2] = value
		case "s":
			result[3] = value
		}
	}

	return result
}

func getConditionResult(condition string, arr xmas) *string {
	if !strings.ContainsRune(condition, ':') {
		return &condition
	}

	indexMap := map[byte]int{'x': 0, 'm': 1, 'a': 2, 's': 3}
	index := indexMap[condition[0]]
	var operator rune
	if strings.ContainsRune(condition, '>') {
		operator = '>'
	} else {
		operator = '<'
	}
	resultString := strings.Split(condition, ":")[1]
	vs := condition[strings.Index(condition, string(operator))+1 : strings.Index(condition, ":")]
	value, _ := strconv.Atoi(vs)

	if index >= len(arr) {
		return nil
	}

	switch operator {
	case '<':
		if arr[index] < value {
			return &resultString
		}
	case '>':
		if arr[index] > value {
			return &resultString
		}
	}
	return nil

}

func readMaps(s []string) map[string][]string {
	m := make(map[string][]string)
	for _, line := range s {
		onCurly := strings.Split(line, "{")
		prefix := onCurly[0][:len(onCurly[0])]
		funcs := strings.Split(strings.Replace(onCurly[1], "}", "", -1), ",")
		thisLinesFuncs := make([]string, 0)
		for _, fun := range funcs {
			thisLinesFuncs = append(thisLinesFuncs, fun)
		}
		m[prefix] = thisLinesFuncs
	}
	return m
}

func readFileToString(filePath string) (string, error) {
	// Read the contents of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the file content bytes to a string
	fileContent := string(content)
	return fileContent, nil
}
