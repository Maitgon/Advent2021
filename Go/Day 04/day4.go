package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"time"
)

type newNum struct {
	num int16
	check bool
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 04/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n\n")

	numbersAux := strings.Split(input[0], ",")
	var numbers []int16
	for _, value := range numbersAux {
		val, _ := strconv.ParseInt(value, 10, 16)
		numbers = append(numbers, int16(val))
	}

	// To read all bingos
	bingosAux := input[1:]
	var bingos [][][]newNum
	for _, value := range bingosAux {
		bingosAux2 := strings.Split(value, "\n")

		// To read each one
		var bingos2 [][]newNum
		for _, value2 := range bingosAux2 {
			if value2[0] == ' ' {
				value2 = value2[1:]
			}
			bingosAux3 := strings.Split(strings.Replace(value2, "  ", " ", -1), " ") // Delete all extra " "
			var bingos3 []newNum
			for _, value3 := range bingosAux3 {
				val, _ := strconv.ParseInt(value3, 10, 16)
				bingos3 = append(bingos3, newNum{num: int16(val), check: false})
			}
			bingos2 = append(bingos2, bingos3)
		}
		bingos = append(bingos, bingos2)
	}

	sol1 := part1(numbers, bingos)
	sol2 := part2(numbers, bingos)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(nums []int16, bingos [][][]newNum) int64 {
	for _, num := range nums {
		bingos = updateBingo(num, bingos)
		for _, bingo := range bingos {
			if checkBingo(bingo){
				return int64(getNonMarked(bingo)) * int64(num)
			}
		}
	}
	return 0
}

func updateBingo(num int16, bingos [][][]newNum) [][][]newNum {
	for n := 0; n < len(bingos); n++ {
		for i := 0; i < len(bingos[0]); i++ {
			for j := 0; j < len(bingos[0][0]); j++ {
				if num == bingos[n][i][j].num {
					bingos[n][i][j].check = true
				}
			}
		}
	}
	return bingos
}

func checkBingo(bingo [][]newNum) bool {
	for i := 0; i < len(bingo); i++ {
		linea := true
		for j := 0; j < len(bingo[0]); j++ {
			linea = linea && bingo[i][j].check
		}
		if linea {
			return true
		}
	}

	for j := 0; j < len(bingo); j++ {
		columna := true
		for i := 0; i < len(bingo[0]); i++ {
			columna = columna && bingo[i][j].check
		}
		if columna {
			return true
		}
	}

	return false
}

func getNonMarked(bingo [][]newNum) int16 {
	sum := int16(0)
	for _, value := range bingo {
		for _, val := range value {
			if !val.check {
				sum += val.num
			}
		}
	}
	return sum
}

func part2(nums []int16, bingos [][][]newNum) int64 {
	yetWon := makeRange(0, len(bingos)) // Los que faltan por ganar
	var last int
	for _, num := range nums {
		bingos = updateBingo(num, bingos)
		for n, bingo := range bingos {
			if checkBingo(bingo){
				yetWon = remove(yetWon, n) // Vamos quitando los que ganan
			}
			if len(yetWon) == 1 {
				last = yetWon[0] // Nos quedamos con el ultimo
			} else if len(yetWon) == 0 {
				return int64(getNonMarked(bingos[last])) * int64(num)
			}
		}
	}
	return 0
}

func makeRange(min, max int) []int {
    a := make([]int, max-min)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func remove(l []int, item int) []int {
    for i, other := range l {
        if other == item {
            return append(l[:i], l[i+1:]...)
        }
    }
	return l
}