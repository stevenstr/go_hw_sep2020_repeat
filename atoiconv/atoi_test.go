package atoiconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDone(t *testing.T) {
	result, err := myStrToInt("1234")
	assert.Equal(t, result, 1234, "OK")
	assert.Nil(t, err)
}

func TestMix(t *testing.T) {
	result, err := myStrToInt("1s234")
	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

func TestOverSize(t *testing.T) {
	result, err := myStrToInt("999999999999999999999999999")
	assert.Equal(t, result, 9223372036854775807, "should be equal")
	assert.Error(t, err, "should be error")
}

func TestEmpty(t *testing.T) {
	result, err := myStrToInt("")
	assert.Equal(t, result, 0, "should be equal")
	assert.Error(t, err, "should be error")
}

func BenchmarkMyStrToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = myStrToInt("1")
	}
}
