---
title: Testing Basics
slug: testing-basics
date: 2020-12-30T00:00:57-08:00
chapter: t
order: 1
tags:
    - golang
    - testing
---

Let's say we have an exported function called `Factorial` in `math.go`

```go
package math

// Factorial - Get the factorial of a positive integer
func Factorial(n int) uint64 {
  if n < 2 {
    return 1
  }
  return uint64(n) * Factorial(n-1)
}
```

We can write a unit test for it in the same package (i.e. same directory) without importing it. The name of the file should end in `_test.go` (e.g. `math_test.go`). Any files named with the `_test.go` suffix will be excluded from the normal build but included when the `go test` command is run. A test file (`*_test.go`) can contain more than one test routine, but since we only have one function to test, we only need one test routine here.

```go
package math

import "testing"

func TestFactorial(t *testing.T) {
  result := Factorial(4)
  if result != 24 {
    t.Errorf("Factorial(%d) = %d, want %d", 4, result, 24)
  }
}
```

> If your tests are in a different package, the functions under test _must_ be exported and you'd have to import them.

Test routines have the form `func TestXxx(*testing.T)`. The only requirement for test routines names is that they begin with `Test` followed by a capital letter (e.g. `TestX`), but the convention is to follow `Test` by the name of the function you are testing (e.g. `TestFactorial`).

The test routines is passed a pointer to [`testing.T`](https://golang.org/pkg/testing/#T), which is an object that keeps track of the test suite's state. You can report the result of each test by using a subset of its methods:

- [`t.Log`](https://golang.org/pkg/testing/#T.Log) - prints out some text to the error log only when a test fails. Its behavior is the same as `fmt.Println`.
- [`t.Logf`](https://golang.org/pkg/testing/#T.Logf) - similar to `t.Log` but output formatted text similar to `fmt.Printf`, instead of plain text like `fmt.Println`.
- [`t.Fail`](https://golang.org/pkg/testing/#T.Fail) - marks a test as failed but continue executing any remaining test routines.
- [`t.Error`](https://golang.org/pkg/testing/#T.Error) - analogous to running `t.Log()` followed by `t.Fail()`.
- [`t.Errorf`](https://golang.org/pkg/testing/#T.Errorf) - analogous to running `t.Logf()` followed by `t.Fail()`.

Within the test routine, you'll prepare parameters, pass them into the function under test, and check that the result is what you expect. If the result is wrong, we can use the `t.Errorf` to  indicate failure and print out an error message.

We can now run `go test` to run our test. This will make Go compile the package source files and test files in the current directory, and run the resulting binary.

If we run our test now, we will see that it passes (`ok`).

```console
$ go test
PASS
ok      _/Users/sysfig/gotest   0.266s
```

The `go test` output also prints out the package name (`_/Users/sysfig/gotest`) as well as the time it took the test to execute (`0.266s`).

To make sure our tests are working (i.e. it reports a failure when there's a mistake), let's introduce an error into our `Factorial` function. Here, we are trying to short-circuit the function by returning early for arguments smaller than 3. This would work for all positive integers except for 0 (`0! == 1`).

```go
func Factorial(n int) uint64 {
  if n < 3 {
    return uint64(n)
  }
  return uint64(n) * Factorial(n-1)
}
```

When we re-run our tests, the error does not surface because we didn't test against `0`.

```console
$ go test
PASS
ok      _/Users/sysfig/gotest   0.261s
```

We rely on tests to tell us when and where our code's logic doesn't meet our expectations. So when a test passes when we know the logic is wrong, it's a bad test. To make it more robust, we can test against more values.

### Testing Multiple Values

Instead of testing just the factorial of `4`, we can try a wider range of numbers - zero, negative numbers, large numbers, etc.

So let's add a new test case into our test routine.

```go
func TestFactorial(t *testing.T) {
  var result uint64
  result = Factorial(0)
  if result != 1 {
    t.Errorf("Factorial(%d) = %d, want %d", 0, result, 1)
  }
  result = Factorial(4)
  if result != 24 {
    t.Errorf("Factorial(%d) = %d, want %d", 4, result, 24)
  }
}
```

Now when we re-run the tests, it will fail.

```console
$ go test
--- FAIL: TestFactorial (0.00s)
    math_test.go:9: Factorial(0) = 0, want 1
FAIL
exit status 1
FAIL    _/Users/sysfig/gotest   0.283s
```

> Note that the test case that passes did not print out anything, this is because we only call `testing.T`'s reporting methods on failure, and not on success.

However, the test logic for both test cases are the same and repeated in the code. To keep things DRY, a common programming practice is to store a map of inputs against expected output, and loop through the map.

```go
func TestFactorial(t *testing.T) {
  testcases := []struct {
    input  int
    output uint64
  }{
    {0, 1},
    {1, 1},
    {2, 2},
    {3, 6},
    {4, 24},
    {10, 3628800},
    {20, 2432902008176640000},
  }
  for _, testcase := range testcases {
    result := Factorial(testcase.input)
    if result != testcase.output {
      t.Errorf("Factorial(%d) = %d, want %d", testcase.input, result, testcase.output)
    }
  }
}
```

This approach is also known as _table-driven tests_.

https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go

Run `go test` again, and it will still report the one failing case for `0`.

```console
$ go test
--- FAIL: TestFactorial (0.00s)
    math_test.go:21: Factorial(0) = 0, want 1
FAIL
exit status 1
FAIL    _/Users/sysfig/gotest   0.292s
```

But if we add an additional check that accounts for `0`:

```go
func Factorial(n int) uint64 {
  if n < 2 {
    return 1
  }
  if n < 3 {
    return uint64(n)
  }
  return uint64(n) * Factorial(n-1)
}
```

All the test cases would once again pass.

```console
$ go test       
PASS
ok      _/Users/sysfig/gotest   0.253s
```
