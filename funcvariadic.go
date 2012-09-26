package main

import (

"fmt"

)

func MinimumInt1(first int, rest ...int) int {

	for _, x := range rest {

		if x < first {
			first = x
		}

	}
	return first

}

func main() {

	numbers :=[]int{7,6,4,3,2,4,6,2}
	fmt.Println(MinimumInt1(numbers[0],numbers[1:]...))

}