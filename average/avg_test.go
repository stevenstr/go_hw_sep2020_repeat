/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Implement average([…]int) float64
function that returns an average value of array (sum / N)
*/

package average

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6}
	result := Average(sl...) //3.5
	assert.Equal(t, 3.5, result)

	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = Average(sl1...) //5.5
	assert.Equal(t, 5.5, result)

	sl2 := []int{1, 2, 3, 4}
	result = Average(sl2...) //2.5
	assert.Equal(t, 2.5, result)

	sl3 := []int{}
	result = Average(sl3...) //0.0
	assert.Equal(t, float64(0), result)

	var sl4 []int
	result = Average(sl4...) //0.0
	assert.Equal(t, float64(0), result)
}
