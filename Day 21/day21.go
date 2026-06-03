package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	pos1 := 7
	pos2 := 1

	sol1 := part1(pos1, pos2)
	sol2 := part2(pos1, pos2)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(pos1, pos2 int) int {
	score1 := 0
	score2 := 0
	dice := 1
	rolled := 0
	for {
		roll1 := dice
		roll(&dice)
		roll1 += dice
		roll(&dice)
		roll1 += dice
		roll(&dice)
		pos1 = actPos(pos1, roll1)
		score1 += pos1

		rolled += 3

		if score1 >= 1000 {
			break
		}

		roll2 := dice
		roll(&dice)
		roll2 += dice
		roll(&dice)
		roll2 += dice
		roll(&dice)
		pos2 = actPos(pos2, roll2)
		score2 += pos2

		rolled += 3

		if score2 >= 1000 {
			break
		}
	}

	var sol int

	if score1 >= 1000 {
		sol = score2
	} else {
		sol = score1
	}

	return sol * rolled
}

func part2(pos1, pos2 int) int64 {
	score1, score2 := 0, 0
	turn := 1
	win13, win23 := part2Aux(pos1, pos2, score1, score2, turn, 3)
	win14, win24 := part2Aux(pos1, pos2, score1, score2, turn, 4)
	win15, win25 := part2Aux(pos1, pos2, score1, score2, turn, 5)
	win16, win26 := part2Aux(pos1, pos2, score1, score2, turn, 6)
	win17, win27 := part2Aux(pos1, pos2, score1, score2, turn, 7)
	win18, win28 := part2Aux(pos1, pos2, score1, score2, turn, 8)
	win19, win29 := part2Aux(pos1, pos2, score1, score2, turn, 9)

	win1 := win13 + 3*win14 + 6*win15 + 7*win16 + 6*win17 + 3*win18 + win19
	win2 := win23 + 3*win24 + 6*win25 + 7*win26 + 6*win27 + 3*win28 + win29
	if win1 > win2 {
		return win1
	} else {
		return win2
	}
}

func part2Aux(pos1, pos2, score1, score2, turn, dice int) (int64, int64) {
	if turn == 1 {
		pos1 = actPos(pos1, dice)
		score1 += pos1
		if score1 >= 21 {
			return 1, 0
		}
		turn = 2
	} else if turn == 2 {
		pos2 = actPos(pos2, dice)
		score2 += pos2
		if score2 >= 21 {
			return 0, 1
		}
		turn = 1
	}

	win13, win23 := part2Aux(pos1, pos2, score1, score2, turn, 3)
	win14, win24 := part2Aux(pos1, pos2, score1, score2, turn, 4)
	win15, win25 := part2Aux(pos1, pos2, score1, score2, turn, 5)
	win16, win26 := part2Aux(pos1, pos2, score1, score2, turn, 6)
	win17, win27 := part2Aux(pos1, pos2, score1, score2, turn, 7)
	win18, win28 := part2Aux(pos1, pos2, score1, score2, turn, 8)
	win19, win29 := part2Aux(pos1, pos2, score1, score2, turn, 9)

	win1 := win13 + 3*win14 + 6*win15 + 7*win16 + 6*win17 + 3*win18 + win19
	win2 := win23 + 3*win24 + 6*win25 + 7*win26 + 6*win27 + 3*win28 + win29
	return win1, win2
}

func roll(dice *int) {
	if *dice == 100 {
		*dice = 1
	} else {
		*dice++
	}
}

func actPos(pos, move int) int {
	if (pos+move)%10 == 0 {
		return 10
	} else {
		return (pos + move) % 10
	}
}
