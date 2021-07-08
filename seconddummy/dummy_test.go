/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Class work 5 unit tests
 */

package seconddummy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	k, err := dummy(11)
	assert.Equal(t, 1, k)
	assert.NotNil(t, err)
}

func TestDummy_2(t *testing.T) {
	k, err := dummy(9)
	assert.Equal(t, 0, k)
	assert.Nil(t, err)
}
