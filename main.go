/*
* Author: Stefan
* Last change: 09.20.2020
* Task: Class work 2 lecture
 */

package main

//"fmt"
//"math"

func main() {
	//ex 1 HW
	//fmt.Println("Hello, world!")

	//ex 2 Variables
	/*
		var a int16 = 123 + 321
		var b int64
		var c = 7
		d := 55

		fmt.Println(a, b, c, d)
		fmt.Printf("%v %T |%v %T |%v %T |%v %T \n\r\n", a, a, b, b, c, c, d, d)
	*/

	//ex 3 default values
	/*
		var (
			a   int
			b   float32
			s   string
			by  byte
			sl  []int
			arr [5]float64
			mp  map[int]string
			bl  bool
		)

		fmt.Printf("%v %T \n\r\n", a, a)
		fmt.Printf("%v %T \n\r\n", b, b)
		fmt.Printf("%v %T \n\r\n", s, s)
		fmt.Printf("%v %T \n\r\n", by, by)
		fmt.Printf("%v %T \n\r\n", sl, sl)
		fmt.Printf("%v %T \n\r\n", arr, arr)
		fmt.Printf("%v %T \n\r\n", mp, mp)
		fmt.Printf("%v %T \n\r\n", bl, bl)
	*/

	//ex 4 var declaration example
	/*
		var name string = "Ivan"
		var (
			//default zero
			age           int
			i             int = 5
			//auto type int and string
			amount, month     = 7, "January"
		)

		fmt.Println(name, age, i, month, amount)
		// %T - type of variable
		fmt.Printf("%T\n", month)
		fmt.Printf("%T\n", amount)
	*/

	//ex 5 Short variable declaration (need to fix)
	/*
		age := 22
		fmt.Println(age)

		cat := "Murzik"
		i, mouse := 1, "Jerry"
		fmt.Println(cat, i, mouse)

		cat = "Tom"
		fmt.Println(cat, i, mouse)

		cat, mouse = "Tom", "NoName"
		fmt.Println(cat, i, mouse)

		cat, newVar, mouse := "Tom", 1, "Jerry"
		fmt.Println(cat, i, mouse, newVar)
	*/

	//ex 6 If-else
	/*
		a := 5

		if a < 0 || a > 3 {
			fmt.Println("foo")
		} else {
			fmt.Println("bar")
		}
	*/

	//ex 7 Loops
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}

		i := 0
		for i < 5 {
			fmt.Println(i)
			i++
		}

		for {
			fmt.Println(i)
			if i == 10 {
				break
			}
			i++
		}
	*/

	//ex 8 Non-fixed size of integer
	/*
		var (
			a int
			b int32
			c int64
		)

		// type missmatched
		//fmt.Println(a == b)
		//fmt.Println(a == c)
		fmt.Println(a == int(b))
		fmt.Println(a == int(c))
	*/

	//ex 9 Integer operators
	/*
		a := 1
		a += 5
		fmt.Printf("%d\n", a) //6
		fmt.Printf("%b\n", a) //110

		a--
		fmt.Println(a) //5

		a++
		fmt.Println(a) //6

		var b uint8 = 255
		fmt.Println(b) //255
		b++
		fmt.Println(b) //0
	*/

	//ex 10 Convert types
	/*
		var apple int32 = 1
		var water int16 = 2

		var compote int = int(apple) + int(water)

		fmt.Printf("%v %T \n\r\n", apple, apple)
		fmt.Printf("%v %T \n\r\n", water, water)
		fmt.Printf("%v %T \n\r\n", compote, compote)
	*/

	//ex 11 Substring operation
	/*
		s := "hello, world!"
		hello := s[:5]
		world := s[7:]

		fmt.Println(s, hello, world)
	*/

	//ex 12 Deals with bytes
	/*
		hello := "п 学"
		fmt.Println(hello)

		fmt.Println(len(hello))                  //6 byte
		fmt.Println(len("п"))                    // 2 byte
		fmt.Println(len("学"))                    // 3 byte
		fmt.Println(utf8.RuneCountInString("学")) // one rune

		//run by rune
		for i, r := range hello {
			fmt.Printf("i = %v, %[2]T - %[2]v\n", i, r)
		}

		fmt.Println()

		//run by byte
		for i := 0; i < len(hello); i++ {
			fmt.Printf("i = %v, %[2]T - %[2]v\n", i, hello[i])
		}
	*/

	//ex 13 CUstom types
	/*
		type Minutes int
		type Hours int

		var a int = 20
		minutes := Minutes(20)
		hours := Hours(2)

		fmt.Printf("%v %T \n\r\n", a, a)
		fmt.Printf("%v %T \n\r\n", minutes, minutes)
		fmt.Printf("%v %T \n\r\n", hours, hours)
	*/

	//ex 14 Constants
	/*
		const name string = "Go"

		fmt.Println(name)

		const (
			e  = 2.7
			pi = 3.14
		)

		fmt.Println(e, pi)

		const (
			a = 1
			b
			c = 2.1
			d
		)

		fmt.Printf("%v %T | %v %T | %v %T | %v %T \n\r\n", a, a, b, b, c, c, d, d)
	*/

	//ex 15 iota
	/*
		const (
			mon = iota
			tues
			wedn
			thur
			er
			sut
			sun
			b
			c = iota + 3
			d
			e
		)

		fmt.Println(mon, tues, wedn, thur, er, sut, sun, b, c, d, e)
	*/

	//ex 16 Function examples
	/*
		fmt.Printf("%T\n\r\n", add)
		fmt.Printf("%T\n\r\n", sub)
		fmt.Printf("%T\n\r\n", first)
		fmt.Printf("%T\n\r\n", zero)
	*/

	//ex 17 Recursion
	/*
		fmt.Println(factorial(5))
	*/

	//ex 18 Multiple return values
	/*
		sum, mult := getValue(5, 6)
		fmt.Println(sum, mult)
	*/

	//ex 19 Multiple return values (need to implement)
	/*
		var (
			sum int
			ok  bool
		)

		if sum, ok = increaseEvenNumber(6); ok {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
		fmt.Println(sum)
	*/

	//ex 20 defer
	/*
		defer fmt.Println("Bye!")
		fmt.Println("Hello!")
	*/

	//ex 21 defer stack
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}

		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
		}
	*/

	//ex 22 Structure example
	/*
		type Book struct {
			ID          int
			Title       string
			Author      string
			Publication time.Time
			Pages       uint
		}

		var HArry Book
		HArry.ID = 1
		HArry.ID += 500

		fmt.Println(HArry.ID)
		fmt.Println(HArry.Pages)
	*/

	//ex 23 Pointer to structure
	/*
		type Book struct {
			ID          int
			Title       string
			Author      string
			Publication time.Time
			Pages       uint
		}

		var a Book
		a.ID = 5
		p := &a

		fmt.Println(a, p)
	*/

	//ex 24 Struct literals
	/*

		type Person struct {
			id   int
			name string
			Age  uint8
		}

		Pike := Person{1, "Rob", 63}
		Tompson := Person{name: "Ken", Age: 76}

		fmt.Println(Pike)
		fmt.Println(Tompson)
	*/

	//ex 25 Struct embedding
	/*
		type Engine struct {
			Power int
			Year  int
		}

		type Car struct {
			engine Engine
			Title  string
			Year   int
		}

		audi := Car{Title: "Audi", Year: 2010, engine: Engine{Power: 450, Year: 2009}}
		bmw := Car{Engine{500, 2005}, "BMW", 2010}
		opel := Car{Engine{400, 2010}, "Opel", 2011}

		opel.engine.Power = 200

		fmt.Println(audi)
		fmt.Println(bmw)
		fmt.Println(opel)
	*/

	//ex 26 Struct embedding (anonymous fields)
	/*
		type Engine struct {
			Power int
			Year  int
		}

		type Car struct {
			Engine
			Title string
			Year  int
		}

		opel := Car{Engine{400, 2010}, "Opel", 2011}

		opel.Power = 200

		fmt.Println(opel.Year)
		fmt.Println(opel.Engine.Year)
	*/

	//ex 27 What a method ?
	/*
		p := Point{1, 1}
		q := Point{2, 5}
		fmt.Println(Distance(p, q))
		fmt.Println(p.Distance(q))
	*/
}

//ex 16 Function examples
/*
func add(x int, y int) int { return x + y }
func sub(x, y int) (z int) { z = x - y; return z }
func first(x, _ int) int   { return x }
func zero(int, int) int    { return 0 }
*/

//ex 17 Recursion
/*
func factorial(i uint) uint {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}
*/

//ex 18 Multiple return values
/*
func getValue(x, y int) (int, int) {
	return x + y, x * y
}
*/

//ex 19 Multiple return values (need to implement)
/*
func increaseEvenNumber(x int) (int, bool) {
	if x%2 == 0 {
		return x * 2, true
	}
	return x, false
}
*/

//ex 27 What a method ?
/*

//Point struct
type Point struct{ X, Y float64 }

// Distance method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance traditional func
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
*/
