/*

Implement string to int converter, like
● func myStrToInt(s str) (int, error){}
● Cover it with tests
*/

package atoiconv

import (
	"strconv"
)

func myStrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
