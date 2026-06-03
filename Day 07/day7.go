package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 07/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), ",")

	var input []int
	for _, value := range inputS {
		val, _ := strconv.ParseInt(value, 10, 16)
		input = append(input, int(val))
	}

	sol1, sol2 := part1y2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1y2(input []int) (int, int) {
	sol1 := 0
	sol2 := 0
	mean2 := Mean(input)
	median1 := Median(input)
	for _, val := range input {
		sol1 += Abs(median1 - val)
		sol2 += sumToVal(Abs(mean2 - val))
	}
	return sol1, sol2
}

func sumToVal(x int) int {
	return (x * (x + 1)) / 2
}

func Mean(input []int) int {
	sum := 0
	for _, val := range input {
		sum += val
	}
	return int(math.Round(float64(sum) / float64(len(input))))
}

func Median(input []int) int {
	sort.Ints(input)
	long := len(input)
	if long%2 == 1 {
		return input[long/2]
	}

	return (input[long/2] + input[long/2-1]) / 2
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
