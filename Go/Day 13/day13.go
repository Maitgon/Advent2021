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
		fmt.Println(string(val))
	}
	fmt.Println("Time: ", end)

}

func part1y2(coordinates []point, folds []folding) (int, [][]rune) {

	foldOnce(&coordinates, folds[0])
	count := len(coordinates)

	for i := 1; i < len(folds); i++ {
		foldOnce(&coordinates, folds[i])
	}

	sol2 := make([][]rune, 6)
	for i := 0; i < 6; i++ {
		aux := make([]rune, 39)
		for j := 0; j < 39; j++ {
			aux[j] = ' '
		}
		sol2[i] = aux
	}

	for _, val := range coordinates {
		sol2[val.y][val.x] = '█' // WTF is even this thing
	}

	return count, sol2
}

func foldOnce(coordinates *[]point, fold folding) {
	if fold.axis == "x" {
		for i := 0; i < len(*coordinates); {
			if (*coordinates)[i].x > fold.pos {
				newPoint := point{x: 2*fold.pos - (*coordinates)[i].x, y: (*coordinates)[i].y}
				if !search(*coordinates, newPoint) {
					(*coordinates)[i] = newPoint
					i++
				} else {
					remove(coordinates, i)
				}
			} else {
				i++
			}
		}
	} else if fold.axis == "y" {
		for i := 0; i < len(*coordinates); {
			if (*coordinates)[i].y > fold.pos {
				newPoint := point{y: 2*fold.pos - (*coordinates)[i].y, x: (*coordinates)[i].x}
				if !search(*coordinates, newPoint) {
					(*coordinates)[i] = newPoint
					i++
				} else {
					remove(coordinates, i)
				}
			} else {
				i++
			}
		}
	}

}

func remove(s *[]point, i int) {
	(*s)[i] = (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
}

func search(s []point, p point) bool {
	for _, elem := range s {
		if elem == p {
			return true
		}
	}
	return false
}
