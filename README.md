# Go Practice

## Introduction

Go is a statically typed, compiled programming language.

Statically typed = types are checked at compile time. 
Dynamically typed = types are checked at run time.

Compiled = source code is translated to machine code before being run.
Interpreted = source code is executed line by line.

## The Basics

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

`package main` means this is the main entrance into the code, not part of a shared library.

`import "fmt"` imports the package `fmt` from the standard library. When you do an import, Go will look in the folders indicated by the GOROOT and GOPATH locations. The standard library is in GOROOT and your custom packages should be in GOPATH.

`func` defines a function. `main` is the name of the function and is the entry point of your code, it takes no arguments. 

`fmt.Println` means we use the function `Println` from the imported `fmt` library.

**YOU CAN ONLY IMPORT FUNCTIONS AND VARIABLES THAT START WITH A CAPITAL LETTER!**

Lines do not have to end with a semicolon. 

Programs can be run with `go run <name>.go`.

Packages can be installed with `go get <name>` and then will be downloaded into `GOPATH/<name>` and then can be imported with `import <name>`. 

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add_2(x, y int) int {
	return x + y
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(add_2(42, 13))
}
```

Functions can take zero or more arguments, arguments are in the form of `<name> <type>`.

They can return zero or more arguments. The return type is indicated after the arguments, `int` here.

If consecutive arguments have the same type, the type can be declared after all of the arguments.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

Here's a function that returns multiple arguments. Ignore the `:=` for now.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(10))
}
```

You can also define which values will be returned in the first line of the function declaration. This means you can just write a "naked" return statement. I don't know why you would ever do this.

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

`var` declares a list of variables (can just be one though), ending with their type.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

`var` can set the initial values of variables. If the variable is given an initial value you don't need to define a type, it is inferred from the initial value, i.e. `bool` `bool` `string` here.

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

`:=` is shorthand for `var` with an initial declaration, i.e. `var k = 3` is the same as `k := 3`.

The types available are: bool, string, int, int8/16/32/64 (and uint versions), byte (same as uint8) float32/64 and some others I will probably never use: rune (same as uint32), complex64/128.

int defaults to int32 on 32-bit systems and 64 on 64-bit systems.

When a variable is initialized without an initial value, i.e. `var i int`, it is actually initialized to that type's "zero value":
- 0 for all numeric types
- false for bool
- "" for strings (empty string)

All type conversions must be done explicitly, i.e. `float64(i)`

When declaring a varialble without specifying the type, i.e. `j := i / 2`, it will infer the type from the right hand side, i.e. if `i` was an int, j will also be an int. That code will fail if i is an odd number though as the result will not be an integer.

Constants are declared with `const Pi = 3.14`. They can be strings, booleans or numerics. You cannot define a const with `:=`. I'm guessing the type is inferred again.

Constants are high precision values. They will take the required type depending on their use if it was not initialized with a type.

## Control Flow

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

Go only has one looping construct, the `for` loop which has the standard three components, each separated by a semi-colon.
- init statement
- conditional
- post statement

This is the same as C. Except don't need parenthesis around the three statements. The init and post statements are optional, so you can do a while loop like so:

```go
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

The semicolons can also be dropped, so you can create an actual while loop:

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

An infinite loop can be done with `for { ... }`.

Also has `if` statements, again they don't need parens.

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {
		fmt.Printf("%g = %g\n", v, lim)
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

You can also put a statement before the conditional in the `if` statement. Variables declared inside the `if` statement are in the scope of the `if` all the way to the end of the `else`. The syntax is a little weird, the `else if` and the `else` need to go on the same line as the curly brace that ends the previous `if` statement or else go will throw an error.

```go
package main

import (
	"fmt"
)

func Sqrt(x float64, n int) float64 {
	z := x / 2
	for i := 0; i < n; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 100))
}
```

Above is Newton's method/

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

We have have the `switch` statement, again it can have a init statement. Unlike C there doesn't need to be a `break` statement in each `case`, go will only run the matching statements - `default` if none match.

Switch statements are always evaluated top to bottom until one matches, it does not try the rest. Useful if one of your `case` checks is a function.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

Can also write switch without a condition which is equal to `switch true`. A clean way to write `if-then-else` chains.

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

Go also has `defer` which does not execute the following statement until the function returns. Important to note that the arguments to the function are evaluated immediately. 

Deferred function calls are pushed to a stack and so when the deferred calls are executed they are in a first-in-last-out order, i.e. the code above prints "counting", "done" and then 9 down to 0.

## More Types!

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

A pointer holds the memory address of a value, type `*T` is a pointer to a `T` value. Pointers have a zero value of `nil`. `var p *int` is a pointer to an integer.

The `&` operator generates a pointer to its operand. `var p *int`, `i := 42`, `p = &i`. Now `p` points to `i`. 

The `*` operator denotes the pointers underlying value. `fmt.Println(*p)` prints 42. `*p = 21` changes `i` to 42. This is known as dereferrencing.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X) // prints 1
	v.X = 4
	fmt.Println(v.X) // prints 4
}
```

`struct` is a collection of fields. When initialized the initialize the values from top to bottom. Fields can be accessed using `.`

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

`struct` fields can be accessed by a pointer to a struct. This is actually shorthand for `(*p)X`.

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{Y: 1}  // X:0 is implicit
	v4 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, v4, p)
}

```

Some more on defining `structs`. Initialized from top to bottom with arguments if not named. Implicitly initialize variables inside the struct to their zero values. `p` is a pointer to a `Vertex`

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	numbers := [6] int {1, 2, 3}
	fmt.Println(numbers) // prints [1, 2, 3, 0, 0, 0]
}
```

Here's an array. Syntax is weird. `var <name> [size]type` - feels like the size should go after the type but what do I know?

Arrays are initialized to their zero values.

Arrays cannot be resized! An array's length is part of its type.

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // prints [3, 5, 7]
}
```

Arrays cannot be resized but we can create a slice into an array. `[]T` is a slice of type `T`.

A slice does not store data, think of it as a reference to an array. 

**Changing the value of a slice changes the underlying value of the array. Other slices that share the same underlying array will also change.**

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

A slice literal is an array literal without the length. It immediately creates the array on the right and then builds a slice that references it. 

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

You can slice into slices.

When slicing, you can omit the start and stop values which default to `0` and the length of the array/slice for the high bound.

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len 6 cap 6

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len 0 cap 6

	// Extend its length.
	s = s[:4]
	printSlice(s) // len 3 cap 6

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len 2 cap 4
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

Slices have length and capacity. Length is the number of elements in the slice. Capacity is number of elements in the underlying array from the start of the slice.

The zero value of a slice is `nil`, `var s []int` initializes to `nil`. `nil` is like Python's `None`. A nil slice has 0 length and 0 capacity.

```go
package main

import "fmt"

func main() {
	a := make([]int, 5) // zeroed slice of length and capacity 5
	printSlice("a", a)

	b := make([]int, 0, 5) // first argument is len, second is cap
	printSlice("b", b) // 0 len 5 cap

	c := b[:2]
	printSlice("c", c) // 2 len 5 cap

	d := c[2:5]
	printSlice("d", d) // 3 len 3 cap
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

Slices can be created with `make` - this is how you make dynamically sized arrays.

`make` allocates a zeroed array and returns a slice into that array.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

Slices can contain any type, including slices.

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

`append` is a built-in function for appending to slices. Returns the appended array with the given value. Can be more than one value. Appending increases length and capacity.

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

`range` is a weird type of `for` loop. When iterating over a slice the two values returned are the index and the value, kind of like a permanent `enumerate` from Python.

If you don't want either, replace with a `_`. `for i := range pow` only returns the index so the only one you'd blank out is the index.

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x+y)/2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
```

From the tour of go tutorial. Make a 2-d array and fill with values. First `make` a 2-d array of length `dy`. `pic` is now an `dy` length array of empty arrays. Iterate over the first dimension of the array with `range` and for each make a `dx` length empty `uint8` array. Then populate these arrays, making sure you cast to the desired type.

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

A `map` maps keys to values (think Python `dict`). The zero value of a map is `nil`. A `nil` map has no keys and keys cannot be added to it. `make` returns a map of a given type that is initialized and ready to use - but it needs to be declared first? Why?

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

Also have map literals. If the top level is just a type name, can be omitted, there above can ignore the two `Vertex` after the string keys. 

```go
package main

import "fmt"

func main() {
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
```

Unlike arrays, maps can be mutated. Can add a key with `m[key] = elem` which also overwrites existing keys. Can also delete elements with the built-in `delete`.

Get a value with `elem = m[key]`. You can test if a key is in a map with `elem, ok := m[key]`. If the key is in the map, `ok` is true, or else it's false. If the key is not in the map, `elem` is the zero value for the map's element type. Above would be integer.

```go
func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	tokens := strings.Fields(s)
	for _, token := range tokens {
		counter[token] += 1
	}
	return counter
}
```

Here's a function that splits a string into spaces with `strings.Fields` and counts how many of each "word" there is. `stings.Fields` returns an array of strings. Because the zero initialization of int is 0, don't need to initialize.

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

Functions can also be passed to functions. Need to specify their arguments and return values too.

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

Can also have closures. A closure is a function that references variables outside of its body. The function may access and assign to the referenced variables. Above, each instance of `adder` references their own `sum` variable.

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

Here's how to do Fibonacci with closures.

## Methods