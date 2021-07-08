/*
* Author: Stefan
* Last change: 09.27.2020
* Task: Class work 4 interfaces, dependency management
 */

package main

func main() {

	//ex 1 Satisfied  implicity
	/*
		d := dog{}
		d.voice()
	*/

	//ex 2 Interfaces as contracts
	/*
		var c ByteCounter
		c.Write([]byte("hello"))
		fmt.Println(c) //5
		c = 0 //reset c
		fmt.Fprintf(&c, "hello %s", "world")//11
		fmt.Println(c)
	*/

	//ex 3 fmt.Stringer
	/*
		s := sl{1, 2, 3}
		fmt.Println(s)

		s1 := []int{1, 2, 3}
		fmt.Println(s1)
	*/

	//ex 4 Method set
	/*
		i := item{}

		i.A()
		i.B()

		p := &item{}
		p.A()
		p.B()

		item{}.A() //ok
		// item{}.B() // cannot take address
	*/

	//ex 5 Interface wrappers
	/*
		os.Stdout.Write([]byte("test")) //ok
		os.Stdout.Close()               //ok

		var w io.Writer //only writer interface

		w = os.Stdout
		w.Write([]byte("test")) //ok
		//w.Close() //dont implement Closer interface
	*/

	//ex 6 Empty inteface
	/*
		var any interface{}
		any = true
		fmt.Printf("%T\n\r", any)

		any = "Golang"
		fmt.Printf("%T\n\r", any)

		any = []int{1, 2, 3, 4}
		fmt.Printf("%T\n\r", any)
	*/

	//ex 7 Assert interface relations
	/*
		var animals animal = cat{}
		animals.voice()

		animals = dog{}
		animals.voice()
	*/

	//ex 8 A concrete type may satisfy many unrelated interfaces
	/*
		type Artifact interface {
			Title() string
			Creators() []string
			Created() time.Time
		}
		type Text interface {
			Pages() uint
			Words() uint
		}
		type Audio interface {
			Stream() (io.ReadCloser, error)
			RunningTime() time.Duration
		}
	*/

	//ex 9 Interface value
	/*
		var w io.Writer
		//interface, dynamic type nil, dynamic value nil
		fmt.Printf("value: %v | type: %T \n\r\n", w, w)

		w = os.Stdout // - it is a pointer to os.File
		//interface, dynamic type *os.File, dynamic value address in memory
		fmt.Printf("value: %v | type: %T \n\r\n", w, w)

		w = new(bytes.Buffer) // - new return a pointer to bytes.Buffer
		//interface, dynamic type *bytes.Buffer, dynamic value empty
		fmt.Printf("value: %v | type: %T \n\r\n", w, w)

		w = nil
		//interface, dynamic type nil, dynamic value nil
		fmt.Printf("value: %v | type: %T \n\r\n", w, w)

		var inter interface{}
		inter = 1
		//interface, dynamic type int, dynamic value 1
		fmt.Printf("value: %v | type: %T \n\r\n", inter, inter)

		inter = []int{1, 2, 3}
		//interface, dynamic type []int, dynamic value {1, 2, 3}
		fmt.Printf("value: %v | type: %T \n\r\n", inter, inter)
	*/

	//ex 10 Interface comparison
	/*
		var x, y interface{}
		fmt.Println(x == y)

		x = [2]int{1, 2}
		y = [2]int{1, 2}
		fmt.Println(x == y)

		x = [2]int{1, 2}
		y = [2]int{1, 1}
		fmt.Println(x == y)
	*/

	//ex 11 How to get the interface dynamic type
	/*
		var (
			x interface{}
			y interface{} = [2]int{1, 2}
			z interface{} = []int{1, 2}
		)

		fmt.Printf("%T\n\r", x)
		fmt.Printf("%T\n\r", y)
		fmt.Printf("%T\n\r", z)
	*/

	//ex 12 Type assertion
	/*
		//voicer implemet only one method voice
		//dog is a object implement 2 method but run was hidden
		var v voicer = dog{}

		v.voice()
		//v.run()

		a := v.(dog) //type assertion
		a.voice()
		a.run()
	*/

	//ex 13
	/*
		type Animal interface {
			voice()
		}
		var w io.Writer
		w = os.Stdout

		rw, ok := w.(io.ReadWriter)
		fmt.Println(ok)
		fmt.Printf("%T\n", rw)

		//extend our interface
		t, ok := w.(Animal) // no panic
		fmt.Println(ok)
		fmt.Printf("%T\n", t)
	*/

	//ex 14 Querying behaviour
	/*
		var d = dog{}
		//put d into empty interface
		var emptyInter interface{} = d
		//extand empty interface
		if resInter, ok := emptyInter.(dog); ok {
			resInter.voice()
		}

		//use animal interface
		var animalInter animal
		animalInter = d
		animalInter.voice()
	*/

	//ex 15 Type switches
	/*

		var a interface{} = true
		var b interface{} = "test"
		var c interface{}
		fmt.Println(dynamicValue(a)) // TRUE
		fmt.Println(dynamicValue(b)) // Unknown type
		fmt.Println(dynamicValue(c)) // NULL
	*/
}

//ex 1 Satisfied  implicity
/*
type dog struct{}

func (_ dog) voice() {
	fmt.Println("Gav!")
}

type animal interface {
	voice()
}
*/

//ex 2 Interfaces as contracts
/*
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
*/

//ex 3 fmt.Stringer
/*
type sl []int

//redefine string method for stringer interface
func (s sl) String() string {
	result := ""
	for i, v := range s {
		result += fmt.Sprintf("%d - %d\n", i, v)
	}
	return result
}
*/

//ex 4 Method set
/*
type item struct{}

func (i item) A() {
	fmt.Println("A")
}

func (i *item) B() {
	fmt.Println("B")
}
*/

//ex 7 Assert interface relations
/*
type cat struct{}

type dog struct{}

type animal interface {
	voice()
}

func (_ cat) voice() {
	fmt.Println("Miay!")
}

func (_ dog) voice() {
	fmt.Println("Gav!")
}
*/

//ex 12 Type assertion
/*
type dog struct{}

func (d dog) voice() {
	fmt.Println("Gav!")
}

func (d dog) run() {
	fmt.Println("RUN!")
}

type voicer interface {
	voice()
}
*/

//ex 14 Querying behaviour
/*
type dog struct{}

func (d dog) voice() {
	fmt.Println("Gav!!!")
}

type animal interface {
	voice()
}
*/

//ex 15 Type switches
/*
func dynamicValue(x interface{}) string {
	switch x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return "INT, UINT"
	case float64:
		return "FLOAT64"
	case bool:
		return "BOOL"
	default:
		return "Unknown type"
	}
}
*/
