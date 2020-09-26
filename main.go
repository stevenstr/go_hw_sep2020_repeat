/*
* Author: Stefan
* Last change: 09.26.2020
* Task:Write max([]string) string function that returns the longest word from
* the slice of strings (the first if there are more than one).
 */

package main

import (
	"fmt"

	"module/maximum"
)

func main() {
	s := []string{"one", "two", "three", "third"}
	fmt.Println(maximum.Max(s))

	s1 := []string{"one", "two"}
	fmt.Println(maximum.Max(s1))
}
