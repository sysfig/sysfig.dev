A panic is a type of error that:

- is not handled by the code (i.e. is not expected/anticipated by the programmer)
- cannot be caught by the compiler
- immediately halts the ordinary flow of control of the function, without executing the rest of the function, returns, and starts panicking.

## What Causes Panics?

- accessing an index of a slice or array beyond its length or capacity
- wrongly asserting the type of a variable
- dereferencing nil pointers
- writing to or closing a closed channel

### Out-of-Bounds Panics

For example, let's say we have an array with a capacity of 4. When we try to access an index of the array that is beyond the capacity (e.g. `10`), the program will panic because we are asking it to do something that it shouldn't do.

```go
package main

func main() {
  scores := [4]int{10, 7, 7, 8}
  scores[10] = 6
}
```

When run, the program will print the following output to the terminal:

```
panic: runtime error: index out of range [10] with length 4

goroutine 1 [running]:
main.main()
        ./panic_1.go:5 +0x1d
exit status 2
```

Let's break this output down:

- `index out of range [10] with length 4` - provides you with information about the root cause of the panic
- `goroutine 1 [running]:` - a Go program can use multiple _goroutines_ to execute code concurrently. This line shows the goroutine ID where the panic occurred, and the state it was in when it occurred.
- `main.main() ./panic_1.go:5 +0x1d` - a _stack trace_ that shows you which package, function, and file and line number the panic originated from.

If the program didn't panic and actually try to write the data (`6`) to the memory address where `scores[10]` would be, it will most likely corrupt the data of whatever is at that address. Therefore, a panic is a safety mechanism to stop a program before it does something harmful and lead to the corruption or loss of data.

### Type Assertion Panics

Let's have another example. Let's suppose we have a variable `i` of type `interface{}` (which means it can be of any type, similar to `any` in languages like TypeScript), and we try to assert that it is an `int`. If it _is_, in fact, an integer, then the program runs fine.

```go
package main

func sum(x, y interface{}) int {
  return x.(int) + y.(int)
}

func main() {
  var x interface{} = 4
  var y interface{} = 5
  sum(x, y)
}
```

```
[no output]
```

But when this type assertion is wrong:

```go
package main

func sum(x, y interface{}) int {
  return x.(int) + y.(int)
}

func main() {
  var x interface{} = "4"
  var y interface{} = 5
  sum(x, y)
}

```

It will produce a panic with the following message:

```
panic: interface conversion: interface {} is string, not int

goroutine 1 [running]:
main.sum(...)
        ./panic_2.go:4
main.main()
        ./panic_2.go:10 +0x45
exit status 2
```

There are two entries in the stack trace this time around, showing that the `main()` function called the `sum()` method on line 10 of `panic_2.go`, but that it's the code on line 4 of `panic_2.go` (within the `sum()` function) that actually caused the panic. By following the stack trace, you can determine the root cause and trigger of the panic.

Type assertions are used to help the compiler determine the type of a variable. Thus, it's not an error that the compiler can catch.

### Dereferencing Nil Pointers

Pointers are references to a specific instance of a type. Let's examine the following code:

```go
package main

type Fruit struct {
  Name string
}

func addCream(f *Fruit) {
  f.Name = f.Name + " with cream"
}

func main() {
  strawberry := &Fruit{"Strawberry"}
  addCream(strawberry)
}
```

`Fruit` is a struct type. `Fruit{"Strawberry"}` is an instance of the `Fruit` type where the `Name` property is set to `"Strawberry"`. Adding a `&` in front of `Fruit{"Strawberry"}` creates a _pointer_ that references that instance. We then assigned that pointer to the variable named `strawberry`. We then passed the pointer to the `addCream` function, which appends the string `" with cream"` to the value of the instance's `Name` property.

The default value for pointers is `nil`. A `nil` pointer does not reference anything. So let's see what happens when we forget to assign a value to `strawberry`.

```go
package main

import "fmt"

type Fruit struct {
  Name string
}

func addCream(f *Fruit) {
  f.Name = f.Name + " with cream"
}

func main() {
  var strawberry *Fruit
  addCream(strawberry)
  fmt.Println("Post-panic code do not continue executing")
}

```

`strawberry` is a `nil` pointer and running the program produces a panic with the following message:

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x105e1a2]

goroutine 1 [running]:
main.addCream(...)
        ./panic_3.go:8
main.main()
        ./panic_3.go:13 +0x22
exit status 2
```

When Go accessed the `Name` property of the `f` struct inside the `addCream()` function call, Go has to first _dereference_ the pointer to the struct. But when the pointer is `nil`, it doesn't reference anything, so dereferencing fails and Go cannot retrieve the value, and thus cannot continue executing the rest of the program. The Go runtime produces a panic within the `addCream()` function, which exits. The panic then bubbles up the call stack to the next function, which is `main()`, and `main()` exits, terminating the program.

The same thing will happen when you try to call a method on a nil pointer.

## Avoiding Panics

Panics produced by the Go runtime are, by definition, errors in the program's logic. So the best way to prevent a panic from happening in production is to write a lot of test cases so that if there's such errors in your code, it will be uncovered by your test suite.

## Capturing Panics

Panics halts the execution of the function and causes the function to panic and returns. If nothing is done about the panic, it will bubble up the call stack and causes the `main()` function to panic and return, at which point the program or goroutine crashes.

But if you are running your program in production and doesn't want a panic to crash the program (perhaps you want to capture the error, log it, and notify the developers to fix it instead), then Go has a featured called _deferred functions_ that can capture the panic and allow the program to continue executing.

Deferred functions are ones that _always_ executes after its calling function has finished executing, even when a panic is encountered. It's similar to `finally` and `ensure` in other languages.

You prepend the `defer` keyword before calling a function to make it a deferred function. We can modify our previous example to include some `defer` function calls that'll inform us when a function has finished executing, regardless of whether it is successful or not.

```go
package main

import "fmt"

type Fruit struct {
  Name string
}

func addCream(f *Fruit) {
  defer fmt.Println("addCream() finished executing")
  f.Name = f.Name + " with cream"
}

func main() {
  defer fmt.Println("main() finished executing")
  var strawberry *Fruit
  addCream(strawberry)
  fmt.Println("Post-panic code continues to execute")
}
```

Running this code generates the following output:

```
addCream() finished executing
main() finished executing
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10a2f48]

goroutine 1 [running]:
main.addCream(0x0)
        ./panic_5.go:11 +0x88
main.main()
        ./panic_5.go:17 +0x8d
exit status 2
```

You can see that both the `fmt.Println` functions were executed despite the panic.

Most of the time, you'd use deferred functions to reliably run some clean up code before the function returns. For example, if you've opened a file for reading with `os.Open`, you can run `file.Close` in a deferred function to ensure that the corresponding file descriptor is _always_ closed, freeing it to be reused, even if an error or panic occurs.

```go
file, err := os.Open("/some/path")
if err != nil {
    return
}
defer file.Close()
```

But we can also use deferred functions to run code that will _capture_ the panic and prevent our program from exiting. We can do this with the built-in `recover()` function.

```go
package main

import "fmt"

type Fruit struct {
  Name string
}

func addCream(f *Fruit) {
  defer func() {
    if error := recover(); error != nil {
      fmt.Println("Panic occurred and recovered:", error)
    }
  }()
  defer fmt.Println("addCream() finished executing")
  f.Name = f.Name + " with cream"
  fmt.Println("Post-panic code in panicking function does not continue to execute")
}

func main() {
  defer fmt.Println("main() finished executing")
  var strawberry *Fruit
  addCream(strawberry)
  fmt.Println("Post-panic code in calling function continues to execute")
}
```

The `recover()` takes no parameters and is only useful when called within a deferred function. When called within the deferred function of a panicking function, it will capture the value of the panic and return it. The programmer can then opt to print an error message but then otherwise continue with the rest of the program.

If the function does not panic, `recover()` returns `nil` and the error message does not get printed.

## Generating Panic

So far, the panics we have encountered are all produced by the Go runtime in the course of executing the code. We know this because the panic message says `panic: runtime error:`. But Go also has a built-in `panic` function that allows you to manually and intentionally generate panics.

`panic()` takes a single string argument, which is the message of the panic.

```go
package main

func bar() {
  panic("baz")
}

func foo() {
  bar()
}

func main() {
  foo()
}
```

```
panic: baz

goroutine 1 [running]:
main.bar(...)
        ./panic_4.go:4
main.foo(...)
        ./panic_4.go:8
main.main()
        ./panic_4.go:12 +0x3a
exit status 2
```

> It's very important to pass a non-`nil` value into `panic()` because otherwise `recover()` won't be able to distinguish between 'no panic', and 'panic with a value of `nil`, and will return `nil` for both scenarios.

So when would you intentionally generate a panic over returning error(s)?

Typically, you don't want to manually generate panic. Panics should only be generated in rare circumstances during the execution of your program. If an error can be handled, then it's better for the function to return an error and handle it. But here are some scenarios when a panic may be preferred:

- you'd use a panic when the error prevents the program or goroutine from correctly executing. For example, if your API service _requires_ the configuration to provide a private key to decrypt tokens but one is not provided, you could argue that even if the API is up, it will not be able to function properly, and thus is not an error that you can recover from, and thus a panic should be produced. Failing fast means you're able to catch errors more quickly.
- in recursive code execution involving a large call stack, it may be easier to panic so that the panic bubbles up to the top of the stack, where it is recovered and an error is returned. The alternative is to return an error and to recursively return the error.
  The `json` package uses panic internally in this way, but its external API only returns `error`.
- you may also use `panic` to mark branches of your code that should _never_ occur. Then, if your code panics because of these `panic` calls, it means the code was able to reach that branch of your code and forces you to revisit your assumptions and fix your code.
