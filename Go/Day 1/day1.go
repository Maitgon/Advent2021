// Day 1.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 1/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	var vals []int64
	for _, value := range input {
		val, _ := strconv.ParseInt(value, 10, 16)
		vals = append(vals, val)
	}
	values := vals[:len(vals)-1]

	sol1 := part1(values)
	fmt.Println("The solution to part 1 is: ", sol1)

	sol2 := part2(values)
	fmt.Println("The solution to part 2 is: ", sol2)

}

func part1(values []int64) int64 {
	count := 0
	for i := 0; i < len(values) - 1; i++ {
		if values[i] < values[i+1] {
			count++
		}
	}
	return int64(count)
}

func part2(values []int64) int64 {
	count := 0
	for i := 0; i < len(values) - 3; i++ {
		if values[i] + values[i+1] + values[i+2] <
		   values[i+1] + values[i+2] + values[i+3] {
			   count++
		   }
	}
	return int64(count)
}