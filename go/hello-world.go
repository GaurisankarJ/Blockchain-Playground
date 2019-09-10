// https://tour.golang.org/list
package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"time"
)

func add(x int, y int) int { // add(x, y int)
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 17 / 7
	y = sum - 7
	return
}

var c, python, java bool
var j, k int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeInference() {
	v := 42.4
	fmt.Printf("v is of type %T\n", v)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places
	// In other words, the binary number that is 1 followed by 100 zeroes
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func theForLoop() {
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// The init and post statements are optional
	sum := 1
	for sum < 10 { // for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) // Use fmt.Sprintf to format a string without printing it
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // Like for, the if statement can start with a short statement to execute before the condition
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func theSwitch() { // The break statement that is needed at the end of each case in those languages is provided automatically in Go
	// fmt.Print("Go runs on ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.\n", os)
	// }

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func theDefer() {
	// defer fmt.Println("world")

	// fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func thePointer() { // Unlike C, Go has no pointer arithmetic
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func theArray() {
	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// var s []int = primes[1:4]
	// fmt.Println(s)

	// Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func theSlice() {
	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

	// y := []int{2, 3, 5, 7, 11, 13}

	// y = y[1:4]
	// fmt.Println(y)

	// y = y[:2]
	// fmt.Println(y)

	// y = y[1:]
	// fmt.Println(y)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	// The length of a slice is the number of elements it contains
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func theMake() {
	a := make([]int, 5) // len(a)=5
	printSliceNew("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSliceNew("b", b)

	c := b[:2]
	printSliceNew("c", c)

	d := c[2:5]
	printSliceNew("d", d)
}

func printSliceNew(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func theSlicesOfSlices() {
	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func theAppend() {
	// func append(s []T, vs ...T) []T
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func theRange() {
	// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// // When ranging over a slice, two values are returned for each iteration, the first is the index, and the second is a copy of the element at that index
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func theMap() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var mn = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	// var mn = map[string]Vertex{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }

	fmt.Println(mn)
}

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func theFunction() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	// fmt.Printf("Hello Go!!!\n")
	// fmt.Println("Math is fun!", rand.Intn(200))
	// fmt.Println("What is pi? ", math.Pi)
	// fmt.Printf("Now you have %g problems.\n", math.Sqrt(77))
	// fmt.Println(add(9, 10))
	// a, b := swap("hello", "world")
	// fmt.Println(b, a)
	// fmt.Println(split(66))
	// var i int
	// fmt.Println(i, c, python, java)
	// var c, python, java = true, false, "No!"
	// l := 10
	// fmt.Println(j, k, c, python, java, l)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	// zeroValues()
	// typeConversion()
	// typeInference()
	// const Truth = true
	// fmt.Println("Go rules?", Truth)
	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))
	// theForLoop()
	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	// theSwitch()
	// theDefer()
	// thePointer()
	// fmt.Println(Vertex{1, 2})
	// v := Vertex{1, 2}
	// v.X = 4
	// fmt.Println(v, v.X)
	// va := Vertex{1, 2}
	// p := &va
	// p.X = 1e9
	// fmt.Println(va)
	// fmt.Println(v1, p, v2, v3) // The special prefix & returns a pointer to the struct value
	// theArray()
	// theSlice()
	// nilSlice()
	// theMake()
	// theSlicesOfSlices()
	// theAppend()
	// theRange()
	// theMap()
	// mutatingMaps()
	// theFunction()
	functionClosure()
}
