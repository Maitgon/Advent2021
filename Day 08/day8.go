package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 08/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	var codes, outputs [][]string
	for _, val := range inputS {
		parts := strings.Split(val, " | ")
		codes = append(codes, strings.Split(parts[0], " "))
		outputs = append(outputs, strings.Split(parts[1], " "))
	}

	sol1 := part1(outputs)
	sol2 := part2(codes, outputs)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(outputs [][]string) int {
	count := 0
	for _, output := range outputs {
		for _, num := range output {
			if len(num) == 2 || len(num) == 3 || len(num) == 4 || len(num) == 7 {
				count++
			}
		}
	}
	return count
}

func part2(codes, outputs [][]string) int {
	sol := 0
	for i, code := range codes {
		mapeo := mapping(code)

		num1 := mapeo[SortString(outputs[i][0])]
		num2 := mapeo[SortString(outputs[i][1])]
		num3 := mapeo[SortString(outputs[i][2])]
		num4 := mapeo[SortString(outputs[i][3])]

		sol += num1*1000 + num2*100 + num3*10 + num4
	}
	return sol
}

func mapping(code []string) map[string]int {

	mapeo := make(map[string]int)

	// Hacemos sort a los codigos ya que luego los outputs tienen distinto orden,
	// pero siguen teniendo las mismas letras

	// Buscamos 1, 4, 7 y 8.

	code1 := getByLength(code, 2)
	code = delete(code, code1)
	mapeo[SortString(code1)] = 1

	code4 := getByLength(code, 4)
	code = delete(code, code4)
	mapeo[SortString(code4)] = 4

	code7 := getByLength(code, 3)
	code = delete(code, code7)
	mapeo[SortString(code7)] = 7

	code8 := getByLength(code, 7)
	code = delete(code, code8)
	mapeo[SortString(code8)] = 8

	// Ahora buscamos 9 0 y 6

	code9 := getByLengthCommon(code, 6, 4, code4)
	code = delete(code, code9)
	mapeo[SortString(code9)] = 9

	code0 := getByLengthCommon(code, 6, 3, code7)
	code = delete(code, code0)
	mapeo[SortString(code0)] = 0

	code6 := getByLength(code, 6)
	code = delete(code, code6)
	mapeo[SortString(code6)] = 6

	// Por último buscamos 3, 5 y 2

	code3 := getByLengthCommon(code, 5, 3, code7)
	code = delete(code, code3)
	mapeo[SortString(code3)] = 3

	code5 := getByLengthCommon(code, 5, 3, code4)
	code = delete(code, code5)
	mapeo[SortString(code5)] = 5

	code2 := getByLength(code, 5)
	mapeo[SortString(code2)] = 2

	return mapeo

}

func getByLength(code []string, length int) string {
	for _, str := range code {
		if len(str) == length {
			return str
		}
	}
	return ""
}

func getByLengthCommon(code []string, length, repeat int, comm string) string {
	for _, str := range code {
		if len(str) == length && common(str, comm) == repeat {
			return str
		}
	}
	return ""
}

func delete(code []string, obj string) []string {
	var sol []string
	for _, str := range code {
		if str != obj {
			sol = append(sol, str)
		}
	}
	return sol
}

func common(str1, str2 string) int {
	count := 0
	for _, c1 := range str1 {
		for _, c2 := range str2 {
			if c1 == c2 {
				count++
				break
			}
		}
	}
	return count
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
