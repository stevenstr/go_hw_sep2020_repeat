/*
* Author: Stefan
* Last change: 09.26.2020
* Task:Write max([]string) string function that returns the longest word from
* the slice of strings (the first if there are more than one).
 */

package maximum

import (
	"unicode/utf8"
)

//Max function that returns the longest word from the slice of strings (the first if there are more than one).
func Max(s []string) string {

	switch {
	case s == nil:
		return "slice are nil!"
	case len(s) == 0:
		return "no elements!"
	case len(s) == 1:
		return s[0]
	}

	l, tmp := 0, 0
	var result string

	for i, v := range s {
		tmp = utf8.RuneCountInString(v)
		if l < tmp {
			l = tmp
			result = s[i]
		}
	}
	return result
}
