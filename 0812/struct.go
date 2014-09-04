package main

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main() {
	/*rect := new(Rect)
	rect.width = 1
	rect.height = 4*/
	rect := NewRect(0, 0, 100, 0.5)
	rect2 := NewRect(10, 2, 0, 0)
	area := rect.Area()
	fmt.Println(area)
	fmt.Println(rect.x, rect.y)
	fmt.Println(rect2.x, rect2.y)
}
