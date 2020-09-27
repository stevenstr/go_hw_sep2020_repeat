/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Class work 5 unit tests
 */

package seconddummy

import "errors"

func dummy(k int) (int, error) {
	if k > 10 {
		return 1, errors.New("must be gt 10")
	}
	return 0, nil
}
