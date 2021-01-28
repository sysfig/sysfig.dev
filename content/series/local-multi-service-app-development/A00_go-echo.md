---
title: Building an Echo Server with Go
slug: go-echo-server
date: 2021-01-27T17:46:37-08:00
chapter: a
order: 0
tags:
    - golang
    - memory
---

The application we are building is called Acme, so let's start by creating a project directory named `acme`.

```console
mkdir acme
```

We are going to be building our echo server in Go - a compiled, statically-typed programming language. So if you haven't already, [download and install Go](https://golang.org/doc/install). By following those steps, you'll be downloading Go's [standard library](https://golang.org/pkg/) as well as the `go` tool, which is a command-line tool we'll use to compile and run our application.

Start by creating a new `echo` directory within our project directory and `cd`-ing into it.

```console
$ mkdir acme/echo && cd acme/echo
```

## Hello World

Then, create a new file named `main.go` where we will write our Go code.

```console
$ touch main.go
```

In the `main.go` file, paste in the following 'Hello World' example:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

In Go, every Go source code file must belong to a _package_. Packages provide a way for us to separate large codebases into smaller [black boxes](https://en.wikipedia.org/wiki/Black_box).

There are two types of packages in Go - applications and libraries. Applications are packages which, after being compiled, can be ran. Web servers and scripts are both examples of an application. Libraries are packages which cannot be run, but, rather, provides utility/helper functions and variables for other libraries and applications.

There's a simple way to distinguish between the two - application packages are named `main`. Within a `main` package, there must always be a function called `main()`, which acts as the entry point where the code will start executing from.

In our example, because we are building an echo server (an application), we are using the `main` package. `Println` is a function that prints an arbitrary string and a newline character to standard output (`stdout`). That function is provided by the [`fmt`](https://golang.org/pkg/fmt/) library package, which is part of the standard library that was downloaded when you downloaded and installed Go.

## Compile and Run

Now that we have some Go code, let's compile and run it. If all goes well, we should expect the application to print the string `Hello, World!` to `stdout`.

We can run `go build` to build an executable from the source code.

```console
$ go build *.go
```

This will generate a binary file called `main` (the same name as the `main.go` file but without the `.go` extension), which we can run.

```console
$ ./main 
Hello, World!
```

As we expected, it printed out a `Hello, World!` message to `stdout`.

## Writing an Echo Server

Now that we know the workflow for compiling and running our Go code, let's continue by implementing the logic of an echo server. An echo server is a simple web server where clients can send requests containing a text message to the server, and the server will simply return the same message.

We could use the [`net/http`](https://golang.org/pkg/net/http/) package from the standard library to implement our server, but there are many web frameworks out there which have a more opinionated but simpler interfaces. One of the more popular frameworks is [Gin](https://gin-gonic.com/) - a _zero-allocation router_ (a router that does not use heap memory).

### Note on Memory Allocation

There are two types of memory that an application can use to store variables - _stack_ and _heap_.

The stack is used for variables declared within a function call. When you declare a variable, stack memory of the right size for that datatype is automatically allocated to hold the values of that variable. And once that function has finished executing, the stack memory used for that variable is deallocated (i.e. reclaimed and freed up) automatically. Thus, the programmer does not need to think about memory allocation or deallocation.

The heap is used when the programmer explicitly allocates a block of memory. This memory is managed by the programmer and once they are finished with it, they must explicitly deallocate it.

In terms of performance, allocating/deallocating as well as accessing stack memory is faster than heap memory. This is because the stack operates as a last-in-first-out (LIFO) stack of contiguous memory blocks of known size, where each block is the combined size of all the data needed to execute a particular function in the call stack. This also means it's fast to allocate memory when a function calls another function because you just allocate memory of the new function from where the memory block for the previous function left off. Accessing memory is also fast because you only ever need to access the top-most block.

Heap memory, on the other hand, is much slower in terms of allocating because the language runtime must find suitable blocks of memory on the end-user's machine. Accessing heap memory is also slower because the runtime must first retrieve the location of the memory block before accessing it.
This is why zero-allocation routers (allocation here refers exclusively to heap allocations) tend to be more performant than routers which allocates memory to the heap.

### Managing Dependencies with Go Modules

Since Go 1.13, the default way to manage (external) dependencies is by using [Go modules](https://blog.golang.org/using-go-modules). A module is simply a collection of related packages that should be version-controlled together.

> Note: A module may contain only a single package.

Go modules uses a `go.mod` file that tracks the dependencies and their versions, as well as the version of Go that the module should be ran on. To use Go modules, we must make our echo server source code its own module as well. We can convert our source code into a module by running `go mod init`.

```console
$ go mod init github.com/sysfig/echo
go: creating new go.mod: module github.com/sysfig/echo
```

This will create a file called `go.mod` in our `echo/` directory with the following content:

```go
module github.com/sysfig/echo

go 1.15
```

The `module` directive defines the _module import path_. ??How is this different from the package's import path??

#### Installing Dependencies

Then, making sure you are within the `echo/` directory, run `go get` to install Gin as a dependency for our module.

```console
$ go get github.com/gin-gonic/gin
go: github.com/gin-gonic/gin upgrade => v1.6.3
```

> Running `go get` within a module directory (one which contains a `go.mod` file) behaves differently than running it outside.

You'll notice that the `go` tool has automatically updated our `go.mod` to include a `require` directive documenting this dependency.

```go
module github.com/sysfig/echo

go 1.15

require github.com/gin-gonic/gin v1.6.3 // indirect
```

The `// indirect` comment marks Gin as an _indirect dependency_, which is a dependency that is not directly imported in our source code. In our case, the `// indirect` comment appears because we haven't updated our source code to use the `gin` package.

You'll also find a new `go.sum` file that records the checksum of all the contents of modules in use as well as the corresponding `go.mod` file. The `go.sum` file is necessary to prevent man-in-the-middle attacks where a malicious party may send you a modified (perhaps insecure) version of the source code.

```
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/gin-contrib/sse v0.1.0 h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=
...
```

### Implementing the Echo Server

Now that we have documented our dependency, update our code to the snippet below where we are using the `gin` package to run a simple API server.

```go
package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.String(200, c.Query("q"))
  })
  r.Run(":8080")
}
```

??`"github.com/gin-gonic/gin"` is the package's _import path_.??

Now when we run `go mod tidy`, the `// indirect` comment is removed because we are now explicitly referencing the `gin` package in our source code.

```go
module github.com/sysfig/echo

go 1.15

require github.com/gin-gonic/gin v1.6.3
```

Next, we can run `go build` to compile the binary and then run it. But during development, we likely only want to run our code and not generate any _[artifacts](https://en.wikipedia.org/wiki/Artifact_(software_development))_. So instead of running `go build` and then executing the resulting binary, we can instead run `go run`. `go run` compiles and run the `main` package without generating any artifacts.

Similar to `go build`, you can pass a list of `.go` source code files to `go run`.

```console
$ go run *.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

Now that the server is up and running, we can test it by sending it a request using [curl](https://curl.se/).

```console
$ curl -i http://localhost:8080/\?q=hello
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Wed, 27 Jan 2021 18:18:16 GMT
Content-Length: 5

hello
```

We passed in a string via the `q` query string parameter, and the echo server echoes back that string in the response body. If we look at the `stdout`, we can also see a new log line for the request.

```log
[GIN] 2021/01/27 - 10:18:16 | 200 |      23.294µs |             ::1 | GET      "/?q=hello"
```

## Summary

In this post, we have gone through writing a Go module that contains a `main` package which, when built and ran, becomes a simple echo server. In the next post, we will take this simple Go application and package it into a Docker container, which will allow us to run and manage it inside a container orchestration framework like Kubernetes.
