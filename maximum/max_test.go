/*
* Author: Stefan
* Last change: 09.26.2020
* Task:Write max([]string) string function that returns the longest word from
* the slice of strings (the first if there are more than one).
 */

package maximum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	s := []string{"one", "two", "three", "third"}
	result := Max(s)
	assert.Equal(t, "three", result)

	s1 := []string{"one", "two"}
	result = Max(s1)
	assert.Equal(t, "one", result)

	var s2 []string
	result = Max(s2)
	assert.Equal(t, "slice are nil!", result)

	s3 := []string{}
	result = Max(s3)
	assert.Equal(t, "no elements!", result)

	s4 := []string{"heheh"}
	result = Max(s4)
	assert.Equal(t, "heheh", result)
}
