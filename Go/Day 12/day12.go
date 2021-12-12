package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return " ", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) IsIn(str string) bool {
	for _, elem := range *s {
		if elem == str {
			return true
		}
	}
	return false
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 12/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	mapping := make(map[string][]string)
	for _, val := range inputS {
		vals := strings.Split(val, "-")
		val1, val2 := vals[0], vals[1]
		if con, ok := mapping[val1]; ok && val1 != "end" && val2 != "start" {
			mapping[val1] = append(con, val2)
		} else if val1 != "end" && val2 != "start" {
			mapping[val1] = []string{val2}
		}
		if con, ok := mapping[val2]; ok && val2 != "end" && val1 != "start" {
			mapping[val2] = append(con, val1)
		} else if val2 != "end" && val1 != "start" {
			mapping[val2] = []string{val1}
		}
	}

	sol1 := part1(mapping)
	sol2 := part2(mapping)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(mapping map[string][]string) int {
	count := 0
	var visited Stack
	visit(mapping, "start", &count, visited)
	return count
}

func visit(mapping map[string][]string, actual string, count *int, visited Stack) {

	// Visitamos el nodo si es pequeño
	if isLowerCase(actual[0]) {
		visited.Push(actual)
	}

	// Visitamos los siguientes
	next := mapping[actual]
	for _, nodeNext := range next {
		if nodeNext == "end" { // If next node is end, then just add 1 to count
			*count++
		} else if !visited.IsIn(nodeNext) { // Don't visit next node if it was visited
			visit(mapping, nodeNext, count, visited)
		}
	}

	// Desvisitamos el nodo si es pequeño
	if isLowerCase(actual[0]) {
		visited.Pop()
	}
}

func part2(mapping map[string][]string) int {
	count := 0
	var visitedOnce, visitedTwice Stack
	visit2(mapping, "start", &count, visitedOnce, visitedTwice)
	return count
}

func visit2(mapping map[string][]string, actual string, count *int, visitedOnce, visitedTwice Stack) {

	// Visitamos el nodo si es pequeño
	if isLowerCase(actual[0]) && !visitedOnce.IsIn(actual) {
		visitedOnce.Push(actual)
	} else if isLowerCase(actual[0]) {
		visitedTwice.Push(actual)
	}

	// Visitamos los siguientes
	next := mapping[actual]
	for _, nodeNext := range next {
		if nodeNext == "end" { // If next node is end, then just add 1 to count
			*count++
		} else if !visitedOnce.IsIn(nodeNext) ||
			!visitedTwice.IsIn(nodeNext) &&
				visitedTwice.IsEmpty() { // Don't visit next node if it was visited
			visit2(mapping, nodeNext, count, visitedOnce, visitedTwice)
		}
	}

	// Desvisitamos el nodo si es pequeño
	if isLowerCase(actual[0]) && !visitedOnce.IsIn(actual) {
		visitedOnce.Pop()
	} else if isLowerCase(actual[0]) {
		visitedTwice.Pop()
	}
}

func isLowerCase(letter byte) bool {
	return 'a' <= letter && letter <= 'z'
}
