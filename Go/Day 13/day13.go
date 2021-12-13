package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

type folding struct {
	pos  int
	axis string
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 13/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n\n")
	coordinatesS := strings.Split(inputS[0], "\n")
	foldsS := strings.Split(inputS[1], "\n")

	coordinates := make([]point, len(coordinatesS))
	for i, val := range coordinatesS {
		coord := strings.Split(string(val), ",")
		x, _ := strconv.ParseInt(coord[0], 10, 16)
		y, _ := strconv.ParseInt(coord[1], 10, 16)
		coordinates[i] = point{x: int(x), y: int(y)}
	}

	folds := make([]folding, len(foldsS))
	for i, val := range foldsS {
		fold := strings.Split(string(val)[11:], "=")
		pos, _ := strconv.ParseInt(fold[1], 10, 16)
		folds[i] = folding{pos: int(pos), axis: fold[0]}
	}

	sol1, sol2 := part1y2(coordinates, folds)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is, uhmmmm, this????????: ")
	for _, val := range sol2 {
		fmt.Println(val)
	}
	fmt.Println("Time: ", end)

}

func part1y2(coordinates []point, folds []folding) (int, [][]string) {
	grid := make([][]string, folds[0].pos*2+1)
	for i := 0; i <= folds[0].pos*2; i++ {
		grid[i] = make([]string, folds[0].pos*2+1)
		for j := 0; j <= folds[0].pos*2; j++ {
			grid[i][j] = "."
		}
	}

	// Rellenamos el tablero
	for _, coord := range coordinates {
		grid[coord.y][coord.x] = "#"
	}

	foldOnce(grid, folds[0])

	count := 0
	for i := range grid {
		for j := range grid {
			if grid[i][j] == "#" {
				count++
			}
		}
	}

	for i := 1; i < len(folds); i++ {
		foldOnce(grid, folds[i])
	}

	sol2 := make([][]string, 6)
	for i := 0; i < 6; i++ {
		sol2[i] = (grid[i])[:39]
	}

	return count, sol2
}

func foldOnce(grid [][]string, fold folding) {
	if fold.axis == "x" {
		for i := range grid {
			for j := fold.pos + 1; j < len(grid); j++ {
				if grid[i][j] == "#" {
					grid[i][2*fold.pos-j] = "#"
					grid[i][j] = "."
				}
			}
		}
	} else if fold.axis == "y" {
		for i := fold.pos + 1; i < len(grid); i++ {
			for j := range grid {
				if grid[i][j] == "#" {
					grid[2*fold.pos-i][j] = "#"
					grid[i][j] = "."
				}
			}
		}
	}
}
