package main

import (

"fmt"
"reflect"

)
func itest(a interface{}) int {

	value := reflect.ValueOf(a)
	l := value.Len()
	fmt.Println(l)
	//oup := value.InterfaceData()
	fmt.Println(value.Index(0).Interface())

	return l

}

func main() {

	x := 8.6
	fmt.Printf("x is type %v with value %v\n",reflect.TypeOf(x),x)
	fmt.Printf("x is type %v with value %v\n",reflect.ValueOf(x),x)
	y := make([]int,4)
	z := make([]int,4)
	y[0]=2
	y[2]=7
	z[0]=2
	z[2]=5
	word := "chameleon"
	value := reflect.ValueOf(word)
	text := value.String()
	fmt.Println(text)
	sliceval := reflect.ValueOf(y)
	valuey := sliceval.Index(2).Int()
	valuez := reflect.ValueOf(z).Index(2).Int()
	valadd := valuey+valuez
	fmt.Printf("\n",valadd)
	new := itest(y)
	println(new)
}