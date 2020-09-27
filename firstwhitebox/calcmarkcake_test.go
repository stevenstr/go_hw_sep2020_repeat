/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Class work 5 unit tests
 */

package firstwhitebox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcMarcCake_1(t *testing.T) {
	ins := []int64{1, 3, 2}
	var out int64 = 11
	assert.Equal(t, out, CalcMarcCake(ins))
}

func TestCalcMarcCake_2(t *testing.T) {
	ins := []int64{7, 4, 9, 6}
	var out int64 = 79
	assert.Equal(t, out, CalcMarcCake(ins))
}
