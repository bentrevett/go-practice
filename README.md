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