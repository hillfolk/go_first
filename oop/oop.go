package main

import "fmt"

//구조체
type Item struct {
	name string
	price float64
	quantity int
}

type Box struct {
	count int
	Initem Item
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (t Box) Cost() float64 {
	return float64(t.count * 1000)
}

func main(){
	shirt := Item{name:"Men's Slim-Fit Shirt", price: 25000, quantity:3}
	shirtbox := Box{count:5,Initem:shirt}
	fmt.Println(shirtbox.Cost())
	fmt.Println(shirtbox.Initem.Cost())
}
