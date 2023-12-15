package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type commandInfo struct {
	Command string
	Lens    int
}

// node represents a node in the linked list
type node struct {
	Data commandInfo
	Next *node
}

// commandInfoLinkedList represents a linked list of commandInfo nodes
type commandInfoLinkedList struct {
	Head *node
}

// AddOrReplaceNode adds a new node to the linked list
func (ll *commandInfoLinkedList) AddOrReplaceNode(data commandInfo) {
	newNode := &node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	if current.Data.Command == data.Command {
		current.Data.Lens = data.Lens
		return
	}
	for current.Next != nil {
		if current.Data.Command == data.Command {
			current.Data.Lens = data.Lens
			return
		}
		current = current.Next
	}
	if current.Data.Command == data.Command {
		current.Data.Lens = data.Lens
		return
	}
	current.Next = newNode
}

func (ll *commandInfoLinkedList) RemoveAll(command string) {
	if ll.Head == nil {
		return
	}

	// Removing nodes from the beginning if they match the command
	for ll.Head != nil && ll.Head.Data.Command == command {
		ll.Head = ll.Head.Next
	}

	current := ll.Head
	for current != nil && current.Next != nil {
		if current.Next.Data.Command == command {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

func main() {
	commands, _ := readFileAndSplitByComma("day15/input.txt")
	boxes := make([]commandInfoLinkedList, 256)
	for i, _ := range boxes {
		boxes[i] = commandInfoLinkedList{}
	}
	for _, entry := range commands {
		if entry[0:2] == "rn" {
			print("here we go")
		}
		if strings.ContainsRune(entry, '=') {
			command := strings.Split(entry, "=")[0]
			lens, _ := strconv.Atoi(strings.Split(entry, "=")[1])
			ci := commandInfo{command, lens}
			boxes[getBox(command)].AddOrReplaceNode(ci)
		} else {
			command := strings.Replace(entry, "-", "", -1)
			boxes[getBox(command)].RemoveAll(command)
		}
	}

	focusingPower := 0
	for i, v := range boxes {
		cur := v.Head
		count := 1
		for cur != nil {
			focusingPower += (i + 1) * count * cur.Data.Lens
			cur = cur.Next
			count++
		}
	}

	fmt.Printf("the answer is: %d", focusingPower)
}

func getBox(command string) int {
	cur := 0
	for _, r := range command {
		cur += int(r)
		cur *= 17
		cur %= 256
	}
	return cur
}

func readFileAndSplitByComma(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")
		content = append(content, words...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}
