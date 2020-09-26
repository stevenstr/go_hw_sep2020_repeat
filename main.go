/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Implement average([…]int) float64
function that returns an average value of array (sum / N)
*/

package main

import (
	"fmt"
	"module/average"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average.Average(sl...)) //3.5

	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(average.Average(sl1...)) //5.5

	sl2 := []int{1, 2, 3, 4}
	fmt.Println(average.Average(sl2...)) //2.5

	sl3 := []int{}
	fmt.Println(average.Average(sl3...)) //0.0

	var sl4 []int
	fmt.Println(average.Average(sl4...)) //0.0

}
