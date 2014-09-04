package main

import (
	"fmt"
)

type PersonInfo struct {
	ID string
	// Name    string
	// Address string
}

func main() {
	/*num := 1
	num2 := -1
	switch {
	case num2 == 0:
		fmt.Println("num < 0")
	case num > 0:
		fallthrough
	case num < -1:
		fmt.Println("nonono")
	}*/

	/*sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)*/

	//无限循环
	/*sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)*/

	/*a := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for _, v := range a {
		fmt.Print(v, " ")
	}*/

	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
