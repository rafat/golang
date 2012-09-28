package main

import (

"fmt"
"strings"

)

type Count int

func (count Count) Increment() Count {count++; return count}
func (count Count) Decrement() Count {count--; return count}

func (count *Count) Pincrement() {*count++}
func (count *Count) Pdecrement() {*count--}

func (count Count) IsZero() bool {return count == 0}

type Part struct {

	Id int
	Name string

}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}
func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func main() {

	
	type StringMap map[string]string
	type FloatChan chan float64
	var i Count = 7
        j := i.Increment()
	fmt.Println(j)
	i.Pdecrement()
	fmt.Println(i,i.IsZero())
	
	sm := make(StringMap)
	sm["January 2008"] = "243.7"
	sm["February 2008"] = "268.2"
	fmt.Println(sm)
	fc := make(FloatChan,1)
	fc <- 2.345
	fmt.Println(<-fc)
	iden := Part{Id: 7,Name: "Cali"}
	iden.UpperCase()
        fmt.Println(iden.Id,iden.Name)

}