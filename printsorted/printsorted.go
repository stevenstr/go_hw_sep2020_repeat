/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Implement printSorted(map[int]string)function that prints map values sorted in order of increasing keys
 */

package printsorted

import (
	"fmt"
	"sort"
)

//PrintSorted function that prints map values sorted in order of increasing keys
func PrintSorted(mp map[int]string) error {

	//checks
	switch {
	case mp == nil:
		return fmt.Errorf("Map are nil")
	case len(mp) == 0:
		return fmt.Errorf("Map are empty")
	}
	//vars for results
	resSlice := []int{}
	resSliceStr := []string{}

	//get values of int
	for k := range mp {
		resSlice = append(resSlice, k)
	}

	//sort indexes for result string slice
	sort.Ints(resSlice)

	//filled our result string slice
	for _, v := range resSlice {
		resSliceStr = append(resSliceStr, mp[v])
	}

	//print result
	//_, err :=
	fmt.Println(resSliceStr)
	/*if err != nil {
		return fmt.Errorf("Problem in last println 42 line")
	}*/
	return nil
}
