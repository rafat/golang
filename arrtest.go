package main

import "fmt"

func rec(i int) {

	if i == 10 {
		return
	}
	fmt.Printf("%d ",i)
	rec(i+1)
	


}

func add(x []int, y []int) []int {
	var lenx,leny int
	lenx=len(x)
	leny=len(y)

	var max,min int
	if lenx > leny {
		max = lenx
		min = leny
	} else {
		max = leny
		min=lenx
	}

	output :=make([]int,max)

	for i := 0;i < min; i++ {
		output[i]=x[i]+y[i]
	}

	for i := min; i < max ; i++ {

		if lenx > leny {
			output[i]=x[i]
		} else {
			output[i]=y[i]
		}

	}

	return output


}

func main() {

	var a [10]int
	a[0]=22
	a[7]=7
	for i := 0; i < 10 ; i++ {

		fmt.Printf("%d ",a[i])
	}
	fmt.Print("\n")
	rec(0)

	s1:=a[0:5]
        s2:=a[5:8]
	var out []int
	out = add(s1,s2)
        fmt.Print(out)


}