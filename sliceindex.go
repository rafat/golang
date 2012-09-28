package main

import (

"fmt"
"reflect"

	)
// See Pages 235-238 for an alternative implementation

func IndexReflectX(xs interface{},x interface{}) int {

	if slice := reflect.ValueOf(xs);slice.Kind() == reflect.Slice {

		for i := 0 ; i < slice.Len(); i++ {
			switch y := slice.Index(i).Interface().(type) {
			case int :
				if y == x.(int) {
					return i
				}
			case string:
				if y == x.(string) {
					return i
				}

			}
		}
	}
	return -1



}

func main()  {

	xs := []int{2, 4, 6, 8}
	fmt.Println("5 @", IndexReflectX(xs, 5), " 6 @", IndexReflectX(xs, 6))

}

