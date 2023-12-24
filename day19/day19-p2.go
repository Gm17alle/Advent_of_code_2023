package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type xmas []int

type minmax struct {
	min int
	max int
}

func (m minmax) bisectGreaterThan(mid int) (minmax, minmax) {
	newLow := minmax{m.min, mid}
	newHigh := minmax{mid + 1, m.max}
	return newLow, newHigh
}

func (m minmax) bisectLessThan(mid int) (minmax, minmax) {
	newLow := minmax{m.min, mid - 1}
	newHigh := minmax{mid, m.max}
	return newLow, newHigh
}

type possibilities struct {
	val map[string]minmax
}

type node struct {
	isDirectMapping bool
	mapsTo          string
	operator        rune
	xmasc           string
	val             int
}

func (p possibilities) getNum() int {
	product := 1
	for _, v := range p.val {
		product = product * (v.max - v.min + 1)
	}
	return product
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str, _ := readFileToString("day19/input.txt")
	inputs := strings.Split(str, "\r\n\r\n")
	m := readMaps(strings.Split(inputs[0], "\r\n"))

	//sum := 0
	//
	//curWorkflow := "in"
	//curIndex := 0
	//
	var x func(cur string, index int, p possibilities) int
	x = func(cur string, index int, p possibilities) int {
		if cur == "A" {
			t := p.getNum()
			if t < 0 {
				fmt.Printf("")
			}
			fmt.Printf("We are at cur %s index %d with total %d and possibilities: %+v\n", cur, index, t, p)
			return t
		} else if cur == "R" {
			return 0
		}
		if index >= len(m[cur]) {
			return 0
		}
		n := m[cur][index]
		total := 0
		if n.isDirectMapping {
			return x(n.mapsTo, 0, p)
		} else if n.operator == '>' {
			newLow, newHigh := p.val[n.xmasc].bisectGreaterThan(n.val)
			p1 := newPos(p, n.xmasc, newLow)
			total += x(cur, index+1, p1)
			p2 := newPos(p, n.xmasc, newHigh)
			total += x(n.mapsTo, 0, p2)
		} else if n.operator == '<' {
			newLow, newHigh := p.val[n.xmasc].bisectLessThan(n.val)
			p1 := newPos(p, n.xmasc, newLow)
			total += x(n.mapsTo, 0, p1)
			p2 := newPos(p, n.xmasc, newHigh)
			total += x(cur, index+1, p2)
		} else {
			panic(":(")
		}
		return total
	}

	v := make(map[string]minmax)
	v["x"] = minmax{1, 4000}
	v["m"] = minmax{1, 4000}
	v["a"] = minmax{1, 4000}
	v["s"] = minmax{1, 4000}
	p := possibilities{v}

	ans := x("in", 0, p)

	fmt.Printf("%d", ans)
}

func newPos(p possibilities, s string, n minmax) possibilities {
	r := make(map[string]minmax)
	for k, v := range p.val {
		if k == s {
			r[k] = n
		} else {
			r[k] = v
		}
	}
	return possibilities{r}
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

func readMaps(s []string) map[string][]node {
	m := make(map[string][]node)
	for _, line := range s {
		onCurly := strings.Split(line, "{")
		prefix := onCurly[0][:len(onCurly[0])]
		funcs := strings.Split(strings.Replace(onCurly[1], "}", "", -1), ",")
		thisLinesFuncs := make([]node, 0)
		for _, fun := range funcs {
			n := getNodeFromFun(fun)
			thisLinesFuncs = append(thisLinesFuncs, n)
		}
		m[prefix] = thisLinesFuncs
	}
	return m
}

func getNodeFromFun(condition string) node {
	if !strings.ContainsRune(condition, ':') {
		return node{
			isDirectMapping: true,
			mapsTo:          condition,
		}
	}

	thisisntaruneforsomefreakingreason := condition[0:1]
	var operator rune
	if strings.ContainsRune(condition, '>') {
		operator = '>'
	} else {
		operator = '<'
	}
	resultString := strings.Split(condition, ":")[1]
	vs := condition[strings.Index(condition, string(operator))+1 : strings.Index(condition, ":")]
	value, _ := strconv.Atoi(vs)

	return node{
		isDirectMapping: false,
		mapsTo:          resultString,
		operator:        operator,
		xmasc:           thisisntaruneforsomefreakingreason,
		val:             value,
	}
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
