package main

import (
	"fmt"
)

/*func Turtle(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}*/

func Turtle(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("unknown")
		}
	}
}

func main() {
	/*a := 1
	b := 2
	Turtle(a, b)*/

	/*a := []int{1, 2, 3}
	Turtle(a...)*/

	a := 1.23
	b := "turtle"
	Turtle(a, b)
}
