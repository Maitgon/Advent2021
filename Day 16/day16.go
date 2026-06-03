package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type packet struct {
	version  int64
	typeID   int64
	lenType  int64    // -1 if typeID == 4
	length   int64    // -1 if typeID == 4
	value    int64    // -1 if typeID != 4
	subPacks []packet // empty if typeID == 4
}

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 16/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("inputAgu.txt")
	}

	hexToBin := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	inputS := make([]string, len(string(bs)))
	for _, char := range string(bs) {
		inputS = append(inputS, hexToBin[char])
	}

	input := strings.Join(inputS, "")

	pos := int64(0)
	p := parsePacket(input, &pos)

	sol1 := getSumVersions(p)
	sol2 := getValue(p)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func parsePacket(input string, pos *int64) packet {

	var p packet

	// Leemos la packet version
	version, _ := strconv.ParseInt(input[*pos:*pos+3], 2, 64)
	*pos += 3

	// Leemos el packet Type ID
	typeID, _ := strconv.ParseInt(input[*pos:*pos+3], 2, 64)
	*pos += 3

	p.version = version
	p.typeID = typeID

	// Si es un literal
	if p.typeID == 4 {
		p.lenType = -1
		p.length = -1

		// Now, we get the value
		var valueS []string
		var control byte
		for control != '0' {
			valueS = append(valueS, input[*pos+1:*pos+5])
			control = input[*pos]
			*pos += 5
		}
		value, _ := strconv.ParseInt(strings.Join(valueS, ""), 2, 64)
		p.value = value

	} else {
		p.value = -1
		p.lenType = int64(input[*pos] - '0')
		*pos++

		if p.lenType == 0 {
			length, _ := strconv.ParseInt(input[*pos:*pos+15], 2, 64)
			p.length = length
			*pos = *pos + 15
			limit := *pos + p.length
			var packets []packet
			for *pos < limit {
				packetAux := parsePacket(input, pos)
				packets = append(packets, packetAux)
			}
			p.subPacks = packets
		} else {
			num, _ := strconv.ParseInt(input[*pos:*pos+11], 2, 64)
			p.length = num
			*pos = *pos + 11
			var packets []packet
			for i := 0; int64(i) < p.length; i++ {
				packetAux := parsePacket(input, pos)
				packets = append(packets, packetAux)
			}
			p.subPacks = packets
		}
	}

	return p

}

func printPacket(p packet) {

	fmt.Println("packet: ", p.version, p.typeID, p.value, p.lenType, p.length)
	for i, pack := range p.subPacks {
		fmt.Println(i, "subpack")
		printPacket(pack)
	}
}

func getSumVersions(p packet) int64 {
	version := p.version
	for _, pAux := range p.subPacks {
		version += getSumVersions(pAux)
	}
	return version
}

func getValue(p packet) int64 {
	val := int64(0)
	switch p.typeID {
	case 0:
		for _, pAux := range p.subPacks {
			val += getValue(pAux)
		}

	case 1:
		val = 1
		for _, pAux := range p.subPacks {
			val *= getValue(pAux)
		}

	case 2:
		min := getValue(p.subPacks[0])
		for _, pAux := range p.subPacks[1:] {
			aux := getValue(pAux)
			if aux < min {
				min = aux
			}
		}
		val = min

	case 3:
		max := getValue(p.subPacks[0])
		for _, pAux := range p.subPacks[1:] {
			aux := getValue(pAux)
			if aux > max {
				max = aux
			}
		}
		val = max

	case 4:
		val = p.value

	case 5:
		if getValue(p.subPacks[0]) > getValue(p.subPacks[1]) {
			val = 1
		} else {
			val = 0
		}

	case 6:
		if getValue(p.subPacks[0]) < getValue(p.subPacks[1]) {
			val = 1
		} else {
			val = 0
		}

	case 7:
		if getValue(p.subPacks[0]) == getValue(p.subPacks[1]) {
			val = 1
		} else {
			val = 0
		}

	default:
		fmt.Println("error Default")

	}

	if val == -1 {
		fmt.Println("error -1")
	}

	return val
}
