/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Write reverse([]int64) ([]int64, error) function that returns the copy of the original
* slice in reverse order. The type of elements is int64.
 */

package main

import (
	"fmt"

	"domain/reverse"
)

func main() {

	sl := []int64{1, 2, 5, 15}
	fmt.Println(reverse.Reverse(sl))
	sl1 := []int64{}
	fmt.Println(reverse.Reverse(sl1))
	var sl2 []int64
	fmt.Println(reverse.Reverse(sl2))
}
