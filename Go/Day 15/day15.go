package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

type cell struct {
	x, y, risk int32
}

type cellHeap []cell

func (h cellHeap) Len() int           { return len(h) }
func (h cellHeap) Less(i, j int) bool { return h[i].risk < h[j].risk }
func (h cellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *cellHeap) Push(x interface{}) {
	*h = append(*h, x.(cell))
}

func (h *cellHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 15/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	matrix := make([][]int32, len(inputS))
	for i, val := range inputS {
		aux := make([]int32, len(inputS[0]))
		for j, num := range val {
			aux[j] = int32(num - '0')
		}
		matrix[i] = aux
	}

	sol1 := part1(matrix)

	matrix2 := make([][]int32, len(matrix)*5)
	for i := range matrix2 {
		matrix2[i] = make([]int32, len(matrix[0])*5)
	}
	for i := range matrix {
		for j := range matrix[0] {
			for ki := 0; ki < 5; ki++ {
				for kj := 0; kj < 5; kj++ {
					if int(matrix[i][j])+kj+ki > 9 {
						matrix2[ki*len(matrix)+i][kj*len(matrix)+j] = ((matrix[i][j] + int32(kj+ki)) % 10) + 1
					} else {
						matrix2[ki*len(matrix)+i][kj*len(matrix)+j] = matrix[i][j] + int32(kj+ki)
					}
				}
			}
		}
	}

	sol2 := part1(matrix2)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(matrix [][]int32) int32 {

	distances := make([][]int32, len(matrix))
	for i := range distances {
		aux := make([]int32, len(matrix[0]))
		for j := range aux {
			aux[j] = math.MaxInt32
		}
		distances[i] = aux
	}

	dx := []int32{-1, 0, 1, 0}
	dy := []int32{0, 1, 0, -1}

	var setCells cellHeap
	setCells.Push(cell{x: 0, y: 0, risk: 0})

	distances[0][0] = matrix[0][0]

	for len(setCells) != 0 {

		cell0 := setCells[0]
		heap.Pop(&setCells)

		for i := 0; i < 4; i++ {
			x := cell0.x + dx[i]
			y := cell0.y + dy[i]

			if !isInsideGrid(matrix, x, y) {
				continue
			} else if distances[x][y] > distances[cell0.x][cell0.y]+matrix[x][y] {
				if distances[x][y] != math.MaxInt32 {
					if ok, k := binarySearch(cell{x: x, y: y, risk: distances[x][y]}, setCells); ok {
						heap.Remove(&setCells, k)
					}
				}
				distances[x][y] = distances[cell0.x][cell0.y] + matrix[x][y]
				setCells.Push(cell{x: x, y: y, risk: distances[x][y]})

			}
		}

	}

	return distances[len(matrix)-1][len(matrix[0])-1] - matrix[0][0]

}

func isInsideGrid(matrix [][]int32, i, j int32) bool {
	return i >= 0 && i < int32(len(matrix)) && j >= 0 && j < int32(len(matrix[0]))
}

func Min(x, y int32) int32 {
	if x <= y {
		return x
	} else {
		return y
	}
}

func binarySearch(needle cell, haystack cellHeap) (bool, int) {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median].risk < needle.risk {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false, -1
	}

	return true, low
}

/*
func removeFirst(s *[]cell) {
	(*s) = (*s)[1:]
}

func search(s []cell, p cell) (bool, int32) {
	for i, elem := range s {
		if elem.x == p.x && elem.y == p.y {
			return true, int32(i)
		}
	}
	return false, -1
}
*/
