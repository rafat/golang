package main

import (

"fmt"
	)

type Exchanger interface {

	Exchange()

}

type StringPair struct {first,second string}

func (strpair *StringPair) Exchange() {strpair.first, strpair.second = strpair.second, strpair.first}

type Point [2]int

func (point *Point) Exchange() {point[0], point[1] = point[1], point[0]}

func ExchangeThese(exchanger ...Exchanger) {

	for _, x := range exchanger {

		x.Exchange()
	}

}

func main() {

	pair := StringPair{first: "Rafat",second: "Hussain"}
	pts := Point{4,8}
	var e Exchanger
	e = &pair
	e.Exchange()
	fmt.Println(e)
	e = &pts
	e.Exchange()
	fmt.Println(e)

	ExchangeThese(&pair,&pts)
	fmt.Println(pair,pts)


}