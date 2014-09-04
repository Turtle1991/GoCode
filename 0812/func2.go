package main

import (
	"fmt"
)

func main() {
	/*add := func(a int, b int) int {
		return a + b
	}
	a := 1
	b := 2
	result := add(a, b)
	fmt.Println(result)*/

	/*a := 1
	b := 3
	result := func(a int, b int) int {
		return a + b
	}(a, b)
	fmt.Println(result)*/

	var j int = 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Printf("i is %d, j is %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
