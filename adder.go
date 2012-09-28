package main

import (

"fmt"
"reflect"

)
func add(x,y interface{}) interface{} {

	sa := reflect.ValueOf(x)
	lenx := sa.Len()
	fmt.Println(lenx)
	sb := reflect.ValueOf(y)
	leny := sb.Len()
	fmt.Println(leny)
	//oup := value.InterfaceData()
	//fmt.Println(value.Index(0).Interface())
	var max,min int

	if lenx > leny {
		max = lenx
		min = leny
		
	} else {
		max = leny
		min=lenx
	
	}


	for i := 0; i < min; i++ {
		tempx := sa.Index(i).Int()
		tempy := sb.Index(i).Int()
		return tempx+tempy
	}
	for i := min; i < max; i++ {
		if lenx > leny {
			tempx := sa.Index(i).Int()
			return tempx
		} else {
			tempy := sb.Index(i).Int()
			return tempy
		}
	}
	return nil
	

}

func AddInt(x,y []int) []int {

	var lenx,leny int
	lenx = len(x)
	leny = len(y)

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

	for i := 0; i < min; i++ {
		output[i]=x[i]+y[i]
	}
	for i := min; i < max; i++ {
		if lenx > leny {
			output[i]=x[i]
		} else {
			output[i]=y[i]
		}
	}
	return output


}

func AddFloat(x,y []float64) []float64 {

	var lenx,leny int
	lenx = len(x)
	leny = len(y)

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

	for i := 0; i < min; i++ {
		output[i]=x[i]+y[i]
	}
	for i := min; i < max; i++ {
		if lenx > leny {
			output[i]=x[i]
		} else {
			output[i]=y[i]
		}
	}
	return output


}


func main() {

	y := make([]int,4)
	z := make([]int,5)
	y[0]=2
	y[2]=7
	z[0]=2
	z[2]=5
	oup := AddInt(y,z)
	fmt.Println(oup)
	yf := make([]float64,4)
	zf := make([]float64,5)
	yf[0]=2.1
	zf[0]=2.2
	yf[1]=1.3
	zf[2]=2.7
	oupf := AddFloat(yf,zf)
	fmt.Println(oupf)

}