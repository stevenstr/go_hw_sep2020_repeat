/*
* Author: Stefam
* Last change: 09.21.2020
* Task: Home work 2.1 Factorial
 */

package factori

//Factorial function
func Factorial(i uint) uint {
	if i == 0 {
		return 1
	}

	var a, result uint = 1, 1
	for a <= i {
		result *= a
		a++
	}
	return result
}
