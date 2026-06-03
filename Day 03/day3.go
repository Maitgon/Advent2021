package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"time"
)

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 03/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")
	input = input[:len(input)-1]

	sol1 := part1(input)
	sol2 := part2(input)


	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(input []string) int64 {
	var x string
	for i := 0; i < len(input[0]); i++ {
		count0 := 0
		count1 := 0
		for j := 0; j < len(input); j++ {
			if input[j][i] == '0' {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			x += "0"
		} else {
			x += "1"
		}
	}
	gammaRate, _ := strconv.ParseInt(x, 2, 16)
	epsilonRate := int64(math.Pow(2,float64(len(x)))) - gammaRate- 1
	return gammaRate * epsilonRate
}

func part2(input []string) int64 {
	gammaRate, _ := strconv.ParseInt(findGamma(input), 2, 16)
	epsilonRate, _ := strconv.ParseInt(findEpsilon(input), 2, 16)
	return gammaRate * epsilonRate
}

func findGamma(input []string) string {
	var gammaRate string
	for i := 0; i < len(input[0]); i++ {
		if len(input) == 1 {
			break
		}
		count0 := 0
		count1 := 0
		for j := 0; j < len(input); j++ {
			if input[j][i] == '0' {
				count0++
			} else {
				count1++
			}
		}
		var target byte
		if count0 > count1 {
			target = '0'
		} else {
			target = '1'
		}
		filtered := []string{}
		for j := 0; j < len(input); j++ {
			if input[j][i] == target {
				if len(input) < 5 {
				}
				filtered = append(filtered, input[j])
			}
		}
		input = filtered
	}

	if len(input) == 1 {
		gammaRate = input[0]
	}
	
	return gammaRate
}

func findEpsilon(input []string) string {

	var epsilonRate string
	for i := 0; i < len(input[0]); i++ {
		if len(input) == 1 {
			epsilonRate = input[0]
			break
		}
		count0 := 0
		count1 := 0
		for j := 0; j < len(input); j++ {
			if input[j][i] == '0' {
				count0++
			} else {
				count1++
			}
		}
		var target byte
		if count1 < count0 {
			target = '1'
		} else {
			target = '0'
		}
		filtered := []string{}
		for j := 0; j < len(input); j++ {
			if input[j][i] == target {
				filtered = append(filtered, input[j])
			}
		}
		input = filtered
	}

	return epsilonRate
}