/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Class work 5 unit tests
 */

package firstwhitebox

import "sort"

//CalcMarcCake function
func CalcMarcCake(d []int64) (s int64) {
	sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })
	var c int64 = 1
	for i := 0; i < len(d); i++ {
		s += d[i] * c
		c = c << 1
	}
	return s
}
