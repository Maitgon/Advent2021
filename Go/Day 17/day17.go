package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type coord struct {
	x, y int
}

func main() {

	start_ := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 17/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs)[15:], ", y=")
	xTargetS := strings.Split(input[0], "..")
	yTargetS := strings.Split(input[1], "..")

	xMin, _ := strconv.ParseInt(xTargetS[0], 10, 64)
	xMax, _ := strconv.ParseInt(xTargetS[1], 10, 64)
	yMin, _ := strconv.ParseInt(yTargetS[0], 10, 64)
	yMax, _ := strconv.ParseInt(yTargetS[1], 10, 64)

	min := coord{x: int(xMin), y: int(yMin)}
	max := coord{x: int(xMax), y: int(yMax)}

	start := coord{x: 0, y: 0}

	sol1 := part1(start, min, max)
	sol2 := part2(start, min, max)

	end := time.Since(start_)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(start, min, max coord) int {
	var xRange []int

	for xVel := 1; xVel < max.x; xVel++ {
		xPos := 0
		for xVelAux := xVel; xVelAux > 0 && xPos <= max.x; xVelAux-- {
			xPos += xVelAux
			if xPos >= min.x && xPos <= max.x && xVelAux == 1 {
				xRange = append(xRange, xVel)
				break
			}
		}
	}

	var yRange []int

	for yVel := 1; yVel <= -min.y; yVel++ {
		yPos := 0
		for yVelAux := yVel; yVelAux > -1000 && yPos >= min.y; yVelAux-- {
			yPos += yVelAux
			if yPos >= min.y && yPos <= max.y {
				yRange = append(yRange, yVel)
				break
			}
		}
	}

	maxY := 0
	for _, xVel := range xRange {
		for _, yVel := range yRange {
			init := start
			maxYAux := 0
			for i := 0; init.x <= max.x && init.y >= min.y; i++ {
				if xVel-i > 0 {
					init.x += xVel - i
				}
				init.y += yVel - i
				if init.y > maxY && init.y > maxYAux {
					maxYAux = init.y
				}
				if init.x >= min.x && init.y >= min.y &&
					init.x <= max.x && init.y <= max.y {
					maxY = maxYAux
				}
			}
		}
	}

	return maxY
}

func part2(start, min, max coord) int {
	var xRange []int

	for xVel := 1; xVel <= max.x; xVel++ {
		xPos := 0
		for xVelAux := xVel; xVelAux > 0 && xPos <= max.x; xVelAux-- {
			xPos += xVelAux
			if xPos >= min.x && xPos <= max.x {
				xRange = append(xRange, xVel)
				break
			}
		}
	}

	var yRange []int

	for yVel := min.y; yVel <= -min.y; yVel++ {
		yPos := 0
		for yVelAux := yVel; yVelAux > -1000 && yPos >= min.y; yVelAux-- {
			yPos += yVelAux
			if yPos >= min.y && yPos <= max.y {
				yRange = append(yRange, yVel)
				break
			}
		}
	}

	count := 0
	for _, xVel := range xRange {
		for _, yVel := range yRange {
			init := start
			for i := 0; init.x <= max.x && init.y >= min.y; i++ {
				if xVel-i > 0 {
					init.x += xVel - i
				}
				init.y += yVel - i
				if init.x >= min.x && init.y >= min.y &&
					init.x <= max.x && init.y <= max.y {
					count++
					break
				}
			}
		}
	}

	return count
}
