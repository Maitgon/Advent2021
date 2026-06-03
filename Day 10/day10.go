package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type Stack []byte

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str byte) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return ' ', false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 10/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(strs []string) int {
	pointsDist := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	points := 0
	for _, str := range strs {
		if ok, letter := isCorrupted(str, pairs); ok {
			points += pointsDist[letter]
		}
	}
	return points
}

func isOpen(str byte) bool {
	return str == '(' || str == '[' || str == '{' || str == '<'
}

func incomplete(code string) Stack {
	var stack Stack
	for i := range code {
		if isOpen(code[i]) {
			stack.Push(code[i])
		} else {
			stack.Pop()
		}
	}
	return stack
}

func isCorrupted(code string, pairs map[byte]byte) (bool, byte) {
	var stack Stack
	for i := range code {
		if isOpen(code[i]) {
			stack.Push(code[i])
		} else {
			top, _ := stack.Pop()
			if pairs[code[i]] != top {
				return true, code[i]
			}
		}
	}
	return false, ' '
}

func part2(strs []string) int {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	var incompletes []string
	for _, str := range strs {
		if ok, _ := isCorrupted(str, pairs); !ok {
			incompletes = append(incompletes, str)
		}
	}
	pointsDist := map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	pointsAll := make([]int, len(incompletes))
	for i, str := range incompletes {
		rest := incomplete(str)

		points := 0
		for !rest.IsEmpty() {
			val, _ := rest.Pop()
			points = 5*points + pointsDist[val]
		}
		pointsAll[i] = points
	}

	sort.Ints(pointsAll)

	return pointsAll[len(pointsAll)/2]
}
