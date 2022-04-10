---
title: Testing
slug: testing
date: 2020-12-30T00:00:57-08:00
chapter: t
order: 0
tags:
    - golang
    - testing
---

Go's standard library contains a [`testing`](https://golang.org/pkg/testing/) package, which provides support for automated testing of Go packages. You define test routines and use the built-in `go test` command to run them.

Let's start with an example. Suppose we want to test a function named `IsStringInSlice` inside a file named `strings.go`:

```go
// IsStringInSlice - Checks if a string is contained within a slice of strings
func IsStringInSlice(s string, slice []string) bool {
  for _, item := range slice {
    if item == s {
      return true
    }
  }
  return false
}
```

First, we would need to create a new file that ends with the suffix `_test.go`. It's intuitive to name the test file after the file from which the function under test is. So we will create a new file named `strings_test.go` with the following content:

```go
import "testing"

func TestIsStringInSlice(t *testing.T) {
  
}
```

In it, we need to import the `testing` package.

The name of test routines starts with the word `Test`. For example, 

The `go` tool also provides commands for benchmarking and code coverage.
