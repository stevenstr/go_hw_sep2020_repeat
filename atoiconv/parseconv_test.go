package atoiconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDone2(t *testing.T) {
	result, err := myStrToInt2("1234")
	assert.Equal(t, result, int64(1234), "OK")
	assert.Nil(t, err)
}

func TestMix2(t *testing.T) {
	result, err := myStrToInt2("1s234")
	assert.Equal(t, result, int64(0), "should be equal")
	assert.Error(t, err, "should be error")
}

func TestOverSize2(t *testing.T) {
	result, err := myStrToInt2("999999999999999999999999999")
	assert.Equal(t, result, int64(9223372036854775807), "should be equal")
	assert.Error(t, err, "should be error")
}

func TestEmpty2(t *testing.T) {
	result, err := myStrToInt2("")
	assert.Equal(t, result, int64(0), "should be equal")
	assert.Error(t, err, "should be error")
}

func BenchmarkMyStrToInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = myStrToInt2("1")
	}
}
