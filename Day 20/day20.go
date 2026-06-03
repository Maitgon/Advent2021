package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var convert = map[byte]byte{
	'#': '1',
	'.': '0',
}

func main() {

	start := time.Now()

	bs, err := ioutil.ReadFile("./Go/Day 20/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n\n")

	enhancer := input[0]

	imageAux := strings.Split(input[1], "\n")
	image := make([][]byte, len(imageAux)+110)
	for i := range image {
		aux := make([]byte, len(imageAux[0])+110)
		for j := range aux {
			if i >= 55 && j >= 55 && i < len(imageAux)+55 && j < len(imageAux[0])+55 {
				aux[j] = convert[imageAux[i-55][j-55]]
			} else {
				aux[j] = '0'
			}
		}
		image[i] = aux
	}

	sol1, sol2 := part1y2(image, enhancer)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1y2(image [][]byte, enhancer string) (int, int) {
	i := 0
	for ; i < 2; i++ {
		image = applyEnhancer(image, enhancer)
	}

	count1 := 0
	for _, row := range image {
		for _, elem := range row {
			if elem == '1' {
				count1++
			}
		}
	}

	for ; i < 50; i++ {
		image = applyEnhancer(image, enhancer)
	}

	count2 := 0
	for _, row := range image {
		for _, elem := range row {
			if elem == '1' {
				count2++
			}
		}
	}

	return count1, count2
}

func applyEnhancer(image [][]byte, enhancer string) [][]byte {
	newImage := make([][]byte, len(image))
	for i := range image {
		auxImage := make([]byte, len(image[0]))
		for j := range image[0] {
			if i == 0 || i == len(image)-1 ||
				j == 0 || j == len(image[0])-1 {
				if image[i][j] == '0' {
					auxImage[j] = convert[enhancer[0]]
				} else {
					auxImage[j] = convert[enhancer[len(enhancer)-1]]
				}
			} else {
				pos := convertToPos(image, i, j)
				auxImage[j] = convert[enhancer[pos]]
			}
		}
		newImage[i] = auxImage
	}

	return newImage
}

func convertToPos(arr [][]byte, i, j int) int {
	var strB []byte
	for iAux := i - 1; iAux < i+2; iAux++ {
		for jAux := j - 1; jAux < j+2; jAux++ {
			strB = append(strB, arr[iAux][jAux])
		}
	}
	str := string(strB)
	pos, _ := strconv.ParseInt(str, 2, 16)
	return int(pos)
}
