package atoiconv

import (
	"strconv"
)

func myStrToInt2(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
