package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Pair struct {
	lVal  int
	rVal  int
	lPair *Pair
	rPair *Pair
	deep  int
}

func (p *Pair) Show() string {
	str := ""
	str += "["

	if p.lVal != -1 {
		str += string(byte(p.lVal) + byte('0'))
	} else {
		str += p.lPair.Show()
	}

	str += ","

	if p.rVal != -1 {
		str += string(byte(p.rVal) + byte('0'))
	} else {
		str += p.rPair.Show()
	}

	str += "]"

	if p.deep == 1 {
		fmt.Println(str)
	}

	return str
}

func parse(str string, pos *int, deep int) *Pair {
	var pair Pair

	// Saltamos el "["
	*pos++

	pair.deep = deep
	// Leemos la parte izquierda
	if str[*pos] == '[' {
		pair.lVal = -1
		pair.lPair = parse(str, pos, deep+1)
	} else {
		pair.lVal = int(str[*pos] - '0')
		*pos++
	}

	// Saltamos la coma
	*pos++

	// Leemos la parte derecha
	if str[*pos] == '[' {
		pair.rVal = -1
		pair.rPair = parse(str, pos, deep+1)
	} else {
		pair.rVal = int(str[*pos] - '0')
		*pos++
	}

	// Saltamos el ']'
	*pos++

	return &pair
}

func main() {
	//start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./Go/Day 18/inputAux2.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	stringS := strings.Split(string(bs), "\n")

	pairs := make([]*Pair, len(stringS))
	for i, val := range stringS {
		pos := 0
		pairs[i] = parse(val, &pos, 1)
	}

	//pairs[0].explode()
	//pairs[0].Show()
	pairs[0].reduct()
	pairs[0].Show()

	//part1(pairs)

}

func (p *Pair) explode() bool {
	var last, pos *int
	last = nil
	pos = nil
	actual := -1
	sol, _ := p.explodeAux(last, &actual, pos)
	return sol
}

func (p *Pair) explodeAux(last, actual, pos *int) (bool, *int) {

	exploded := false

	if p.lVal != -1 {
		if *actual != -1 {
			p.lVal += *actual
			*actual = -1
			return true, last
		}
		last = &p.lVal
	} else if p.deep == 4 && *actual == -1 {
		if last != nil {
			fmt.Println(nil, last)
			*last += p.lPair.lVal
		}
		*actual = p.lPair.rVal
		p.lPair = nil
		p.lVal = 0
	} else {
		exploded, last = p.lPair.explodeAux(last, actual, pos)
	}

	if exploded {
		return true, last
	}

	if p.rVal != -1 {
		if *actual != -1 {
			p.rVal += *actual
			*actual = -1
			return true, last
		}
		last = &p.rVal
	} else if p.deep == 4 && *actual == -1 {
		if last != nil {
			*last += p.rPair.lVal
		}
		*actual = p.rPair.rVal
		p.rPair = nil
		p.rVal = 0
	} else {
		exploded, last = p.rPair.explodeAux(last, actual, pos)
	}

	return exploded, last

}

func (p *Pair) split() bool {

	splited := false

	if p.lVal > 9 {
		p.lPair = &Pair{lVal: p.lVal / 2, rVal: p.lVal/2 + p.lVal%2, lPair: nil, rPair: nil, deep: p.deep + 1}
		p.lVal = -1
		return true
	} else if p.lVal == -1 {
		splited = p.lPair.split()
	}

	if splited {
		return true
	}

	if p.rVal > 9 {
		p.rPair = &Pair{lVal: p.rVal / 2, rVal: p.rVal/2 + p.rVal%2, lPair: nil, rPair: nil, deep: p.deep + 1}
		p.rVal = -1
		return true
	} else if p.rVal == -1 {
		splited = p.rPair.split()
	}

	return splited
}

func equal(p1, p2 Pair) bool {

	var aux1 bool
	if p1.lPair != p2.lPair {
		aux1 = false
	} else if p1.lPair != nil {
		aux1 = equal(*p1.lPair, *p2.lPair)
	}

	var aux2 bool
	if p1.rPair != p2.rPair {
		aux2 = false
	} else if p1.rPair != nil {
		aux2 = equal(*p1.rPair, *p2.rPair)
	}

	return p1.lVal == p2.lVal &&
		p1.rVal == p2.rVal &&
		p1.deep == p2.deep &&
		aux1 && aux2
}

func (p *Pair) reduct() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			p.Show()
			p.explode()
		}
		for j := 0; j < 10; j++ {
			p.Show()
			p.split()
		}
	}
}

func sum(p1, p2 *Pair) *Pair {
	moreDeep(p1)
	moreDeep(p2)
	return &Pair{lPair: p1, rPair: p2, lVal: -1, rVal: -1, deep: 1}
}

func moreDeep(p *Pair) {
	if p != nil {
		p.deep++
		moreDeep(p.lPair)
		moreDeep(p.rPair)
	}
}

func part1(pairs []*Pair) {
	p := pairs[0]
	for _, pAux := range pairs[1:] {
		p = sum(p, pAux)
		p.reduct()
	}

	p.Show()
}
