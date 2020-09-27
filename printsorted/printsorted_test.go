/*
* Author: Stefan
* Last change: 09.26.2020
* Task: Implement printSorted(map[int]string)function that prints map values sorted in order of increasing keys
 */

package printsorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintSorted(t *testing.T) {

	//checks on nil error
	mp := map[int]string{2: "a", 0: "b", 1: "c"}
	err := PrintSorted(mp)
	assert.Nil(t, err, nil)

	mpSec := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	err = PrintSorted(mpSec)
	assert.Nil(t, err, nil)

	//check on nil map
	var mpNil map[int]string
	err = PrintSorted(mpNil)
	assert.EqualError(t, err, "Map are nil")

	//check on empty map
	var mpEmpty = map[int]string{}
	err = PrintSorted(mpEmpty)
	assert.EqualError(t, err, "Map are empty")
}
