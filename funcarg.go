package main

import (

"fmt"
"math"

)

func Heron(a,b,c int) float64 {

	a1,b1,c1 := float64(a),float64(b),float64(c)
	s := (a1+b1+c1)/2
	return math.Sqrt(s * (s-a1) * (s-b1) * (s-c1))

}

func PythagoreanTriple(m,n int) (a,b,c int) {

	if m < n {

		m,n = n,m
	}
	return (m * m) - (n*n), (2 * m * n), (m * m) + (n * n)



}

func main() {

	for i := 1;i <= 4; i++ {

		x,y,z := PythagoreanTriple(i,i+1) 
		d1 := Heron(x,y,z)
		d2 := Heron(PythagoreanTriple(i,i+1))
		fmt.Printf("d1 == %10f == d2 == %10f\n", d1,d2)
	}

}