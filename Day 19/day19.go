package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y, z int
}

var undefined = Point{x: -1, y: 2, z: -37983798}

type Scanner struct {
	pos     Point
	beacons Set
}

func (s *Scanner) Show() {
	for beacon, _ := range s.beacons.list {
		fmt.Println(beacon.x, beacon.y, beacon.z)
	}
}

type Set struct {
	list map[Point]struct{}
}

func (s *Set) Has(v Point) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v Point) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v Point) {
	delete(s.list, v)
}

func (s *Set) Clear() {
	s.list = make(map[Point]struct{})
}

func (s *Set) Size() int {
	return len(s.list)
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[Point]struct{})
	return s
}

//AddMulti Add multiple values in the set
func (s *Set) AddMulti(list ...Point) {
	for _, v := range list {
		s.Add(v)
	}
}

func (s *Set) Union(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

func (s *Set) Difference(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 19/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	stringS := strings.Split(string(bs), "\n\n")

	scanners := make([]Scanner, len(stringS))
	for i, stringAux := range stringS {
		stringSS := strings.Split(stringAux, "\n")
		scannerAux := NewSet()
		for _, stringAux2 := range stringSS[1:] {
			stringSSS := strings.Split(stringAux2, ",")
			xPoint, _ := strconv.ParseInt(stringSSS[0], 10, 64)
			yPoint, _ := strconv.ParseInt(stringSSS[1], 10, 64)
			zPoint, _ := strconv.ParseInt(stringSSS[2], 10, 64)
			scannerAux.Add(Point{x: int(xPoint), y: int(yPoint), z: int(zPoint)})
		}
		if i == 0 {
			scanners[i] = Scanner{pos: Point{x: 0, y: 0, z: 0}, beacons: *scannerAux}
		} else {
			scanners[i] = Scanner{pos: undefined, beacons: *scannerAux}
		}
	}

	sol1, sol2 := part1y2(scanners)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1y2(scanners []Scanner) (int, int) {
	var completed bool
	scan0 := scanners[0]
	for !completed {
		completed = true
		for j, _ := range scanners[1:] {
			completed = completed && scanners[j+1].pos != undefined

			if scanners[j+1].pos == undefined {
				if newBeacons, ok := Match(&scan0, &scanners[j+1]); ok {
					scan0.beacons = *scan0.beacons.Union(&newBeacons)
				}
			}

		}

	}

	maxDist := 0
	for _, s1 := range scanners {
		for _, s2 := range scanners {
			dist := Manhattan(s1.pos, s2.pos)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return scan0.beacons.Size(), maxDist

}

func Manhattan(p1, p2 Point) int {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y) + Abs(p1.z-p2.z)
}

func Abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -1
	}
}

func Match(scan0, scan1 *Scanner) (Set, bool) {
	rotMatrix := getAllMatrices()
	maxVal := 0
	for _, rot := range rotMatrix {
		matches := make(map[Point]int)
		for point1 := range scan1.beacons.list {
			for point0 := range scan0.beacons.list {
				rotPoint := MultPoint(rot, point1)
				dif := minus(point0, rotPoint)
				matches[dif] += 1
			}
		}

		for key, val := range matches {
			if val >= 12 {
				scan1.pos = key

				newBeacons := NewSet()
				for point1, _ := range scan1.beacons.list {
					rotPoint := MultPoint(rot, point1)
					newBeacons.Add(add(rotPoint, scan1.pos))
				}
				return *newBeacons, true
			} else {
				if maxVal < val {
					maxVal = val
				}
			}
		}

	}

	//fmt.Println(maxVal)

	return *NewSet(), false
}

func minus(p1, p2 Point) Point {
	return Point{x: p1.x - p2.x, y: p1.y - p2.y, z: p1.z - p2.z}
}

func add(p1, p2 Point) Point {
	return Point{x: p1.x + p2.x, y: p1.y + p2.y, z: p1.z + p2.z}
}

func getAllMatrices1() [24][3][3]int {
	var sol [24][3][3]int
	xRot := [3][3]int{
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	}
	yRot := [3][3]int{
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	}
	zRot := [3][3]int{
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	}

	sol[0] = [3][3]int{ // I
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	sol[1] = xRot                                                        // X
	sol[2] = yRot                                                        // Y
	sol[3] = zRot                                                        // Z
	sol[4] = MultMatrix(xRot, xRot)                                      // XX
	sol[5] = MultMatrix(xRot, yRot)                                      // XY
	sol[6] = MultMatrix(xRot, zRot)                                      // XZ
	sol[7] = MultMatrix(yRot, xRot)                                      // YX
	sol[8] = MultMatrix(yRot, yRot)                                      // YY
	sol[9] = MultMatrix(zRot, yRot)                                      // ZY
	sol[10] = MultMatrix(zRot, zRot)                                     // ZZ
	sol[11] = MultMatrix(MultMatrix(xRot, xRot), xRot)                   // XXX
	sol[12] = MultMatrix(MultMatrix(xRot, xRot), yRot)                   // XXY
	sol[13] = MultMatrix(MultMatrix(xRot, xRot), zRot)                   // XXZ
	sol[14] = MultMatrix(MultMatrix(xRot, yRot), xRot)                   // XYX
	sol[15] = MultMatrix(MultMatrix(xRot, yRot), yRot)                   // XYY
	sol[16] = MultMatrix(MultMatrix(xRot, zRot), zRot)                   // XZZ
	sol[17] = MultMatrix(MultMatrix(yRot, xRot), xRot)                   // YXX
	sol[18] = MultMatrix(MultMatrix(yRot, yRot), yRot)                   // YYY
	sol[19] = MultMatrix(MultMatrix(zRot, zRot), zRot)                   // ZZZ
	sol[20] = MultMatrix(MultMatrix(MultMatrix(xRot, xRot), xRot), yRot) // XXXY
	sol[21] = MultMatrix(MultMatrix(MultMatrix(xRot, xRot), yRot), xRot) // XXYX
	sol[22] = MultMatrix(MultMatrix(MultMatrix(xRot, yRot), xRot), xRot) // XYXX
	sol[23] = MultMatrix(MultMatrix(MultMatrix(yRot, xRot), xRot), xRot) // YXXX

	return sol

}

func getAllMatrices() [24][3][3]int {
	var sol [24][3][3]int
	xRot := [3][3]int{
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	}
	yRot := [3][3]int{
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	}
	sol[0] = [3][3]int{ // I
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	sol[1] = xRot                                                                          // X
	sol[2] = yRot                                                                          // Y
	sol[3] = MultMatrix(xRot, xRot)                                                        // XX
	sol[4] = MultMatrix(xRot, yRot)                                                        // XY
	sol[5] = MultMatrix(yRot, xRot)                                                        // YX
	sol[6] = MultMatrix(yRot, yRot)                                                        // YY
	sol[7] = MultMatrix(MultMatrix(xRot, xRot), xRot)                                      // XXX
	sol[8] = MultMatrix(MultMatrix(xRot, xRot), yRot)                                      // XXY
	sol[9] = MultMatrix(MultMatrix(xRot, yRot), xRot)                                      // XYX
	sol[10] = MultMatrix(MultMatrix(xRot, yRot), yRot)                                     // XYY
	sol[11] = MultMatrix(MultMatrix(yRot, xRot), xRot)                                     // YXX
	sol[12] = MultMatrix(MultMatrix(yRot, yRot), xRot)                                     // YYX
	sol[13] = MultMatrix(MultMatrix(yRot, yRot), yRot)                                     // YYX
	sol[14] = MultMatrix(MultMatrix(MultMatrix(xRot, xRot), xRot), yRot)                   // XXXY
	sol[15] = MultMatrix(MultMatrix(MultMatrix(xRot, xRot), yRot), xRot)                   // XXYX
	sol[16] = MultMatrix(MultMatrix(MultMatrix(xRot, xRot), yRot), yRot)                   // XXYY
	sol[17] = MultMatrix(MultMatrix(MultMatrix(xRot, yRot), xRot), xRot)                   // XYXX
	sol[18] = MultMatrix(MultMatrix(MultMatrix(xRot, yRot), yRot), yRot)                   // XYYY
	sol[19] = MultMatrix(MultMatrix(MultMatrix(yRot, xRot), xRot), xRot)                   // YXXX
	sol[20] = MultMatrix(MultMatrix(MultMatrix(yRot, yRot), yRot), xRot)                   // YYYX
	sol[21] = MultMatrix(MultMatrix(MultMatrix(MultMatrix(xRot, xRot), xRot), yRot), xRot) // XXXYX
	sol[22] = MultMatrix(MultMatrix(MultMatrix(MultMatrix(xRot, yRot), xRot), xRot), xRot) // XXXYX
	sol[23] = MultMatrix(MultMatrix(MultMatrix(MultMatrix(xRot, yRot), yRot), yRot), xRot) // XXXYX

	return sol

}

func MultMatrix(x, y [3][3]int) [3][3]int {
	var sol [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var aux int
			for k := 0; k < 3; k++ {
				aux += x[i][k] * y[k][j]
			}
			sol[i][j] = aux
		}
	}

	return sol
}

func MultPoint(m [3][3]int, p Point) Point {
	x := m[0][0]*p.x + m[0][1]*p.y + m[0][2]*p.z
	y := m[1][0]*p.x + m[1][1]*p.y + m[1][2]*p.z
	z := m[2][0]*p.x + m[2][1]*p.y + m[2][2]*p.z

	return Point{x: x, y: y, z: z}
}
