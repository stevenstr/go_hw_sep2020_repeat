/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Implement Square and Circle structures which implements
	Figure interface.
	Add test cases for Perimeter / Area calculation.
*/

package figure

const pi float64 = 3.14159265359

//Figure interface
type Figure interface {
	Area() float64
	Perimeter() float64
}

//Square struct
type Square struct {
	A float64
}

//Area square method Methods for square
func (s Square) Area() float64 {
	return s.A * s.A
}

//Perimeter square method
func (s Square) Perimeter() float64 {
	return s.A * 4.0
}

//Circle struct
type Circle struct {
	R float64
}

//Area circle method Methods for Circle
func (c Circle) Area() float64 {
	return pi * (c.R * c.R)
}

//Perimeter circle method
func (c Circle) Perimeter() float64 {
	return 2.0 * pi * c.R
}
