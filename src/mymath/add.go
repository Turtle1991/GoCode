package mymath

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("dp in mymath...")
}

func Jia(a int, b int) (re int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!!!!!")
		return
	}
	return a + b, nil
}
