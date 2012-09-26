package main

import (

	"fmt"
	"math"
)

//Shaper is an interface and has a single function Area that returns an int.
type Shaper interface {
   Area() float64
}

type Rectangle struct {
   length, width float64
}

//this function Area works on the type Rectangle and has the same function signature defined in the interface Shaper.  Therefore, Rectangle now implements the interface Shaper.
func (r Rectangle) Area() float64 {
   return r.length * r.width
}

type Circle struct {

	radius float64
}

func (c Circle) Area() float64 {

	return  math.Pi * c.radius * c.radius

}


func main() {
   r := Rectangle{length:5, width:3}
   fmt.Println("Rectangle r details are: ", r)  
   fmt.Println("Rectangle r's area is: ", r.Area())  

   s := Shaper(r)
   fmt.Println("Area of the Shape r is: ", s.Area())

   c := Circle{radius:3}
   fmt.Println("Circle c details are: ", c)  
   fmt.Println("Circle c's area is: ", c.Area())  

   s = Shaper(c)
   fmt.Println("Area of the Shape c is: ", s.Area())  

   
}