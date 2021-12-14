package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 14/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n\n")
	polymerT := inputS[0]
	rulesS := strings.Split(inputS[1], "\n")

	rules := make(map[string]string)
	for _, val := range rulesS {
		parts := strings.Split(val, " -> ")
		rules[parts[0]] = parts[1]
	}
	sol1 := part1(polymerT, rules, 10)
	sol2 := part1(polymerT, rules, 40)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(polymerT string, rules map[string]string, repeat int) int {

	values := make(map[string]int)
	valuesLetters := make(map[string]int)
	for i := 0; i < len(polymerT)-1; i++ {
		pair := string(polymerT[i]) + string(polymerT[i+1])
		if _, ok := values[pair]; ok {
			values[pair]++
		} else {
			values[pair] = 1
		}
		letter := string(polymerT[i])
		if _, ok := valuesLetters[letter]; ok {
			valuesLetters[letter]++
		} else {
			valuesLetters[letter] = 1
		}
	}

	if _, ok := valuesLetters[string(polymerT[len(polymerT)-1])]; ok {
		valuesLetters[string(polymerT[len(polymerT)-1])]++
	} else {
		valuesLetters[string(polymerT[len(polymerT)-1])] = 1
	}

	for i := 0; i < repeat; i++ {
		values = applyRules(rules, values, valuesLetters)
		//for key, value := range values {
		//fmt.Println(key, value)
		//}
		//for key, value := range valuesLetters {
		//fmt.Println(key, value)
		//}
	}

	maxN := 0
	minN := valuesLetters[string(polymerT[0])]
	//var max, min string
	for _, value := range valuesLetters {
		if value > maxN {
			maxN = value
			//max = key
		}
		if value < minN {
			minN = value
			//min = key
		}
	}

	return maxN - minN
}

func applyRules(rules map[string]string, values, valuesLetters map[string]int) map[string]int {
	newValues := make(map[string]int)
	for key, value := range values {
		nextChar := rules[key]
		if _, ok := valuesLetters[nextChar]; ok {
			valuesLetters[nextChar] += value
		} else {
			valuesLetters[nextChar] = value
		}

		newOne := string(key[0]) + nextChar
		newTwo := nextChar + string(key[1])

		if _, ok := newValues[newOne]; ok {
			newValues[newOne] += value
		} else {
			newValues[newOne] = value
		}

		if _, ok := newValues[newTwo]; ok {
			newValues[newTwo] += value
		} else {
			newValues[newTwo] = value
		}
	}

	return newValues
}
