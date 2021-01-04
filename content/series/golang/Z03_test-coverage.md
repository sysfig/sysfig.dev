---
title: Test Coverage
slug: testing-coverage
date: 2020-12-30T16:14:33-08:00
chapter: z
order: 2
tags:
    - golang
    - testing
draft: true
---

We can add the `-cover` option to `go test` to also display the coverage alongside the test results.

```console
$ go test -cover
PASS
coverage: 100.0% of statements
ok      _/Users/sysfig/gotest   0.119s
```

Go presents code coverage as statement coverage and not line coverage.

You can also save the coverage results into a text file by using the `-coverprofile` option and specifying a filename to save to.

```console
$ go test -coverprofile=coverage
PASS
coverage: 100.0% of statements
ok      _/Users/sysfig/gotest   0.239s
```

The resulting file looks like this:

```
mode: set
/Users/sysfig/gotest/math.go:12.30,13.12 1 1
/Users/sysfig/gotest/math.go:18.2,18.35 1 1
/Users/sysfig/gotest/math.go:13.12,15.3 1 1
/Users/sysfig/gotest/math.go:15.8,15.18 1 1
/Users/sysfig/gotest/math.go:15.18,17.3 1 1
```

We can then use the `cover` tool from `go tool` to read this coverage file, and it will output a HTML page which we can open in a browser and examine the coverage results in a more human-friendly way.

```console
$ go tool cover -html=coverage -o coverage.html
```
