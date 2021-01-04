---
title: Benchmarking
slug: testing-benchmarking
date: 2020-12-30T16:14:33-08:00
chapter: t
order: 2
tags:
    - golang
    - testing
    - benchmarking
    - performance
draft: true
---

There are many ways to write a `Factorial` function. In the previous post, we showed two variations of the function that uses the same recursive strategy.

```go
// Version 1
func Factorial(n int) uint64 {
  if n < 2 {
    return 1
  }
  return uint64(n) * Factorial(n-1)
}

// Version 2
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

So how do we know which variation performs better? And does it always perform better regardless of the input? We can find out by writing benchmark functions for both of these variations against different inputs.

Benchmark functions have the form `func BenchmarkXxx(*testing.B)`, and follows this structure:

```go
var result <result-type>
func BenchmarkXyz(b *testing.B) {
    var r <result-type>
    for i := 0; i < b.N; i++ {
        r = Xyz()
    }
    result = r
}
```

The value of `b.N` is adjusted during the execution of the benchmark until either the minimum benchmark time is reached (defaults to 1 second), or a reliably/stable benchmarked time can be determined, whichever is longer.

The variables `r` and `result` are needed to prevent the compiler from optimizing or eliminating the function from being run, and thus reducing the benchmarked time.

Thus, to benchmark our `Factorial` function using the input value of `20`, we'd write:

```go
var result uint64
func BenchmarkFactorial(b *testing.B) {
    var r uint64
    for i := 0; i < b.N; i++ {
        r = Factorial(20)
    }
    result = r
}
```

We can run all benchmark functions by adding the `-bench` option to `go test`

```console
$ go test -bench=.
```

The `-bench` option takes a regular expression that matches the name of the benchmarks to run. The output shows the benchmark functions ran, the number of iterations ran (i.e. the final value of `b.N`), and the average run time for the operation.

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial20-8          21411552                59.4 ns/op
PASS
ok      _/Users/sysfig/gotest   1.271s
```

The benchmarks would run for at least a second each. For very simple or very complex functions, you can override this using the `-benchtime` option.

## Benchmarking Against Different Inputs

But to make it easier to test against different values of input, we can write a private function that takes the input value as a parameter to the function, and call it from the public benchmark functions.

```go
var result uint64

func benchmarkFactorial(i int, b *testing.B) {
  var r uint64
  for n := 0; n < b.N; n++ {
    r = Factorial(i)
  }
  result = r
}

func BenchmarkFactorial0(b *testing.B)  { benchmarkFactorial(0, b) }
func BenchmarkFactorial1(b *testing.B)  { benchmarkFactorial(1, b) }
func BenchmarkFactorial2(b *testing.B)  { benchmarkFactorial(2, b) }
func BenchmarkFactorial3(b *testing.B)  { benchmarkFactorial(3, b) }
func BenchmarkFactorial4(b *testing.B)  { benchmarkFactorial(4, b) }
func BenchmarkFactorial10(b *testing.B) { benchmarkFactorial(10, b) }
func BenchmarkFactorial20(b *testing.B) { benchmarkFactorial(20, b) }
```

Comment out version 2 and run the benchmarks with `go test -bench=.`.

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           701284012                1.64 ns/op
BenchmarkFactorial1-8           720004132                1.64 ns/op
BenchmarkFactorial2-8           379386813                3.09 ns/op
BenchmarkFactorial3-8           248128510                4.81 ns/op
BenchmarkFactorial4-8           179111528                7.29 ns/op
BenchmarkFactorial10-8          47191048                26.7 ns/op
BenchmarkFactorial20-8          21411552                59.4 ns/op
PASS
ok      _/Users/sysfig/gotest   10.550s
```

Then comment out version 1 and uncomment version 2 and re-run the benchmarks.

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           692024133                1.76 ns/op
BenchmarkFactorial1-8           714156800                1.63 ns/op
BenchmarkFactorial2-8           675448486                1.72 ns/op
BenchmarkFactorial3-8           363829836                3.26 ns/op
BenchmarkFactorial4-8           221823795                5.36 ns/op
BenchmarkFactorial10-8          57344438                18.8 ns/op
BenchmarkFactorial20-8          21389043                52.6 ns/op
PASS
ok      _/Users/sysfig/gotest   9.892s
```

It looks like both functions perform equally well for `Factorial(0)` and `Factorial(1)`, which makes sense as they both boil down to the same `if X return 1` logic. But version 2, with its micro-optimization of `if n < 3 { return uint64(n) }` actually reduces the benchmarked time. It appears that removing a single recursion loop chips off a few nanoseconds there.

We can try some other approaches for writing the function. Here, instead of recursion, we are using a local variable and a `for` loop.

```go
// Version 3
func Factorial(n int) uint64 {
  var result uint64 = 1
  for i := 1; i <= n; i++ {
    result *= uint64(i)
  }
  return result
}
```

When we run `go test -bench=.`, we get:

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           703206807                1.73 ns/op
BenchmarkFactorial1-8           520482672                2.25 ns/op
BenchmarkFactorial2-8           339613926                3.53 ns/op
BenchmarkFactorial3-8           279914770                4.29 ns/op
BenchmarkFactorial4-8           278382530                4.27 ns/op
BenchmarkFactorial10-8          237474589                5.05 ns/op
BenchmarkFactorial20-8          157904125                7.56 ns/op
PASS
ok      _/Users/sysfig/gotest   12.545s
```

It seems like this version 3 performs better for larger (4+) inputs, but slower than version 2 for inputs below 4.

We can try combining the short-curcuit approach along with the `for` loop to give us something like this:

```go
// Version 4
func Factorial(n int) uint64 {
  if n < 2 {
    return 1
  }
  if n < 3 {
    return uint64(n)
  }
  var result uint64 = 1
  for i := 1; i <= n; i++ {
    result *= uint64(i)
  }
  return result
}
```

Version 4's benchmarks look like this:

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           540542161                2.01 ns/op
BenchmarkFactorial1-8           579944492                2.02 ns/op
BenchmarkFactorial2-8           646798371                1.82 ns/op
BenchmarkFactorial3-8           404616813                2.99 ns/op
BenchmarkFactorial4-8           280392439                4.08 ns/op
BenchmarkFactorial10-8          249285301                4.76 ns/op
BenchmarkFactorial20-8          164231257                7.49 ns/op
PASS
ok      _/Users/sysfig/gotest   11.063s
```

Let's try one more approach. As you were writing the unit tests, you may come to the realization that there's a limit to how high the input value can be. If we try to write a test case for `30!`, which is `265252859812191058636308480000000` (33 digits), Go would fail to build the binary because `265252859812191058636308480000000` does not fit inside 64 bits.

```console
$ go test
# _/Users/sysfig/gotest [_/Users/sysfig/gotest.test]
./math_test.go:17:8: constant 265252859812191058636308480000000 overflows uint64
FAIL    _/Users/sysfig/gotest [build failed]
```

`uint64` (64-bit unsigned integer) can have a maximum value of `2^64 − 1`, or `18446744073709551615` (20 digits). `20!` is `2432902008176640000` (19 digits) but `21!` is already `51090942171709440000` (20 digits and over `2^64 − 1`). This means our `Factorial` function can only accept positive integers between 0 and 20.

Functions like `Multiply` or `Sum` have an enormously-huge number of possible combinations for its inputs; but out `Factorial` function can only take 21 different values (`0` to `20`). And so we can simply create a map with the all possible inputs as the keys, and the pre-calculated outputs as the values.

```go
// Version 5
var factorials = map[int]uint64{
  0:  1,
  1:  1,
  2:  2,
  3:  6,
  4:  24,
  5:  120,
  6:  720,
  7:  5040,
  8:  40320,
  9:  362880,
  10: 3628800,
  11: 39916800,
  12: 479001600,
  13: 6227020800,
  14: 87178291200,
  15: 1307674368000,
  16: 20922789888000,
  17: 355687428096000,
  18: 6402373705728000,
  19: 121645100408832000,
  20: 2432902008176640000,
}

func Factorial(n int) uint64 {
  return factorials[n]
}
```

And this is the benchmark results for version 5:

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           168200143                7.07 ns/op
BenchmarkFactorial1-8           159913008                7.68 ns/op
BenchmarkFactorial2-8           165133173                7.17 ns/op
BenchmarkFactorial3-8           161319036                7.40 ns/op
BenchmarkFactorial4-8           168475332                7.20 ns/op
BenchmarkFactorial10-8          145387818                8.58 ns/op
BenchmarkFactorial20-8          100000000               10.1 ns/op
PASS
ok      _/Users/sysfig/gotest   13.130s
```

I had expected this method to be the fastest, since `factorials` is a constant and the `Factorial` function is simply doing a map lookup. But the benchmark data actually shows it is significantly slower than version 4. So it looks like creating the map and performing a map lookup is relatively slow.

Let's try another variation where we use a bunch of `if` statements instead of a map.

```go
// Version 6
func Factorial(n int) uint64 {
  if n < 2 {
    return 1
  }
  if n == 2 {
    return 2
  }
  if n == 3 {
    return 6
  }
  if n == 4 {
    return 24
  }
  if n == 5 {
    return 120
  }
  if n == 6 {
    return 720
  }
  if n == 7 {
    return 5040
  }
  if n == 8 {
    return 40320
  }
  if n == 9 {
    return 362880
  }
  if n == 10 {
    return 3628800
  }
  if n == 11 {
    return 39916800
  }
  if n == 12 {
    return 479001600
  }
  if n == 13 {
    return 6227020800
  }
  if n == 14 {
    return 87178291200
  }
  if n == 15 {
    return 1307674368000
  }
  if n == 16 {
    return 20922789888000
  }
  if n == 17 {
    return 355687428096000
  }
  if n == 18 {
    return 6402373705728000
  }
  if n == 19 {
    return 121645100408832000
  }
  return 2432902008176640000
}
```

The benchmarks for this is much more favorable.

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           766182746                1.52 ns/op
BenchmarkFactorial1-8           779824089                1.55 ns/op
BenchmarkFactorial2-8           749115752                1.56 ns/op
BenchmarkFactorial3-8           738537469                1.57 ns/op
BenchmarkFactorial4-8           931112314                1.26 ns/op
BenchmarkFactorial10-8          524524347                2.26 ns/op
BenchmarkFactorial20-8          360830468                3.26 ns/op
PASS
ok      _/Users/sysfig/gotest   9.899s
```

We can also try a long `switch` statement:

```go
// Version 7
func Factorial(n int) uint64 {
  switch n {
  case 0:
    return 1
  case 1:
    return 1
  case 2:
    return 2
  case 3:
    return 6
  case 4:
    return 24
  case 5:
    return 120
  case 6:
    return 720
  case 7:
    return 5040
  case 8:
    return 40320
  case 9:
    return 362880
  case 10:
    return 3628800
  case 11:
    return 39916800
  case 12:
    return 479001600
  case 13:
    return 6227020800
  case 14:
    return 87178291200
  case 15:
    return 1307674368000
  case 16:
    return 20922789888000
  case 17:
    return 355687428096000
  case 18:
    return 6402373705728000
  case 19:
    return 121645100408832000
  }
  return 2432902008176640000
}
```

Which produces a similar benchmark to the `if` statements.

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8           759149526                1.51 ns/op
BenchmarkFactorial1-8           660821928                1.81 ns/op
BenchmarkFactorial2-8           590159343                2.20 ns/op
BenchmarkFactorial3-8           620782159                1.91 ns/op
BenchmarkFactorial4-8           589425992                2.00 ns/op
BenchmarkFactorial10-8          654949964                1.94 ns/op
BenchmarkFactorial20-8          457468696                2.68 ns/op
PASS
ok      _/Users/sysfig/gotest   10.037s
```

In summary, the benchmark data for our different versions of the `Factorial` functions are.

| version\input  | 0    | 1    | 2    | 3    | 4    | 10   | 20   |
|----------------|------|------|------|------|------|------|------|
| 1 (recursion)  | 1.64 | 1.64 | 3.09 | 4.81 | 7.29 | 26.7 | 59.4 |
| 2 (recursion+) | 1.76 | 1.63 | 1.72 | 3.26 | 5.36 | 18.8 | 52.6 |
| 3 (for loop)   | 1.73 | 2.25 | 3.53 | 4.29 | 4.27 | 5.05 | 7.56 |
| 4 (for loop+)  | 2.01 | 2.02 | 1.82 | 2.99 | 4.08 | 4.76 | 7.49 |
| 5 (map)        | 7.07 | 7.68 | 7.17 | 7.40 | 7.20 | 8.58 | 10.1 |
| 6 (if)         | 1.52 | 1.55 | 1.56 | 1.57 | 1.26 | 2.26 | 3.26 |
| 7 (switch)     | 1.51 | 1.81 | 2.20 | 1.91 | 2.00 | 1.94 | 2.68 |

![](/img/factorial-benchmarks.png)

So it seems like the `if` and `switch` approaches are the best.

## Advanced Topics

### Memory Allocations

So far, our benchmarks tell us how fast each of our functions run, but it doesn't tell us how much heap memory each function uses, or how many times the `malloc()` system call is used to allocate that heap memory.

Oftentimes, especially in memory-constrained environments, the "best" function to use isn't always the fastest, but the one that uses the least heap memory. Typically, you'd strike a balance between the two.

So how do we know how much heap memory is used by our function? We have two options:

- Call [`b.ReportAllocs()`](https://golang.org/pkg/testing/#B.ReportAllocs) within our benchmark function to log and report memory allocation statistics of the benchmarked function, or
- Run `go test -bench` with the `-test.benchmem` flag, which will include memory allocation statistics for _all_ benchmarks ran.

```go
func benchmarkFactorial(i int, b *testing.B) {
  b.ReportAllocs()
  var r uint64
  for n := 0; n < b.N; n++ {
    r = Factorial(i)
  }
  result = r
}
```

Let's run the benchmarks against the version 1 of our `Factorial()` function.

```console
$ go test -bench=.
goos: darwin
goarch: amd64
BenchmarkFactorial0-8   704701876  1.68 ns/op  0 B/op  0 allocs/op
BenchmarkFactorial1-8   695891988  1.72 ns/op  0 B/op  0 allocs/op
BenchmarkFactorial2-8   374519816  3.16 ns/op  0 B/op  0 allocs/op
BenchmarkFactorial3-8   240368280  4.94 ns/op  0 B/op  0 allocs/op
BenchmarkFactorial4-8   175318573  6.82 ns/op  0 B/op  0 allocs/op
BenchmarkFactorial10-8  43424978   25.7 ns/op  0 B/op  0 allocs/op
BenchmarkFactorial20-8  20320018   60.5 ns/op  0 B/op  0 allocs/op
PASS
ok      _/Users/sysfig/gotest   10.613s
```

`go test` reports that, on average, there are 0 bytes allocated using 0 allocation calls, no heap memory allocated by the code, which makes sense since ??. (??variables are allocated on the stack??)

??The compiler may optimize your code??

###

Profiling tools `pprof` can help us identify the allocation-heavy portions of the code. It does both CPU profiling and allocation profiling.

https://blog.golang.org/pprof

CPU profile
`-cpuprofile profile_cpu.out`

`go tool pprof -svg profile_cpu.out > profile_cpu.svg`
`go tool pprof -svg profile_cpu.out > profile_cpu.pdf`

Get a call graph

```
$ go test -bench=. -cpuprofile v3.prof  
...
% go tool pprof -pdf v3.prof > v3.pdf 
failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in $PATH
```

`dot` and `gv`

Debian-based: `apt-get install graphviz gv`
macOS: `brew install graphviz`


Flame graphs

Previously, you had to install a tool by Uber called [`go-torch`](https://github.com/uber-archive/go-torch), but as of Go v1.11, flamegraphs are part of the `pprof` tool.

`pprof -http=":8081" [binary] [profile]`

---

configuring power management on your machine. For better or worse, modern CPUs rely heavily on active thermal management which can add noise to benchmark results.

https://blog.golang.org/pprof
