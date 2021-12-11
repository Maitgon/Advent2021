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
	bs, err := ioutil.ReadFile("./Go/Day 11/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	matrix := make([][]int8, len(inputS))
	for i, val := range inputS {
		aux := make([]int8, len(inputS[0]))
		for j, num := range val {
			aux[j] = int8(num - '0')
		}
		matrix[i] = aux
	}

	sol1, sol2 := part1y2(matrix)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1y2(matrix [][]int8) (int, int) {
	flashes := 0
	i := 0
	for ; i < 100; i++ {
		oneStep(matrix, &flashes)
	}
	sol1 := flashes
	for ; !allZeroes(matrix); i++ {
		oneStep(matrix, &flashes)
	}
	return sol1, i
}

func oneStep(matrix [][]int8, flashes *int) {

	// Sum 1 to all positions
	// There shouldn't be any position with 9
	for i := range matrix {
		for j := range matrix {
			matrix[i][j]++
		}
	}

	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] == 10 {
				explosion(matrix, i, j, flashes)
			}
		}
	}
}

func explosion(matrix [][]int8, i, j int, flashes *int) {
	matrix[i][j] = 0

	*flashes++
	for i1 := -1; i1 <= 1; i1++ {
		for j1 := -1; j1 <= 1; j1++ {
			if !(0 <= i+i1 && i+i1 < len(matrix) &&
				0 <= j+j1 && j+j1 < len(matrix[0])) {
				continue
			} else if matrix[i+i1][j+j1] == 9 {
				explosion(matrix, i+i1, j+j1, flashes)
			} else if matrix[i+i1][j+j1] > 0 &&
				matrix[i+i1][j+j1] <= 9 {
				matrix[i+i1][j+j1]++
			}
		}
	}
}

func allZeroes(matrix [][]int8) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
