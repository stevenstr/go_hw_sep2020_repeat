package domain

//Point structure
type Point struct {
	X, Y int
}

//Square structure
type Square struct {
	Start Point
	A     uint
}

//End method
func (s Square) End() (p Point) {
	p.X = s.Start.X + int(s.A)
	p.Y = s.Start.Y + int(s.A)
	return p
}

//Perimeter method
func (s Square) Perimeter() uint {
	perimeter := s.A * uint(4)
	return perimeter
}

//Area method
func (s Square) Area() uint {
	area := s.A * s.A
	return area
}
