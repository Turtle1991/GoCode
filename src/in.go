package main

import "fmt"
import (
	"mymath"
)

func init() {
	fmt.Println("run main ")
}

func main() {
	a := 1
	b := 5
	c, err := mymath.Jia(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
