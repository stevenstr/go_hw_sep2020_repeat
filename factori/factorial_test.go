package factori

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
* go test -coverprofile=cover
* go tool cover -html=cover -o coverout.html
 */
//Test_Factorial function
func Test_Factorial(t *testing.T) {
	res := Factorial(0)
	assert.Equal(t, uint(1), res)

	res = Factorial(5)
	assert.Equal(t, uint(120), res)

	res = Factorial(1)
	assert.Equal(t, uint(1), res)
}
