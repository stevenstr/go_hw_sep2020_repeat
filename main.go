/*
* Author: Stefan
* Last change: 09.21.2020
* Task: Home work 2.2 methods
 */

package main

import (
	"fmt"
	"module/domain"
)

func main() {
	s := domain.Square{domain.Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
