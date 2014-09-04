package main

import (
	"fmt"
)

func main() {
	/*array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:1]
	fmt.Println("The values of array is:")
	for _, v := range array {
		fmt.Print(v, " ")
	}
	fmt.Println("\nThe values of slice is:")
	for _, v := range slice {
		fmt.Print(v, " ")
	}*/

	/*slice := make([]int, 5, 10)
	slice2 := []int{4, 5, 6}
	slice = append(slice, slice2...)
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	l := len(slice)
	fmt.Println("\n", l)
	c := cap(slice)
	fmt.Println(c)*/

	/*oneSlice := make([]int, 4, 10)
	newSlice := oneSlice[:9]
	for _, v := range newSlice {
		fmt.Print(v, " ")
		}*/

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	// copy(slice2, slice1)
	copy(slice1, slice2)
	for _, v := range slice1 {
		fmt.Print(v, " ")
	}
}
