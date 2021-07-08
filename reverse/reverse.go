/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Write reverse([]int64) ([]int64, error) function that returns the copy of the original
* slice in reverse order. The type of elements is int64.
 */

package reverse

import "fmt"

//Reverse function Reverse([]int64) ([]int64, error) function that returns the copy of the original slice in reverse order. The type of elements is int64.
func Reverse(sl []int64) ([]int64, error) {

	switch {
	case sl == nil:
		s := []int64{}
		return s, fmt.Errorf("Slice are nil")
	case len(sl) == 0:
		s := []int64{}
		return s, fmt.Errorf("Are empty")
	}
	rev := make([]int64, len(sl))
	revStartInd := 0

	for i := len(sl) - 1; i >= 0; i-- {
		rev[revStartInd] = sl[i]
		revStartInd++
	}

	return rev, nil
}
