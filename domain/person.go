/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Implement sort.Interface interface for People type.
	Sort people in order of increasing age.
	If two people have the same age - sort them by name.
	Add test cases.
*/

package domain

import (
	"fmt"
	"time"
)

//Interface interface
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//Person structure
type Person struct {
	FirstName string
	LastName  string
	BirthDay  time.Time
}

//People structure
type People []Person

//Len function
func (p People) Len() int {
	return len(p)
}

//Less function
func (p People) Less(i, j int) bool {
	switch {
	case p[i].BirthDay.Sub(p[j].BirthDay) > 0:
		return true
	case p[i].BirthDay.Sub(p[j].BirthDay) < 0:
		return false
	case p[i].FirstName < p[j].FirstName:
		return true
	case p[i].FirstName > p[j].FirstName:
		return false
	}
	return p[i].LastName < p[j].LastName
}

//Swap function
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//Show function
func (p People) Show(s string) {
	fmt.Println(s)
	for _, q := range p {
		fmt.Println(q.FirstName, q.LastName, q.BirthDay)
	}
	fmt.Println()
}
