package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type cuboid struct {
	x1, x2, y1, y2, z1, z2 int
	sign                   bool // true is on an false is off
}

func parseCuboid(s string) cuboid {
	var c cuboid

	sAux1 := strings.Split(s, " ")
	c.sign = sAux1[0] == "on"

	sAux2 := strings.Split(sAux1[1], ",")

	c.x1, c.x2 = parseCuboidAux(sAux2[0])
	c.y1, c.y2 = parseCuboidAux(sAux2[1])
	c.z1, c.z2 = parseCuboidAux(sAux2[2])

	return c
}

func parseCuboidAux(s string) (int, int) {
	sAux := strings.Split(s[2:], "..")

	part1, _ := strconv.ParseInt(sAux[0], 10, 64)
	part2, _ := strconv.ParseInt(sAux[1], 10, 64)

	return int(part1), int(part2)
}

func printCuboid(c cuboid) {
	fmt.Println(c.sign, c.x1, c.x2, c.y1, c.y2, c.z1, c.z2)
}

func (c cuboid) intersection(cAux cuboid) (cuboid, bool) {
	x1 := maxInt(c.x1, cAux.x1)
	x2 := minInt(c.x2, cAux.x2)

	y1 := maxInt(c.y1, cAux.y1)
	y2 := minInt(c.y2, cAux.y2)

	z1 := maxInt(c.z1, cAux.z1)
	z2 := minInt(c.z2, cAux.z2)

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return cuboid{}, false
	}

	var newSign bool
	if c.sign && cAux.sign {
		newSign = false
	} else if !c.sign && !cAux.sign {
		newSign = true
	} else {
		newSign = cAux.sign
	}

	newCuboid := cuboid{x1, x2, y1, y2, z1, z2, newSign}

	return newCuboid, true
}

func (c cuboid) volume() int {
	dx := c.x2 - c.x1 + 1
	dy := c.y2 - c.y1 + 1
	dz := c.z2 - c.z1 + 1

	vol := dx * dy * dz

	if c.sign {
		return vol
	} else {
		return -vol
	}
}

func main() {

	start := time.Now()

	bs, err := ioutil.ReadFile("./Go/Day 22/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	input := make([]cuboid, len(inputS))
	for i, elem := range inputS {
		input[i] = parseCuboid(elem)
	}

	var stop int
	for i, elem := range input {
		if elem.x1 > 50 || elem.x1 < -50 {
			stop = i
			break
		}
	}

	inputPart1 := input[:stop]

	end := time.Since(start)

	sol1 := solve(inputPart1)
	sol2 := solve(input)

	fmt.Printf("Solution to part 1 is: %d\n", sol1)
	fmt.Printf("Solution to part 2 is: %d\n", sol2)
	fmt.Println("Time: ", end)

}

func solve(input []cuboid) int {

	var newList []cuboid

	for _, cub := range input {

		var toAdd []cuboid

		for _, cubAux := range newList {
			newCub, inters := cubAux.intersection(cub)
			if inters {
				toAdd = append(toAdd, newCub)
			}
		}

		newList = append(newList, toAdd...)

		if cub.sign {
			newList = append(newList, cub)
		}
	}

	totalVol := 0
	for _, cub := range newList {
		totalVol += cub.volume()
	}

	return totalVol

}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
