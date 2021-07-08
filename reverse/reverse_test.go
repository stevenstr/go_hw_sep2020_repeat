/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Write reverse([]int64) ([]int64, error) function that returns the copy of the original
* slice in reverse order. The type of elements is int64.
 */

package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	sl := []int64{1, 2, 5, 15}
	result, err := Reverse(sl)
	assert.Equal(t, []int64{15, 5, 2, 1}, result)
	assert.Nil(t, err, nil)

	sl1 := []int64{}
	result, err = Reverse(sl1)
	assert.Equal(t, []int64{}, result)
	assert.EqualError(t, err, "Are empty")

	var sl2 []int64
	result, err = Reverse(sl2)
	assert.Equal(t, []int64{}, result)
	assert.EqualError(t, err, "Slice are nil")
}
