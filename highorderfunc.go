package main

import (
"fmt"
)

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func main() {

	xs := []int{2, 4, 6, 8}
	t := SliceIndex(len(xs),func(i int) bool { return xs[i] == 4 })
	fmt.Println(t)

}
