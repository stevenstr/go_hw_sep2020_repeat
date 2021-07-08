/*
* Author: Stefan
* Last changes: 09.26.2020
* Task: Class work Lecture 3 Array, slice, map
 */

package main

func main() {

	//ex 1 Simple example
	/*
		var x [5]int
		c := [5]int{1, 2, 3, 4, 5}
		v := [3]int{
			1,
			2,
			3,
		}

		fmt.Println(x)
		fmt.Println(c)
		fmt.Println(v)
	*/

	//ex 2 Automatic array length declaration
	/*
		x := [...]int{1, 2, 3, 4}
		fmt.Println(x, len(x))
		for i, v := range x {
			fmt.Println(i, v)
		}
	*/

	//ex 3 Array literals p.1
	/*
		arr := [5]int{6, 7, 8, 9, 10}
		arr2 := [5]int{0: 6, 7, 8, 9, 10}
		arr3 := [5]int{0: 6, 1: 7, 2: 8, 3: 9, 4: 10}

		fmt.Println("Array 1:", arr, "Len:", len(arr))
		fmt.Println("Array 2:", arr2, "Len:", len(arr2))
		fmt.Println("Array 3:", arr3, "Len:", len(arr3))
	*/

	//ex 4 Array literals p.2
	/*
		arr := [5]int{3: 4, 5, 0: 1}
		fmt.Println(arr)
		for i, v := range arr {
			fmt.Println(i, v)
		}
	*/

	//ex 5 Size is a part of type
	/*
		var x, y [5]int
		//var c [6]int

		fmt.Println(x == y)
		//fmt.Println(x == c)// error missmatched type
	*/

	//ex 6 Array iteration
	/*
		arr := [3]string{
			"kolya",
			"ivan",
			"lesha",
		}

		arr1 := [...]string{
			"kolya",
			"ivan",
			"lesha",
		}
		for i, v := range arr {
			fmt.Println(i, v)
		}

		for i, v := range arr1 {
			fmt.Println(i, v)
		}

		for i := 0; i < len(arr); i++ {
			fmt.Println(i, arr[i])
		}

		for _, v := range arr {
			fmt.Println(v)
		}
	*/

	//ex 7 Passed by value
	/*
		arr := [3]int{1, 2, 3}
		fmt.Println("original array:", arr)
		fmt.Println(arr)
		changeFirst(arr)
		fmt.Println("after chage:", arr)
	*/

	//ex 8 Multi-dimensional arrays
	/*
		arr := [3][2]int{
			[2]int{1, 2},
			[2]int{3, 4},
			[2]int{4, 6},
		}

		arr1 := [2][2]int{
			[2]int{1, 2},
			[2]int{3, 4},
		}

		fmt.Printf("%T | %v\n\r\n", arr, arr)
		fmt.Printf("%T | %v\n\r\n", arr1, arr)
	*/

	//ex 9 Slice declaration
	/*
		sl := []int{}
		var sl1 []int
		sl2 := make([]int, 3)
		sl3 := make([]int, 3, 7)
		sl4 := new([]int)

		fmt.Printf("%T | %v | sl == nil? - %v\n\r\n", sl, sl, sl == nil)
		fmt.Printf("%T | %v | sl == nil? - %v\n\r\n", sl1, sl1, sl1 == nil)
		fmt.Printf("%T | %v | sl == nil? - %v\n\r\n", sl2, sl2, sl2 == nil)
		fmt.Printf("%T | %v | sl == nil? - %v\n\r\n", sl3, sl3, sl3 == nil)
		fmt.Printf("%T | %v | sl == nil? - %v\n\r\n", sl4, sl4, sl4 == nil)
	*/

	//ex 10 Built-in function make
	/*
		s := make([]int, 2, 5)
		fmt.Println(s, len(s), cap(s))
		s1 := make([]int, 6)
		fmt.Println(s1, len(s1), cap(s1))
		s1 = s1[:4]
		fmt.Println(s1, len(s1), cap(s1))
		s1 = s1[:cap(s1)]
		fmt.Println(s1, len(s1), cap(s1))
		s1 = s1[2:]
		fmt.Println(s1, len(s1), cap(s1))
	*/

	//ex 11 slicing
	/*
		months := []string{
			"jan",
			"feb",
			"mar",
			"apr",
			"may",
			"jun",
		}

		q1 := months[:4]
		q2 := months[3:]

		fmt.Println(months, "len:", len(months), "cap:", cap(months))
		fmt.Println(q1, "len:", len(q1), "cap:", cap(q1))
		fmt.Println(q2, "len:", len(q2), "cap:", cap(q2))

		q2[0] = "CHANGED"
		fmt.Println(months, "len:", len(months), "cap:", cap(months))
		fmt.Println(q1, "len:", len(q1), "cap:", cap(q1))
		fmt.Println(q2, "len:", len(q2), "cap:", cap(q2))

		q1 = q1[:cap(q1)]

		fmt.Println(months, "len:", len(months), "cap:", cap(months))
		fmt.Println(q1, "len:", len(q1), "cap:", cap(q1))
		fmt.Println(q2, "len:", len(q2), "cap:", cap(q2))
	*/

	//ex 12 Growing slices
	/*
		sl := []int{1, 2, 3}
		sl1 := []int{4, 5, 6}

		fmt.Printf("%v, len: %v, cap: %v\n\r\n", sl, len(sl), cap(sl))
		fmt.Printf("%v, len: %v, cap: %v\n\r\n", sl1, len(sl1), cap(sl1))

		sl = append(sl, sl1...)
		fmt.Printf("%v, len: %v, cap: %v\n\r\n", sl, len(sl), cap(sl))

		sl = append(sl, 7, 8, 9)
		fmt.Printf("%v, len: %v, cap: %v\n\r\n", sl, len(sl), cap(sl))
	*/

	//ex 13 Growing slices
	/*
		fmt.Println("original")
		a := [...]int{1, 2, 3, 4, 5}
		s := a[:3]
		fmt.Printf("a: %v, len: %v, cap: %v\n\r\n", a, len(a), cap(a))
		fmt.Printf("s: %v, len: %v, cap: %v\n\r\n", s, len(s), cap(s))

		fmt.Println("s1 is a s slice append 1")
		s1 := append(s, 1)
		fmt.Printf("a: %v, len: %v, cap: %v\n\r\n", a, len(a), cap(a))
		fmt.Printf("s: %v, len: %v, cap: %v\n\r\n", s, len(s), cap(s))
		fmt.Printf("s1: %v, len: %v, cap: %v\n\r\n", s1, len(s1), cap(s1))

		fmt.Println("s1 slice append 1")
		s1 = append(s1, 1)
		fmt.Printf("a: %v, len: %v, cap: %v\n\r\n", a, len(a), cap(a))
		fmt.Printf("s: %v, len: %v, cap: %v\n\r\n", s, len(s), cap(s))
		fmt.Printf("s1: %v, len: %v, cap: %v\n\r\n", s1, len(s1), cap(s1))

		fmt.Println("s1 slice append 1")
		s1 = append(s1, 1)
		fmt.Printf("a: %v, len: %v, cap: %v\n\r\n", a, len(a), cap(a))
		fmt.Printf("s: %v, len: %v, cap: %v\n\r\n", s, len(s), cap(s))
		fmt.Printf("s1: %v, len: %v, cap: %v\n\r\n", s1, len(s1), cap(s1))

		fmt.Println("cap \"a\" was eatten and \"s1\" now create new memory like 5*2 = 10 and next append do not affect on \"a\" slice")
		s1[3] = 5555
		fmt.Printf("a: %v, len: %v, cap: %v\n\r\n", a, len(a), cap(a))
		fmt.Printf("s: %v, len: %v, cap: %v\n\r\n", s, len(s), cap(s))
		fmt.Printf("s1: %v, len: %v, cap: %v\n\r\n", s1, len(s1), cap(s1))
	*/

	//ex 14 Growing slice experiment
	/*
		var x []int
		for i := 0; i < 65; i++ {
			x = append(x, 0)
			fmt.Printf("len: %v | cap: %v | Value: %v\r\n\r", len(x), cap(x), x)
		}
	*/

	//ex 15 Compare slices !!!only with nil
	/*
		var s []int
		fmt.Println(s == nil)
		var s1 = []int{}
		fmt.Println(s1 == nil)
		var s2 = []int(nil)
		fmt.Println(s2 == nil)
	*/

	//ex 16 Reference type ?
	/*
		sl := []int{1, 2, 3, 4}
		fmt.Println(sl)
		changeSlice(sl)
		fmt.Println(sl)
	*/

	//ex 17 Useful tricks
	/*
		var arr [5]int
		sl := arr[:]

		fmt.Println(arr, len(arr))
		fmt.Println(sl, len(sl), cap(sl))
	*/

	//ex 18 Maps
	/*
		var a map[string]int
		var b = map[string]int{"1": 1, "2": 2}

		e := make(map[string]int, 3)
		f := make(map[int]int)

		fmt.Printf("a %T | %v \n\r\n", a, a)
		fmt.Printf("b %T | %v \n\r\n", b, b)
		fmt.Printf("e %T | %v \n\r\n", e, e)
		fmt.Printf("f %T | %v \n\r\n", f, f)

		e["first"] = 1
		e["second"] = 2
		fmt.Printf("e %T | %v \n\r\n", e, e)

		first := e["first"]
		fmt.Println(first)
		third := e["third"]
		fmt.Println(third)

		third, ok := e["third"]
		if !ok {
			fmt.Println("third does not exist!")
		}

		first, ok = e["first"]
		if !ok {
			fmt.Println("first does not exist!")
		}

		fmt.Println(e, "len ", len(e))

		delete(e, "first")
		fmt.Println(e, "len ", len(e))
	*/

}

//ex 7 Passed by value
/*
func changeFirst(a [3]int) {
	a[1] = 555
	fmt.Println("inside func:", a)
}
*/

//ex 16 Reference type ?
/*
func changeSlice(sl []int) {
	if len(sl) > 0 {
		sl[0] = 999
	}
}
*/
