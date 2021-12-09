package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type intSeen struct {
	num   int
	visto bool
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 09/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	var matrix [][]intSeen
	for _, val := range inputS {
		var aux []intSeen
		for _, num := range val {
			aux = append(aux, intSeen{int(num - '0'), false})
		}
		matrix = append(matrix, aux)
	}

	sol1 := part1(matrix)
	sol2 := part2(matrix)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func isLowPoint(matrix [][]intSeen, i, j int) bool {
	cumplen := 0
	if i-1 < 0 || matrix[i-1][j].num > matrix[i][j].num {
		cumplen++
	}
	if i+1 >= len(matrix) || matrix[i+1][j].num > matrix[i][j].num {
		cumplen++
	}
	if j-1 < 0 || matrix[i][j-1].num > matrix[i][j].num {
		cumplen++
	}
	if j+1 >= len(matrix[0]) || matrix[i][j+1].num > matrix[i][j].num {
		cumplen++
	}
	return cumplen == 4
}

func part1(matrix [][]intSeen) int {
	risk := 0
	for i := range matrix {
		for j := range matrix[0] {
			if isLowPoint(matrix, i, j) {
				risk += 1 + matrix[i][j].num
			}
		}
	}
	return risk
}

func part2(matrix [][]intSeen) int {
	sols := [3]int{0, 0, 0}
	for i := range matrix {
		for j := range matrix[0] {
			if isLowPoint(matrix, i, j) {
				size := getBasinSize(matrix, i, j)
				if sols[0] < size {
					sols[0], sols[1], sols[2] = size, sols[0], sols[1]
				} else if sols[1] < size {
					sols[1], sols[2] = size, sols[1]
				} else if sols[2] < size {
					sols[2] = size
				}
			}
		}
	}

	return sols[0] * sols[1] * sols[2]
}

func getBasinSize(matrix [][]intSeen, i, j int) int {
	size := 0
	if !matrix[i][j].visto {
		matrix[i][j].visto = true
		size++
	}

	// Ahora ejecutamos el mismo algoritmo en las diferentes direcciones
	// si es posible

	// Para la izquierda
	if j-1 >= 0 && !matrix[i][j-1].visto &&
		matrix[i][j-1].num > matrix[i][j].num &&
		matrix[i][j-1].num != 9 {
		size += getBasinSize(matrix, i, j-1)
	}

	// Para la derecha
	if j+1 < len(matrix[0]) && !matrix[i][j+1].visto &&
		matrix[i][j+1].num > matrix[i][j].num &&
		matrix[i][j+1].num != 9 {
		size += getBasinSize(matrix, i, j+1)
	}

	// Para abajo
	if i+1 < len(matrix) && !matrix[i+1][j].visto &&
		matrix[i+1][j].num > matrix[i][j].num &&
		matrix[i+1][j].num != 9 {
		size += getBasinSize(matrix, i+1, j)
	}

	// Para arriba
	if i-1 >= 0 && !matrix[i-1][j].visto &&
		matrix[i-1][j].num > matrix[i][j].num &&
		matrix[i-1][j].num != 9 {
		size += getBasinSize(matrix, i-1, j)
	}

	return size
}
