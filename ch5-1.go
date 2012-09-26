package main

import "fmt"

// Type Assertion

type adder interface{

	add() int

}

type val struct{

	x,y int

}

func (v val) add() int {
	
	return v.x+v.y
}

func main() {

	var i interface{} = 99
	j := i.(int)
	fmt.Printf("%d->%T\n",j,j)
//	var out int
	var a,b int = 2,3
	vv := val{x: a, y: b}
	var w adder
	w=vv
	fmt.Print(w.add())

}