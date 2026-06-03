// Day 2.

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

// Part 1 structs

type Move struct {
	dir string
	num int
}

type Position struct {
	depth, horizontal int
}

func moveOnce(pos Position, mov Move) Position {
	switch mov.dir {
	case "forward": pos.horizontal += mov.num
	case "down": pos.depth += mov.num
	case "up": pos.depth -= mov.num
	}
	return pos
}

func move(pos Position, movList []Move) Position {
	for _, move := range movList {
		pos = moveOnce(pos, move)
	}
	return pos
}

// Part 2 structs

type Position2 struct {
	depth, horizontal, aim int
}

func moveOnce2(pos Position2, mov Move) Position2 {
	switch mov.dir {
	case "forward": pos.horizontal += mov.num
				    pos.depth += pos.aim*mov.num
	case "down": pos.aim += mov.num
	case "up": pos.aim -= mov.num
	}
	return pos
}

func move2(pos Position2, movList []Move) Position2 {
	for _, move := range movList {
		pos = moveOnce2(pos, move)
	}
	return pos
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 02/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")
	input = input[:len(input)-1]

	var vals []Move
	for _, value := range input {
		aux := strings.Split(value, " ")
		val, _ := strconv.ParseInt(aux[1], 10, 16)
		vals = append(vals, Move{dir: aux[0], num: int(val)})
	}

	sol1 := part1(vals)
	sol2 := part2(vals)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(movList []Move) int {
	newPos := Position{depth: 0, horizontal: 0}
	finalPos := move(newPos, movList)
	return finalPos.depth * finalPos.horizontal
}

func part2(movList []Move) int {
	newPos := Position2{depth: 0, horizontal: 0, aim: 0}
	finalPos := move2(newPos, movList)
	return finalPos.depth * finalPos.horizontal
}