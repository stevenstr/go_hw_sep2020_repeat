/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Implement average([…]int) float64
function that returns an average value of array (sum / N)
*/

package average

//Average function function that returns an average value of array (sum / N)
func Average(sl ...int) float64 {

	switch {
	case len(sl) == 0:
		return 0.0
	}

	count := 0
	for _, v := range sl {
		count += v
	}

	result := float64(count) / float64(len(sl))
	return result
}
