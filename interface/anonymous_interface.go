package main

import "fmt"

type rect struct {width, heigth float64}

func (r rect) show(){
	fmt.Printf("width: %f, heigth: %f \n", r.width , r.heigth)
}

type circle struct {radius float64}

func (c circle) show() {
	fmt.Printf("radius: %f\n", c.radius)
}


func display(s interface { show()}) {
s.show()
}

func main(){
	r := rect{3,4}
	c := circle{2.5}

	display(r)
	display(c)
}
