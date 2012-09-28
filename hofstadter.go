package main

import (

"fmt"

)

func HofstadterFemale(n int) int {

	if n<=0 {
		return 1
	}

	return n - HofstadterMale(HofstadterFemale(n-1))

}

func HofstadterMale(n int) int {

	if n<=0 {
		return 0
	}

	return n - HofstadterFemale(HofstadterMale(n-1))

}


func main() {
	females := make([]int,20)
	males := make([]int,len(females))
	for n:= range females {

		females[n]=HofstadterFemale(n)
		males[n]=HofstadterMale(n)
	}

	fmt.Println("F",females)
	fmt.Println("M",males)


}