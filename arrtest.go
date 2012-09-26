package main

import "fmt"


func rec(i int) {

	if i == 10 {
		return
	}
	fmt.Printf("%d ",i)
	rec(i+1)
	


}

type Interface interface{
	Len() int
	Val(i int)
}

func add(x, y Interface) Interface {
	var lenx,leny int
	lenx = x.Len()
	leny = y.Len()

	var max,min int
	output := x
	if lenx > leny {
		max = lenx
		min = leny
		output = x
	} else {
		max = leny
		min=lenx
		output = y
	}


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

type IntAdd []int

func (p IntAdd) Len() int {len(p)}
func (p IntAdd) Val(i int) {return p[i]}

type FloatAdd []float32
func (p FloatAdd) Val(i int) {return p[i]}
func (p FloatAdd) Len() int {len(p)}

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