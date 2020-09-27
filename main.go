/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Implement Square and Circle structures which implements
	Figure interface.
	Add test cases for Perimeter / Area calculation.
*/

package main

import (
	"fmt"

	"module/figure"
)

func main() {
	var s figure.Figure
	ss := figure.Square{A: 3}

	s = ss

	var c figure.Figure
	cc := figure.Circle{R: 3}
	c = cc

	fmt.Println(s.Area(), s.Perimeter())
	fmt.Println(c.Area(), c.Perimeter())
}
