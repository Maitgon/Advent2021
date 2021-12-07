package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 06/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), ",")
	//inputS := strings.Split("3,4,3,1,2", ",")

	input := make([]int64, 9)
	for i := 0; i <= 8; i++ {
		input[i] = 0
	}
	for _, value := range inputS {
		val, _ := strconv.ParseInt(value, 10, 16)
		input[uint(val)] += 1
	}

	sol1, sol2 := part1y2(input, 80, 256)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func oneDay(input []int64) []int64 {
	result := make([]int64, 9)
	for _, i := range [7]uint{0, 1, 2, 3, 4, 5, 7} {
		result[i] = input[i+1]
	}
	result[6] = input[0] + input[7]
	result[8] = input[0]
	return result
}

func part1y2(input []int64, days1, days2 uint16) (int64, int64) {
	i := uint16(0)

	// Part 1
	for ; i < days1; i++ {
		input = oneDay(input)
	}
	sol1 := int64(0)
	for _, val := range input {
		sol1 += val
	}

	// Part 2
	for ; i < days2; i++ {
		input = oneDay(input)
	}
	sol2 := int64(0)
	for _, val := range input {
		sol2 += val
	}

	return sol1, sol2
}
