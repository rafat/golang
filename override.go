package main

import (

"fmt"

)

type Item struct {

	id string
	price float64
	quantity int

}

func (item *Item) Cost() float64 {

	return item.price * float64(item.quantity)
}

type SpecialItem struct {

	Item
	CatalogId int
	markup float64

}

func (item *SpecialItem) Cost() float64 {

	return item.Item.Cost() * item.markup
}

func main() {

	it := Item{"Red",5.23,20}
	fmt.Println(it)
	newitem := SpecialItem{it,34567,1.20}
	fmt.Println(newitem)
	fmt.Println(newitem.Cost())


}