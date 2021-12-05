package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

type cord struct {
	x, y int64
}

type path struct {
	start, end cord
}

func (p *path) isDiagonal() bool {
	return !(p.start.x == p.end.x || p.start.y == p.end.y)
}

func (p *path) generatePath() []cord {

	if p.isDiagonal() {
		fmt.Println("Es diagonal ", p.start.x, p.start.y, p.end.x,   p.end.y)
		return make([]cord, 0)
	}

	var completePath []cord

	if p.start.x == p.end.x {
		if p.start.y < p.end.y {
			for i := p.start.y; i <= p.end.y; i++ {
				completePath = append(completePath, cord{x: p.start.x, y: i})
			}
		} else {
			for i := p.start.y; i >= p.end.y; i-- {
				completePath = append(completePath, cord{x: p.start.x, y: i})
			}
		}
	} else if p.start.y == p.end.y {
		if p.start.x < p.end.x {
			for i := p.start.x; i <= p.end.x; i++ {
				completePath = append(completePath, cord{x: i, y: p.start.y})
			}
		} else {
			for i := p.start.x; i >= p.end.x; i-- {
				completePath = append(completePath, cord{x: i, y: p.start.y})
			}
		}
	}

	return completePath
}

func (p *path) generatePathDia() []cord {

	if !p.isDiagonal() {
		fmt.Println("Es diagonal ", p.start.x, p.start.y, p.end.x,   p.end.y)
		return make([]cord, 0)
	}

	var completePath []cord

	if p.start.x < p.end.x && p.start.y < p.end.y {
		for i := int64(0); i <= p.end.x - p.start.x; i++ {
			completePath = append(completePath, cord{x: p.start.x + i, y: p.start.y + i})
		}
	} else if p.start.x < p.end.x && p.start.y > p.end.y {
		for i := int64(0); i <= p.end.x - p.start.x; i++ {
			completePath = append(completePath, cord{x: p.start.x + i, y: p.start.y - i})
		}
	} else if p.start.x > p.end.x && p.start.y > p.end.y {
		for i := int64(0); i <= p.start.x - p.end.x; i++ {
			completePath = append(completePath, cord{x: p.start.x - i, y: p.start.y - i})
		}
	} else if p.start.x > p.end.x && p.start.y < p.end.y {
		for i := int64(0); i <= p.start.x - p.end.x; i++ {
			completePath = append(completePath, cord{x: p.start.x - i, y: p.start.y + i})
		}
	}

	return completePath
}



func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 05/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	var input2 []path
	for _, vals := range input[:len(input)-1] {
		aux := strings.Split(vals, " -> ")
		firstAux := strings.Split(aux[0], ",")
		startPathX, _ := strconv.ParseInt(firstAux[0], 10, 16)
		startPathY, _ := strconv.ParseInt(firstAux[1], 10, 16)
		secondAux := strings.Split(aux[1], ",")
		endPathX, _ := strconv.ParseInt(secondAux[0], 10, 16)
		endPathY, _ := strconv.ParseInt(secondAux[1], 10, 16)
		input2 = append(input2, path{start:cord{x: startPathX, y:startPathY}, end:cord{x: endPathX, y: endPathY}})
	}

	sol1, sol2 := part1y2(input2)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1y2(input []path) (int64, int64) {

	cordsAppearances := make(map[cord]int)

	for _, paths := range input {
		if !paths.isDiagonal() {
			cords := paths.generatePath()
			for _, cord := range cords {
				if _, ok := cordsAppearances[cord]; ok {
					cordsAppearances[cord] += 1
				} else {
					cordsAppearances[cord] = 1
				}
			}
		}
	}

	sol1 := 0
	for _, value := range cordsAppearances {
		if value >= 2 {
			sol1++
		}
	}

	for _, paths := range input {
		if paths.isDiagonal() {
			cords := paths.generatePathDia()
			for _, cord := range cords {
				if _, ok := cordsAppearances[cord]; ok {
					cordsAppearances[cord] += 1
				} else {
					cordsAppearances[cord] = 1
				}
			}
		}
	}

	sol2 := 0
	for _, value := range cordsAppearances {
		if value >= 2 {
			sol2++
		}
	}

	return int64(sol1), int64(sol2)
}