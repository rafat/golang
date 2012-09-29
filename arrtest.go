package main

import (
"fmt"
)


func rec(i int) {

    if i == 10 {
        return
    }
    fmt.Printf("%d ",i)
    rec(i+1)
    


}

type Interface interface{
    Len() int
    SetVal(a interface{}, i int)
    GetVal(i int) interface{}
    AddVal(a interface{}, i int) interface{}
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



    for i := 0; i < min; i++ {
        output.SetVal(x.AddVal(y.GetVal(i),i),i)
    }
    for i := min; i < max; i++ {
        if lenx > leny {
            output.SetVal(x.GetVal(i),i)
        } else {
            output.SetVal(y.GetVal(i),i)
        }
    }




    return output


}

type IntAdd []int

func (p IntAdd) Len() int {return len(p)}
func (p IntAdd) SetVal(val interface{},i int) {p[i] = val.(int)}
func (p IntAdd) GetVal(i int) interface{} {return p[i]}
func (p IntAdd) AddVal(val interface{},i int) interface{} {return p[i]+val.(int)}


type FloatAdd []float64

func (p FloatAdd) Len() int {return len(p)}
func (p FloatAdd) SetVal(val interface{},i int) {p[i] = val.(float64)}
func (p FloatAdd) GetVal(i int) interface{} {return p[i]}
func (p FloatAdd) AddVal(val interface{},i int) interface{} {return p[i]+val.(float64)}

func addint(x,y []int) []int {
	oup := add(IntAdd(x),IntAdd(y))
	loup := oup.Len()
	oup2 := make([]int,loup)
	for i := 0; i < loup; i++ {
		oup2[i]=(oup.GetVal(i)).(int)
	}
	return oup2
}

func addfloat(x,y []float64) []float64 {
	oup := add(FloatAdd(x),FloatAdd(y))
	loup := oup.Len()
	oup2 := make([]float64,loup)
	for i := 0; i < loup; i++ {
		oup2[i]=(oup.GetVal(i)).(float64)
	}
	return oup2
}

func main() {

    var a [10]float64
    a[0]=22.2
    a[7]=7.3
    for i := 0; i < 10 ; i++ {

        fmt.Printf("%d ",a[i])
    }
    fmt.Print("\n")
    rec(0)

    s1:=a[0:5]
    s2:=a[5:8]
    out := addfloat(s1,s2)
    fmt.Print(out)



}