// Day 1.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 1/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	var vals []int16
	for _, value := range input {
		val, _ := strconv.ParseInt(value, 10, 16)
		vals = append(vals, int16(val))
	}
	values := vals[:len(vals)-1]

	sol1 := part1(values)

	sol2 := part2(values)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(values []int16) int16 {
	count := 0
	for i := 0; i < len(values) - 1; i++ {
		if values[i] < values[i+1] {
			count++
		}
	}
	return int16(count)
}

func part2(values []int16) int16 {
	count := 0
	for i := 0; i < len(values) - 3; i++ {
		if values[i] + values[i+1] + values[i+2] <
		   values[i+1] + values[i+2] + values[i+3] {
			   count++
		   }
	}
	return int16(count)
}