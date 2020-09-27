/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Implement printSorted(map[int]string)function that prints map values sorted in order of increasing keys
 */

package main

import "domain/printsorted"

func main() {
	mp := map[int]string{2: "a", 0: "b", 1: "c"}
	mpSec := map[int]string{10: "aa", 0: "bb", 500: "cc"}

	printsorted.PrintSorted(mp)
	printsorted.PrintSorted(mpSec)
}
