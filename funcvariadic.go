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

func Minimum(first interface{}, rest ...interface{}) interface{} {


	minimum := first
	for _, x := range rest {

		switch x := x.(type) {

		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}


		} 



	}
return minimum


}

func main() {

	//numbers :=[]int{7,6,4,3,2,4,6,2}
        number := Minimum(7.1,6.2,4.2,3.0,2.2,4.1,6.1,1.3).(float64)
	fmt.Println(number)

}