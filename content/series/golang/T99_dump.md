
### Testing Multiple Scenarios

It's good practice to have one test routine for one function. ??How does different scenarios of the same test work then??

sub-tests

## Integration Testing

go help test
go help testflag - flags that control execution of the test

go help testfunc

### Test Data

The `go` tool will ignore any directories named `testdata`, which we can use to hold test data.

```go
file, err := os.Open("./testdata/users.csv")
if err != nil {
  log.Fatal(err)
}
// Do something with the file
```

A test fixture is a clean, fixed, and well-defined environment within which tests run so that tests results are reliably repeatable. You can set up a test fixture by:

- running test suites on a clean container
- wiping/deleting the output directory and databases which may contain artefacts from the previous test run, this resets the environment to its original state.
- loading sample data the tests will consume into a database
- copying a set of sample files that the tests will consume

If you used testing frameworks before, you typically set up your test fixture in functions called `setup`, `before`/`beforeAll` and resets your fixture using functions called `teardown`, `after`/`afterAll`.

## Testing Mutliple Packages

Running tests on files in the current directory using `go test` (with no arguments) is known as _local directory mode_. We can also explicitly specify a package or packages (i.e. directories) for `go test` to test; when we do, it's called _package list mode_. For example:

- `go test .` - test the package in the current directory
- `go test ./math` - test the package in the `./math` directory
- `go test math` - test the `math` module that is found in `$GOPATH`
- `go test ./...` - (ran within a module directory containing packages) test all packages in this module

If we are testing multiple packages, the results of each package would be listed, as well as the result of the entire test run.

When testing packages in package list mode, `go test` will cache all successful package tests that was invoked using a test command that uses only a limited subset of 'cacheable' test flags (`-cpu`, `-list`, `-parallel`, `-run`, `-short`, and `-v`). On subsequent runs, if any of the package source files have not changed, `go test` will use the cached results instead of re-running the tests, and the output will display `(cached)` instead of the elaspsed time in the summary line.


## Tips

### Running a Subset of Tests

`go test` has a `-run` option which allows you to only run test routines whose names matches a regular expression pattern. So `go test -run TestFactorial` will run `TestFactorial/??` and `TestFactorial/??`.


## Test Packages

### Mocks

https://golang.org/pkg/net/http/httptest/
https://golang.org/pkg/testing/iotest/

https://github.com/rakyll/gotest - `go test` but the output is colored
https://github.com/cweill/gotests - test boilerplate generator
(2) https://github.com/stretchr/testify - assert

