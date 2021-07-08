/*
* Author: Stefan
* Last change: 09.20.2020
* Task: Home work 1.1 - write hello worl program
 */

package main

import "fmt"

const (
	sep  = ","
	sep2 = "!"
)

func main() {
	var hello string = "Hello"
	world := "world"

	result := hello + sep + world + sep2

	fmt.Println(result)

	for i, v := range result {
		fmt.Printf("index: %v | rune: %c \n\r\n", i, v)
	}
}
